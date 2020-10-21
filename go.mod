module github.com/terra-project/mantle-official

go 1.14

replace github.com/CosmWasm/go-cosmwasm => github.com/terra-project/go-cosmwasm v0.10.1-terra

replace github.com/terra-project/mantle-sdk => ../mantle-sdk

require (
	github.com/cosmos/cosmos-sdk v0.39.1
	github.com/iancoleman/strcase v0.1.2
	github.com/terra-project/core v0.4.0-rc.5
	github.com/terra-project/mantle-sdk v0.2.0
	github.com/terra-project/mantle-compatibility v1.2.1-tequila-rc4
)
