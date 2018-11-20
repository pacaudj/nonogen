# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=nonogen

all: build

build:
		$(GOBUILD) -o $(BINARY_NAME) -v

test:
		$(GOTEST) -v .
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm  ./test-output/*
run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)
