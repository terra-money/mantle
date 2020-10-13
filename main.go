package main

import (
	"github.com/terra-project/mantle-official/indexers/account_txs"
	"github.com/terra-project/mantle-official/indexers/blocks"
	"github.com/terra-project/mantle-official/indexers/cw20"
	"github.com/terra-project/mantle-official/indexers/tx_infos"
	"github.com/terra-project/mantle-official/indexers/txs"
	"github.com/terra-project/mantle/utils"
	"log"
	"os"
	"reflect"
	"strconv"

	_ "net/http/pprof"

	"github.com/terra-project/mantle/app"
	"github.com/terra-project/mantle/db/badger"
)

func main() {
	// env variables
	var dbDir = os.Getenv("DB_DIR")
	var genesisPath = invariant(os.Getenv("GENESIS_PATH"), "genesis is not supplied")
	var endpoint = invariant(os.Getenv("ENDPOINT"), "tendermint endpoint is not supplied")
	var syncUntil = func() uint64 {
		syncUntilString := os.Getenv("SYNC_UNTIL")
		syncUntil, err := strconv.Atoi(syncUntilString)
		if err != nil {
			return 0
		}

		return uint64(syncUntil)
	}()

	log.Printf(
		"[mantle] initializing, dbDir=%s, genesisPath=%s, endpoint=%s, syncUntil=%d",
		dbDir,
		genesisPath,
		endpoint,
		syncUntil,
	)

	badgerdb := badger.NewBadgerDB(dbDir)
	defer badgerdb.Close()

	mantle := app.NewMantle(
		badgerdb,
		utils.GenesisDocFromFile(genesisPath),
		account_txs.RegisterAccountTxs,
		tx_infos.RegisterTxInfos,
		txs.RegisterTxs,
		blocks.RegisterBlocks,
		cw20.RegisterCW20Transfers,
	)

	mantle.Server(1337)
	mantle.Sync(app.SyncConfiguration{
		TendermintEndpoint: endpoint,
		SyncUntil:          syncUntil,
	})
}

func invariant(value string, errorMsg string) string {
	if reflect.ValueOf(value).IsZero() {
		panic(errorMsg)
	}

	return value
}