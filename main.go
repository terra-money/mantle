package main

import (
	"github.com/terra-project/mantle/utils"

	_ "net/http/pprof"

	"github.com/terra-project/mantle/app"
	"github.com/terra-project/mantle/db/badger"
)

func main() {
	badgerdb := badger.NewBadgerDB("")
	defer badgerdb.Close()

	mantle := app.NewMantle(
		badgerdb,
		utils.GenesisDocFromFile("./columbus.json"),
		//indexers.RegisterIndexAsset,
		//indexers.RegisterMirrorTx,
	)

	mantle.Server()
	mantle.Sync(app.SyncConfiguration{
		TendermintEndpoint:     "http://public-node.terra.dev:26657/",
	})

}
