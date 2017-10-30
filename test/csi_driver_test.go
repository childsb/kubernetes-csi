package test

import (
	"net"
	"sync"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Example simple driver
// This example assumes that your driver will create the server and listen on
// some unix domain socket or port for tests.
type simpleDriver struct {
	listener net.Listener
	server   *grpc.Server
	wg       sync.WaitGroup
}

func (s *simpleDriver) GetSupportedVersions(
	context.Context, *csi.GetSupportedVersionsRequest) (*csi.GetSupportedVersionsResponse, error) {
	return &csi.GetSupportedVersionsResponse{
		Reply: &csi.GetSupportedVersionsResponse_Result_{
			Result: &csi.GetSupportedVersionsResponse_Result{
				SupportedVersions: []*csi.Version{
					&csi.Version{
						Major: 0,
						Minor: 1,
						Patch: 0,
					},
				},
			},
		},
	}, nil
}

func (s *simpleDriver) GetPluginInfo(
	context.Context, *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
	return &csi.GetPluginInfoResponse{
		Reply: &csi.GetPluginInfoResponse_Result_{
			Result: &csi.GetPluginInfoResponse_Result{
				Name:          "simpleDriver",
				VendorVersion: "0.1.1",
				Manifest: map[string]string{
					"hello": "world",
				},
			},
		},
	}, nil
}

func (s *simpleDriver) goServe() {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.server.Serve(s.listener)
	}()
}

func (s *simpleDriver) Address() string {
	return s.listener.Addr().String()
}

func (s *simpleDriver) Start() error {
	// Listen on a port assigned by the net package
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	s.listener = l

	// Create a new grpc server
	s.server = grpc.NewServer()

	csi.RegisterIdentityServer(s.server, s)
	reflection.Register(s.server)

	// Start listening for requests
	s.goServe()
	return nil
}

func (s *simpleDriver) Stop() {
	s.server.Stop()
	s.wg.Wait()
}

//
// Tests
//
func TestSimpleDriver(t *testing.T) {

	// Setup simple driver
	s := &simpleDriver{}
	err := s.Start()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	defer s.Stop()

	// Setup a connection to the driver
	conn, err := grpc.Dial(s.Address(), grpc.WithInsecure())
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	defer conn.Close()

	// Make a call
	c := csi.NewIdentityClient(conn)
	r, err := c.GetPluginInfo(context.Background(), &csi.GetPluginInfoRequest{})
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	// Verify
	name := r.GetResult().GetName()
	if name != "simpleDriver" {
		t.Errorf("Unknown name: %s\n", name)
	}
}

