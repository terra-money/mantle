module github.com/terra-project/mantle

go 1.14

replace github.com/CosmWasm/go-cosmwasm => github.com/terra-project/go-cosmwasm v0.10.4

require (
	github.com/cosmos/cosmos-sdk v0.39.2
	github.com/getsentry/sentry-go v0.7.0
	github.com/iancoleman/strcase v0.1.2
	github.com/spf13/viper v1.7.0
	github.com/tendermint/tendermint v0.33.9
	github.com/terra-project/core v0.4.2
	github.com/terra-project/mantle-sdk v0.3.7-public-lcd
)
