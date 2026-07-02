SHELL=/usr/bin/env bash

all: build
.PHONY: all

unexport GOFLAGS

GOVERSION:=$(shell go version | cut -d' ' -f 3 | sed 's/^go//' | awk -F. '{printf "%d%03d%03d", $$1, $$2, $$3}')
ifeq ($(shell expr $(GOVERSION) \< 1018000), 1)
$(warning Your Golang version is go$(shell expr $(GOVERSION) / 1000000).$(shell expr $(GOVERSION) % 1000000 / 1000).$(shell expr $(GOVERSION) % 1000))
$(error Update Golang version to at least 1.18.0)
endif

BINS:=

ldflags=-X=github.com/unibaseio/da-sdk-go/build.CurrentCommit=+git.$(subst -,.,$(shell git describe --always --match=NeVeRmAtCh --dirty 2>/dev/null || git rev-parse --short HEAD 2>/dev/null))+$(shell date "+%F.%T%Z")
ifneq ($(strip $(LDFLAGS)),)
	ldflags+=-extldflags=$(LDFLAGS)
endif

GOFLAGS+=-ldflags="$(ldflags)"

hub: $(BUILD_DEPS)
	rm -f bin/hub-edge
	GOOS=linux GOARCH=amd64 go build $(GOFLAGS) -o bin/hub-edge ./app/hub
.PHONY: hub
BINS+=bin/hub

# unibase CLI — a client tool, built native (not cross-compiled).
cli: $(BUILD_DEPS)
	rm -f bin/unibase
	go build $(GOFLAGS) -o bin/unibase ./app/cli
.PHONY: cli
BINS+=bin/unibase


build: hub 

.PHONY: build

clean:
	rm -f $(BINS)
.PHONY: clean