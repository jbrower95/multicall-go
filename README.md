# Multicall Client for Ethereum

![Tests Passing](https://github.com/jbrower95/multicall-go/actions/workflows/tests.yaml/badge.svg)


This package provides a Go implementation of the Ethereum [`Multicall3`](https://multicall3.com) contract, enabling efficient batched calls to on-chain contracts using the Ethereum JSON-RPC API. It simplifies making multiple contract read-only (view) calls in a single network request.

## Features

- **Multicall Support**: Aggregate multiple contract calls into a single request, saving on network requests and gas.
- **Batching**: Automatically handles batching of requests when the size exceeds a defined limit.
- **Customizable**: Supports custom deserialization of return data, allowing for flexibility in handling different contract responses.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
  - [Creating the Client](#creating-the-client)
  - [Getting Balance](#getting-balance)
  - [Getting Block Number](#getting-block-number)
  - [Custom Contract Calls](#custom-contract-calls)
- [Testing](#testing)
- [License](#license)

## Installation

Install the Go package via:

```bash
go get github.com/jbrower95/multicall-go
```

## Usage

### Creating the Client

### Performing two RPC calls at once

### Performing list of RPC calls

## Testing

`go test` is run automatically in CI.
