# based on https://github.com/vincentbernat/hellogopher/blob/master/Makefile 
MODULE   = $(shell env GO111MODULE=on $(GO) list -m)

GOCMD       = go 
GOBUILD     = $(GOCMD) build
GOTEST      = $(GOCMD) test
BINARY_NAME = dister
BIN_DEST    = bin/$(BINARY_NAME)

COMMIT_ID   ?= $(shell git rev-parse HEAD|| echo unknown) #  git id if available
VERSION     ?= $(shell git describe --tags --always   --match=v* 2> /dev/null || \
                 echo v0.0.1) # versioning if available 
export GOCACHE := $(CURDIR)/build_cache

all: test build 

build: 

	$(GOBUILD) \
		-tags release \
        -ldflags '-X dist/cmd.Version=$(VERSION) -X dist/cmd.commitID=$(COMMIT_ID)' \
		-o $(BIN_DEST)_$(VERSION) main.go 

test: 
	$(GOTEST) -v ./...
	

compile: 	
	GOOS=linux GOARCH=386 $(GOBUILD) -o $(BIN_DEST)_$(VERSION)-linux-386 main.go
