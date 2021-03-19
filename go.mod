module github.com/terra-project/mantle

go 1.14

replace github.com/CosmWasm/go-cosmwasm => github.com/terra-project/go-cosmwasm v0.10.5

replace github.com/cosmos/cosmos-sdk => github.com/terra-project/cosmos-sdk v0.39.2-public.5

require (
	github.com/cosmos/cosmos-sdk v0.39.2
	github.com/getsentry/sentry-go v0.7.0
	github.com/iancoleman/strcase v0.1.2
	github.com/spf13/viper v1.7.0
	github.com/tendermint/tendermint v0.33.9
	github.com/terra-project/core v0.4.3-0.20210311085928-3a5c3ed1df40
	github.com/terra-project/mantle-sdk v0.3.9
)
