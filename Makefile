# Copyright 2016 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

VERSION :=
TAG := $(shell git describe --abbrev=0 --tags HEAD 2>/dev/null)
COMMIT := $(shell git rev-parse HEAD)
ifeq ($(TAG),)
    VERSION := latest
else
    ifeq ($(COMMIT), $(shell git rev-list -n1 $(TAG)))
        VERSION := $(TAG)
    else
        VERSION := $(TAG)-$(COMMIT)
    endif
endif

container: build quick-container
.PHONY: container

clean:
	rm -f kubernetes-csi
	rm -rf _output
	rm ./pkg/driver/driver.mock.go
.PHONY: clean

quick-container:
	docker build -t kubernetes-csi:latest .
.PHONY: quick-container

glide:
	glide up -v
.PHONY: glide

glide-clean:
	rm -rf ./vendor
	rm -f glide.lock
	glide cc
.PHONY: glide-clean

test:
	go test -v  ./test/csi_driver_test.go
	go test -v  ./test/co_test.go
.PHONY: test

csi-mock:
	mkdir -p ./pkg/driver/
	mockgen -source=${GOPATH}/src/github.com/container-storage-interface/spec/lib/go/csi/csi.pb.go -imports .=github.com/container-storage-interface/spec/lib/go/csi -package=driver -destination=./pkg/driver/driver.mock.go
.PHONY: csi-mock

all build: csi-mock glide
	mkdir -p _output
	go build -o _output/flex-provisioner ./cmd/flex-provisioner/ 
.PHONY: all build
