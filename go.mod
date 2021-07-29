module github.com/terra-money/mantle

go 1.14

replace github.com/CosmWasm/go-cosmwasm => github.com/terra-money/go-cosmwasm v0.10.5

replace github.com/terra-money/mantle-sdk => ../mantle-sdk

replace github.com/cosmos/cosmos-sdk => github.com/terra-money/cosmos-sdk v0.39.2-public.6-mutex-2

replace github.com/terra-money/mantle-compatibility => ../mantle-compatibility/columbus-4

require (
	github.com/cosmos/cosmos-sdk v0.39.2
	github.com/getsentry/sentry-go v0.7.0
	github.com/iancoleman/strcase v0.1.2
	github.com/spf13/viper v1.7.0
	github.com/tendermint/tendermint v0.33.9
	github.com/terra-money/mantle-sdk v0.3.12-0.20210625171532-d0b70a63fbe8
	github.com/terra-project/core v0.4.6-0.20210408065940-14f6b52b7624
)
