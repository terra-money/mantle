module github.com/terra-money/mantle

go 1.14

replace github.com/CosmWasm/go-cosmwasm => github.com/terra-money/go-cosmwasm v0.10.5

replace github.com/cosmos/cosmos-sdk => github.com/terra-money/cosmos-sdk v0.39.2-public.5

require (
	github.com/cosmos/cosmos-sdk v0.39.2
	github.com/getsentry/sentry-go v0.7.0
	github.com/iancoleman/strcase v0.1.2
	github.com/spf13/viper v1.7.0
	github.com/tendermint/tendermint v0.33.9
	github.com/terra-money/core v0.4.6-0.20210408065940-14f6b52b7624
	github.com/terra-money/mantle-sdk v0.3.11
)
