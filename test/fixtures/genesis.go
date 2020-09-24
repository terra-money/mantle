package fixtures

import (
	"fmt"
	"github.com/terra-project/mantle-compatibility/genesis"
	"github.com/terra-project/mantle/app"
	"github.com/terra-project/mantle/test"
	"github.com/terra-project/mantle/types"
)

var maxAccounts = 10
func NewTestChain(indexers ...types.IndexerRegisterer) (*app.Mantle, []test.TestAccount){
	// create accounts, register in keyring
	accounts := make([]test.TestAccount, maxAccounts)
	for i := 0; i< maxAccounts; i++ {
		accounts[i] = test.NewAccount(fmt.Sprintf("test-%d", i))
	}

	// give genesis coins
	genesisAccounts := make([]test.GenesisAccount, maxAccounts)
	for i, account := range accounts {
		genesisAccounts[i] = genesis.NewGenesisAccount(account.GetAddress(), test.Coins{
			{ Denom: "uluna", Amount: test.NewInt(10000000000) },
			{ Denom: "ukrw", Amount: test.NewInt(10000000000) },
			{ Denom: "uusd", Amount: test.NewInt(10000000000) },
		})
	}

	// create genesis
	genesis := genesis.NewGenesis("mantle-test", genesisAccounts...)

	// create app
	simapp := app.NewSimMantle(
		genesis,
		indexers...,
	)

	simapp.Server()

	return simapp, accounts
}