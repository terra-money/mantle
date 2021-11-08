# Mantle

> 🚧 This package is deprecated. For columbus-5 compatible version of mantle, [this](https://github.com/terra-money/mantlemint) is what you're looking for. 🚧

This repository contains official indexers built on [mantle-sdk](https://github.com/terra-money/mantle-sdk).

## Installation and run

> Installation process is likely to change once it hits a major release! Current installation steps are intended for internal developers only.

In order to run an instance of mantle, 

```sh
git clone https://github.com/terra-money/mantle.git
cd mantle
go mod download
# then, follow instructions
```

### Dependencies

`mantle` depends on 3 modules:
- [terra-money/core](https://github.com/terra-money/core): official terra network repository.
- [terra-money/mantle-sdk](https://github.com/terra-money/mantle-sdk): official `mantle-sdk`.
- [terra-money/mantle-compatibility](https://:github.com/terra-money/mantle-compatibility): compatibility provider for mantle-sdk.

Versions of these packages need to adjusted in `go.mod` file to match the version of the network you are trying to index.

> The `go.mod` file in `master` branch is set to run with the latest mainnet.

### Adjusting package versions

> Unlisted networks are to be supported in the future.

#### terra-money/core

`core` version **MUST** match with the network version you're trying to index. Fix core's version in `go.mod` file accordingly.

- columbus-3: github.com/terra-money/core v0.3.7
- columbus-4: github.com/terra-money/core v0.4.0
- tequila-4: github.com/terra-money/core v0.4.0-rc.5

#### terra-money/compatibility

Adjust package version accordingly:

- columbus-3: github.com/terra-money/mantle-compatibility v1.2.1-columbus-3
- columbus-4: github.com/terra-money/mantle-compatibility v1.2.1-tequila-rc4
- tequila-4: github.com/terra-money/mantle-compatibility v1.2.1-tequila-rc4

#### terra-money/mantle

We recommed using the latest version of mantle.


### Build

```
go build main.go
```

### Load genesis file

You'll need to use the same genesis file for the network you're trying to index.

- Download from [here](https://docs.terra.money/node/join-network.html#picking-a-network).
- Save it as `genesis.json`.

### Run mantle instance

```
GENESIS_PATH=genesis.json \
ENDPOINT=<your-network-tendermint-rpc> \
DB_DIR=mantle-db \
GRAPHQL_PORT=1337
./mantle
```

### Access indexed data through graphql

Default graphql port is `1337`, which you can change by setting `GRAPHQL_PORT` environmental variable.

There's also a graphql playground available. You can access it at `http://localhost:1337`.


# LICENSE


This software is licensed under the Apache 2.0 license. Read more about it here.

© 2020 Terraform Labs, PTE.
