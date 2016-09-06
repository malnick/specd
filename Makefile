VERSION := $(shell git describe --tags)
REVISION := $(shell git rev-parse --short HEAD)

LDFLAGS := -X github.com/mesosphere/dcos-signal/signal.VERSION=$(VERSION) -X github.com/mesosphere/dcos-signal/signal.REVISION=$(REVISION) 
FILES := $(shell go list ./... | grep -v vendor)

all: test install

test:
	@echo "+$@"
	go test $(FILES)  -cover

build: 
	@echo "+$@"
	go build -v -ldflags '$(LDFLAGS)' $(FILES)

install:
	@echo "+$@"
	go install -v -ldflags '$(LDFLAGS)' $(FILES)
