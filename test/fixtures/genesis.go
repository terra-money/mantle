package fixtures

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/bech32"
	types2 "github.com/tendermint/tendermint/types"
	"github.com/terra-project/mantle-compatibility/genesis"
	"github.com/terra-project/mantle-sdk/app"
	"github.com/terra-project/mantle-sdk/test"
	"github.com/terra-project/mantle-sdk/types"
)

var maxAccounts = 10
func NewTestChain(indexers ...types.IndexerRegisterer) (*app.Mantle, []test.TestAccount){
	// create accounts, register in keyring
	accounts := make([]test.TestAccount, maxAccounts)
	for i := 0; i< maxAccounts; i++ {
		accounts[i] = test.NewAccount(fmt.Sprintf("test-%d", i))
	}

	// give genesis coins
	// first one would be the validator
	genesisAccounts := make([]test.GenesisAccount, maxAccounts)
	for i, account := range accounts {
		genesisAccounts[i] = genesis.NewGenesisAccount(account.GetAddress(), test.Coins{
			{ Denom: "uluna", Amount: test.NewInt(10000000000) },
			{ Denom: "ukrw", Amount: test.NewInt(10000000000) },
			{ Denom: "uusd", Amount: test.NewInt(10000000000) },
		})
	}

	_, converted, _ := bech32.DecodeAndConvert(genesisAccounts[0].GetAddress().String())
	valString, _ := bech32.ConvertAndEncode("terravaloper", converted)
	valAddr, _ := sdk.ValAddressFromBech32(valString)

	// create genesis
	genesis := genesis.NewGenesis("mantle-test", genesisAccounts...)

	fmt.Println("!!!", accounts[0].GetPubKey())

	genesis.Validators = []types2.GenesisValidator{
		{
			Address: valAddr.Bytes(),
			PubKey:  accounts[0].GetPubKey(),
			Power:   100,
			Name:    "Test",
		},
	}

	// create app
	simapp := app.NewSimMantle(
		genesis,
		indexers...,
	)

	return simapp, accounts
}