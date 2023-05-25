[![GoDoc](https://godoc.org/github.com/INFURA/go-ethlibs?status.svg)](http://godoc.org/github.com/INFURA/go-ethlibs)
[![Go Report Card](https://goreportcard.com/badge/github.com/INFURA/go-ethlibs)](https://goreportcard.com/report/github.com/INFURA/go-ethlibs)

# go-ethlibs: Go Ethereum Libraries

A sandbox of helpers related to interacting with Ethereum nodes.

We hope code snippets in this repository will come in useful for others, but
please don't expect any stability guarantees. We'd like for the best pieces to
eventually graduate into stand-alone packages.


## Overview

- `eth`: Helpers for serializing/deserializing Ethereum JSONRPC types
- `jsonrpc`: JSONRPC request and response parsing
- `node`: A proto-ethclient in the `node` namespace
- `rlp`: Independent implementation of RLP parsing

## Compile

https://geth.ethereum.org/docs/tools/abigen

```sh
# generate abi
solc --abi ./contracts/Counter.sol -o gen
# generate bytecode
solc --bin ./contracts/Counter.sol -o ./gen/Counter.bin 
# gen binding files
abigen --abi ./gen/Counter.abi --pkg gen --type Counter --out Counter.go
```

## License

MIT
