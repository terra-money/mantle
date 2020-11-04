package main

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/terra-project/mantle/indexers/account_txs"
	"github.com/terra-project/mantle/indexers/blocks"
	"github.com/terra-project/mantle/indexers/cw20"
	"github.com/terra-project/mantle/indexers/tx_infos"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"

	_ "net/http/pprof"

	tmtypes "github.com/tendermint/tendermint/types"
	"github.com/terra-project/mantle-sdk/app"
	"github.com/terra-project/mantle-sdk/db/badger"
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

	genesis, genesisErr := tmtypes.GenesisDocFromFile(genesisPath)

	// check shasum
	jsonBlob, _ := ioutil.ReadFile(genesisPath)
	shasum := sha1.New()
	shasum.Write(jsonBlob)
	sum := hex.EncodeToString(shasum.Sum(nil))
	
	log.Printf("genesis shasum(sha1)=%v", sum)

	if genesisErr != nil {
		panic(genesisErr)
	}
	mantle := app.NewMantle(
		badgerdb,
		genesis,
		account_txs.RegisterAccountTxs,
		tx_infos.RegisterTxInfos,
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
