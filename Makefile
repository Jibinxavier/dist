GOCMD       = go 
GOBUILD     = $(GOCMD) build
GOTEST      = $(GOCMD) test
BINARY_NAME = dister
BIN_DEST    = bin/$(BINARY_NAME)

COMMIT_ID   ?= $(shell git rev-parse HEAD|| echo unknown) #  git id if available
VERSION     ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
                 echo v0.0.1) # versioning if available 


all: test build 

build: 

	$(GOBUILD) \
		-tags release \
        -ldflags '-X dist/cmd.Version=$(VERSION) dist/cmd.commitID=$(COMMIT_ID)'\
		-o $(BIN_DEST)_$(VERSION) \  
		-v

test: 
	$(GOTEST) -v ./...
	

compile: 	
	GOOS=linux GOARCH=386 $(GOBUILD) -o $(BIN_DEST)_$(VERSION)-linux-386 main.go
