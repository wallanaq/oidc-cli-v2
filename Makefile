# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
BINARY_NAME=oidc

# GoReleaser
GORELEASER_CMD=goreleaser

.PHONY: all build test clean release-snapshot

all: build

# Build the binary
build:
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/oidc/

# Run tests
test:
	$(GOTEST) -v ./...

# Clean up build artifacts and binaries
clean:
	rm -f $(BINARY_NAME)
	rm -rf ./dist

# Create a local snapshot release for testing
release-snapshot:
	$(GORELEASER_CMD) release --snapshot --clean
