# Multicall Client for Ethereum

![Tests Passing](https://github.com/jbrower95/multicall-go/actions/workflows/tests.yaml/badge.svg)


This package provides a Go implementation of the Ethereum [`Multicall3`](https://multicall3.com) contract, enabling efficient batched calls to on-chain contracts using the Ethereum JSON-RPC API. It simplifies making multiple contract read-only (view) calls in a single network request.

## Features

- **Multicall Support**: Aggregate multiple contract calls into a single request, saving on network requests and gas.
- **Batching**: Automatically handles batching of requests when the size exceeds a defined limit.
- **Customizable**: Supports custom block height, batching sizes, and deserialization functions.

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

To create a `MulticallClient`, pass a `context.Context`, an `ethclient.Client`, and optional settings.

```go
client, _ := ethclient.Dial("https://your_rpc_url")
ctx := context.Background()
mc, _ := multicall.NewMulticallClient(ctx, client, nil)
```

### Performing two RPC calls at once

You can perform three kinds of actions;
- Checking the balance of an address
- Fetching the block number that the multicall was read from, and
- Performing an arbitrary smart contract read

In this example, we fetch the balance of `0xdeadbeefdeadbeef` while also 
fetching the current block number. We use `Do()` for this.

```go
balanceCall := multicallClient.GetBalance(common.HexToAddress("0xdeadbeefdeadbeef"))
blockNumberCall := multicallClient.GetBlockNumber()

balance, blockNumber, _ := multicall.Do(multicallClient, balanceCall, blockNumberCall)
fmt.Println("Balance:", balance, "Block Number:", blockNumber)
```

### Performing list of RPC calls

Use DoMany to perform many similarly-typed RPC calls at once:

```go
calls := []*multicall.MultiCallMetaData[big.Int]{
    multicallClient.GetBalance(common.HexToAddress("0xAddress1")),
    multicallClient.GetBalance(common.HexToAddress("0xAddress2")),
}

results, _ := multicall.DoMany(multicallClient, calls...)
for i, result := range *results {
    fmt.Printf("Call %d result: %v\n", i+1, result)
}
```

## Testing

`go test` is run automatically in CI.
