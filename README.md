# oidc-cli

[![release](https://github.com/wallanaq/oidc-cli-v2/actions/workflows/release.yaml/badge.svg)](https://github.com/wallanaq/oidc-cli-v2/actions/workflows/release.yaml)
[![Latest Release](https://img.shields.io/github/v/release/wallanaq/oidc-cli-v2)](https://github.com/wallanaq/oidc-cli-v2/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/wallanaq/oidc-cli-v2)](https://goreportcard.com/report/github.com/wallanaq/oidc-cli-v2)
[![Go Reference](https://pkg.go.dev/badge/github.com/wallanaq/oidc-cli-v2.svg)](https://pkg.go.dev/github.com/wallanaq/oidc-cli-v2)

A simple and effective command-line tool for interacting with OpenID Connect (OIDC) providers.

## Features

*   **Version Information**: Check the current version of the CLI (`oidc version`).
*   **Update Checker**: Notify you when a new version is available (`oidc update-check`).

### Work in Progress (WIP)

We are actively working on adding more features to make OIDC interactions seamless from your terminal:

*   `login`: Authenticate with an OIDC provider.
*   `logout`: End the session with the provider.
*   `me`: Fetch user profile information from the `/userinfo` endpoint.
*   And more!

## Installation

### From Release

You can download the pre-compiled binaries for your operating system from the GitHub Releases page.

### From Source

Ensure you have Go installed (version 1.25+ is recommended).

```bash
git clone https://github.com/wallanaq/oidc-cli-v2.git
cd oidc-cli-v2
go build -o oidc ./cmd/oidc/
```

## Usage

Get a list of all available commands:
```bash
oidc --help
```

### Check Version

```sh
oidc version
```

### Check for Updates

```sh
oidc update-check
```

## Development

This project uses GoReleaser to manage releases.

### Building for Local Testing

To build a local version for testing, you can use the `Makefile`.

First, ensure you have GoReleaser installed.

Then, run the following command from the project root to simulate a release without publishing anything:

```bash
# This will create the binaries in a `dist` folder.
make release-snapshot
```

### Using Make

The provided `Makefile` simplifies the development workflow:

*   `make build`: Compiles the project into a single binary named `oidc`.
*   `make test`: Runs all the tests in the project.
*   `make release-snapshot`: Creates a local test release using GoReleaser.
*   `make clean`: Removes build artifacts and the `dist` directory.
