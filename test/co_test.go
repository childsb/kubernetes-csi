package test

import (
	"golang.org/x/net/context"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	gomock "github.com/golang/mock/gomock"
	mock_driver "github.com/csi-volumes/kubernetes-csi/pkg/driver"

	"fmt"
)

// SafeGoroutineTester is an implementation of the mock ... interface
// which can be used to use the mock functions in another go routine.
//
// The major issue is that the golang mock framework uses t.Fatalf()
// which causes a deadlock when called in another goroutine. To avoid
// this issue, this simple implementation prints the error then panics,
// which avoids the deadlock.
type SafeGoroutineTester struct{}

// Errorf prints the error to the screen then panics
func (s *SafeGoroutineTester) Errorf(format string, args ...interface{}) {
	fmt.Printf(format, args)
	panic("MOCK TEST ERROR")
}

// Fatalf prints the error to the screen then panics
func (s *SafeGoroutineTester) Fatalf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
	panic("MOCK TEST FATAL FAILURE")
}

func TestPluginInfoResponse(t *testing.T) {

	// Setup mock
	m := gomock.NewController(t)
	defer m.Finish()
	driver := mock_driver.NewMockIdentityServer(m)

	// Setup input
	in := &csi.GetPluginInfoRequest{
		Version: &csi.Version{
			Major: 0,
			Minor: 1,
			Patch: 0,
		},
	}

	// Setup mock outout
	out := &csi.GetPluginInfoResponse{
		Reply: &csi.GetPluginInfoResponse_Result_{
			Result: &csi.GetPluginInfoResponse_Result{
				Name:          "mock",
				VendorVersion: "0.1.1",
				Manifest: map[string]string{
					"hello": "world",
				},
			},
		},
	}

	// Setup expectation
	driver.EXPECT().GetPluginInfo(nil, in).Return(out, nil).Times(1)

	// Actual call
	r, err := driver.GetPluginInfo(nil, in)
	name := r.GetResult().GetName()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	if name != "mock" {
		t.Errorf("Unknown name: %s\n", name)
	}
}

func TestGRPCGetPluginInfoReponse(t *testing.T) {

	// Setup mock
	m := gomock.NewController(&SafeGoroutineTester{})
	defer m.Finish()
	driver := mock_driver.NewMockIdentityServer(m)

	// Setup input
	in := &csi.GetPluginInfoRequest{
		Version: &csi.Version{
			Major: 0,
			Minor: 1,
			Patch: 0,
		},
	}

	// Setup mock outout
	out := &csi.GetPluginInfoResponse{
		Reply: &csi.GetPluginInfoResponse_Result_{
			Result: &csi.GetPluginInfoResponse_Result{
				Name:          "mock",
				VendorVersion: "0.1.1",
				Manifest: map[string]string{
					"hello": "world",
				},
			},
		},
	}

	// Setup expectation
	// !IMPORTANT!: Must set context expected value to gomock.Any() to match any value
	driver.EXPECT().GetPluginInfo(gomock.Any(), in).Return(out, nil).Times(1)

	// Create a new RPC
	server := mock_driver.NewMockCSIDriver(&mock_driver.MockCSIDriverServers{
		Identity: driver,
	})
	conn, err := server.Nexus()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	defer server.Close()

	// Make call
	c := csi.NewIdentityClient(conn)
	r, err := c.GetPluginInfo(context.Background(), in)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	name := r.GetResult().GetName()
	if name != "mock" {
		t.Errorf("Unknown name: %s\n", name)
	}
}
