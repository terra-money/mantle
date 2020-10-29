module github.com/terra-project/mantle

go 1.14

replace github.com/CosmWasm/go-cosmwasm => github.com/terra-project/go-cosmwasm v0.10.1-terra

require (
	github.com/cosmos/cosmos-sdk v0.39.1
	github.com/iancoleman/strcase v0.1.2
	github.com/tendermint/tendermint v0.33.7
	github.com/terra-project/core v0.4.0-rc.4
	github.com/terra-project/mantle-compatibility v1.3.2-tequila-0004
	github.com/terra-project/mantle-sdk v0.2.3
)
