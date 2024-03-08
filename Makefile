# Build target
BINARY_NAME = algo
BINARY_PATH = ./target
BINARY = $(BINARY_PATH)/$(BINARY_NAME)

all: test build

build:
	go build -o $(BINARY) -v

test:
	go test -v ./...

clean:
	go clean
	rm -f $(BINARY)

run: build
	$(BINARY) --help

.PHONY: all build test clean run
.DEFAULT_GOAL := run