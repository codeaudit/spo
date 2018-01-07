# SPO

[![Build Status](https://travis-ci.org/spaco/spo.svg)](https://travis-ci.org/spaco/spo)
[![GoDoc](https://godoc.org/github.com/spaco/spo?status.svg)](https://godoc.org/github.com/spaco/spo)
[![Go Report Card](https://goreportcard.com/badge/github.com/spaco/spo)](https://goreportcard.com/report/github.com/spaco/spo)

SPO is a next-generation token build on SPACO network.

SPACO network improves on HTTP/HTTPS in too many ways to be addressed here.

SPACO is small part of OP Redecentralize and OP next-generation net Plan.

## Links

* [spaco.io](https://www.spaco.io)
* [Spaco Blog](https://blog.spaco.io)
* [Spaco Blockchain Explorer](https://explorer.spaco.io)
* [Spaco Distribution Event](https://event.spaco.io)

## Table of Contents

<!-- MarkdownTOC depth="2" autolink="true" bracket="round" -->

- [Installation](#installation)
    - [Go 1.9+ Installation and Setup](#go-19-installation-and-setup)
    - [Go get SPACO](#go-get-SPACO)
    - [Run SPACO from the command line](#run-SPACO-from-the-command-line)
    - [Show SPACO node options](#show-SPACO-node-options)
    - [Run SPACO with options](#run-SPACO-with-options)
- [API Documentation](#api-documentation)
    - [Wallet REST API](#wallet-rest-api)
    - [JSON-RPC 2.0 API](#json-rpc-20-api)
    - [SPACO command line interface](#SPACO-command-line-interface)
- [Development](#development)
    - [Modules](#modules)
    - [Running Tests](#running-tests)
    - [Formatting](#formatting)
    - [Code Linting](#code-linting)
    - [Dependency Management](#dependency-management)
    - [Wallet GUI Development](#wallet-gui-development)
    - [Releases](#releases)
- [Changelog](#changelog)

<!-- /MarkdownTOC -->

## Installation

### Go 1.9+ Installation and Setup

[Golang 1.9+ Installation/Setup](./Installation.md)

### Go get SPACO

```sh
go get https://github.com/spaco/spo/...
```

This will download `github.com/spaco/spo` to `$GOPATH/src/github.com/spaco/spo`.

You can also clone the repo directly with `git clone https://github.com/spaco/spo`,
but it must be cloned to this path: `$GOPATH/src/github.com/spaco/spo`.

### Run SPACO from the command line

```sh
cd $GOPATH/src/github.com/spaco/spo
make run
```

### Show SPACO node options

```sh
cd $GOPATH/src/github.com/spaco/spo
make run-help
```

### Run SPACO with options

```sh
cd $GOPATH/src/github.com/spaco/spo
make ARGS="--launch-browser=false" run
```

## API Documentation

### Wallet REST API

[Wallet REST API](src/gui/README.md).

### JSON-RPC 2.0 API

[JSON-RPC 2.0 README](src/api/webrpc/README.md).

### SPACO command line interface

[CLI command API](cmd/cli/README.md).

## Development

We have two branches: `master` and `develop`.

`develop` is the default branch and will have the latest code.

`master` will always be equal to the current stable release on the website, and should correspond with the latest release tag.

### Modules

* `/src/cipher` - cryptography library
* `/src/coin` - the blockchain
* `/src/daemon` - networking and wire protocol
* `/src/visor` - the top level, client
* `/src/gui` - the web wallet and json client interface
* `/src/wallet` - the private key storage library
* `/src/api/webrpc` - JSON-RPC 2.0 API
* `/src/api/cli` - CLI library

### Running Tests

```sh
make test
```

### Formatting

All `.go` source files should be formatted with `gofmt` or `goimports`.

### Code Linting

Install prerequisites:

```sh
make install-linters
```

Run linters:

```sh
make lint
```

### Dependency Management

Dependencies are managed with [dep](https://github.com/golang/dep).

To install `dep`:

```sh
go get -u github.com/golang/dep
```

`dep` vendors all dependencies into the repo.

If you change the dependencies, you should update them as needed with `dep ensure`.

Use `dep help` for instructions on vendoring a specific version of a dependency, or updating them.

After adding a new dependency (with `dep ensure`), run `dep prune` to remove any unnecessary subpackages from the dependency.

When updating or initializing, `dep` will find the latest version of a dependency that will compile.

Examples:

Initialize all dependencies:

```sh
dep init
dep prune
```

Update all dependencies:

```sh
dep ensure -update -v
dep prune
```

Add a single dependency (latest version):

```sh
dep ensure github.com/foo/bar
dep prune
```

Add a single dependency (more specific version), or downgrade an existing dependency:

```sh
dep ensure github.com/foo/bar@tag
dep prune
```

### Wallet GUI Development

The compiled wallet source should be checked in to the repo, so that others do not need to install node to run the software.

Instructions for doing this:

[Wallet GUI Development README](src/gui/static/README.md)

### Releases

0. If the `master` branch has commits that are not in `develop` (e.g. due to a hotfix applied to `master`), merge `master` into `develop`
1. Compile the `src/gui/dist/` to make sure that it is up to date (see [Wallet GUI Development README](src/gui/static/README.md))
2. Update all version strings in the repo (grep for them) to the new version
3. Update `CHANGELOG.md`: move the "unreleased" changes to the version and add the date
4. Merge these changes to `develop`
5. On the `develop` branch, make sure that the client runs properly from the command line (`./run.sh`)
6. Build the releases and make sure that the Electron client runs properly on Windows, Linux and macOS. Delete these releases when done.
7. Make a PR merging `develop` into `master`
8. Review the PR and merge it
9. Tag the master branch with the version number. Version tags start with `v`, e.g. `v0.20.0`.
10. Make sure that the client runs properly from the `master` branch
11. Create the release builds from the `master` branch (see [Create Release builds](electron/README.md))

If there are problems discovered after merging to master, start over, and increment the 3rd version number.
For example, `v0.20.0` becomes `v0.20.1`, for minor fixes.

#### Creating release builds

[Create Release builds](electron/README.md).

## Changelog

[CHANGELOG.md](CHANGELOG.md)

## Acknowledgement

Thank Skycoin team especially Synth and Steve! Thank you for your support！
