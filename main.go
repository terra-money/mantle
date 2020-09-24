package main

import (
	"github.com/terra-project/mantle-official/indexers"
	"github.com/terra-project/mantle/utils"

	_ "net/http/pprof"

	"github.com/terra-project/mantle/app"
	"github.com/terra-project/mantle/db/badger"
)

func main() {
	badgerdb := badger.NewBadgerDB("mantle-db")
	defer badgerdb.Close()

	mantle := app.NewMantle(
		badgerdb,
		utils.GenesisDocFromFile("./tequila.json"),
		indexers.RegisterAccountTxs,
		indexers.RegisterTxs,
	)

	mantle.Server()
	mantle.Sync(app.SyncConfiguration{
		TendermintEndpoint:     "http://public-node.terra.dev:26657/",
	})

}
