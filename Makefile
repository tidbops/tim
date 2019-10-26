# Set DEBUGGER=1 to build debug symbols
LDFLAGS = $(if $(DEBUGGER),,-s -w) $(shell ./hack/version.sh)

# SET DOCKER_REGISTRY to change the docker registry
DOCKER_REGISTRY := $(if $(DOCKER_REGISTRY),$(DOCKER_REGISTRY),localhost:5000)

GOVER_MAJOR := $(shell go version | sed -E -e "s/.*go([0-9]+)[.]([0-9]+).*/\1/")
GOVER_MINOR := $(shell go version | sed -E -e "s/.*go([0-9]+)[.]([0-9]+).*/\2/")
GO111 := $(shell [ $(GOVER_MAJOR) -gt 1 ] || [ $(GOVER_MAJOR) -eq 1 ] && [ $(GOVER_MINOR) -ge 11 ]; echo $$?)
ifeq ($(GO111), 1)
$(error Please upgrade your Go compiler to 1.11 or higher version)
endif

# Enable GO111MODULE=on explicitly, disable it with GO111MODULE=off when necessary.
export GO111MODULE := on
GOOS := $(if $(GOOS),$(GOOS),linux)
GOARCH := $(if $(GOARCH),$(GOARCH),amd64)
GOENV  := GO15VENDOREXPERIMENT="1" GOOS=$(GOOS) GOARCH=$(GOARCH)
GO     := $(GOENV) CGO_ENABLED=0 go build
CGO   := $(GOENV) CGO_ENABLED=1 go build
GOTEST := CGO_ENABLED=0 go test -v -cover

default: build

build: tim tim-server

tim:
	$(CGO) -ldflags '$(LDFLAGS)' -o bin/tim cmd/tim/*.go

tim-server:
	GO111MODULE=off go get github.com/jessevdk/go-assets-builder
	go-assets-builder pkg/server/dashboard/templates -o pkg/server/dashboard/templates/assets.go -s /pkg/server/dashboard/templates/  -p templates
	$(CGO) -ldflags '$(LDFLAGS)' -o bin/tim-server cmd/tim-server/*.go
