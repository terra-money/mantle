package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	sentry "github.com/getsentry/sentry-go"
	"github.com/terra-project/mantle/indexers/account_txs"
	"github.com/terra-project/mantle/indexers/blocks"
	"github.com/terra-project/mantle/indexers/cw20"
	"github.com/terra-project/mantle/indexers/tx_infos"

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
	var sentryDsn = os.Getenv("SENTRY_DSN")
	var debugPort = func() uint64 {
		debugPortString := os.Getenv("DEBUG_PORT")
		debugPort, err := strconv.Atoi(debugPortString)
		if err != nil {
			return 0
		}

		return uint64(debugPort)
	}()
	var syncUntil = func() uint64 {
		syncUntilString := os.Getenv("SYNC_UNTIL")
		syncUntil, err := strconv.Atoi(syncUntilString)
		if err != nil {
			return 0
		}

		return uint64(syncUntil)
	}()

	// init sentry
	if sentryDsn != "" {
		initSentry(sentryDsn)
		defer sentry.Flush(time.Second * 2)
	}

	// init mantle
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

	if debugPort != 0 {
		go func() {
			log.Println(http.ListenAndServe(fmt.Sprintf("localhost:%d", debugPort), nil))
		}()
	}

	mantle.Server(1337)
	mantle.Sync(app.SyncConfiguration{
		TendermintEndpoint: endpoint,
		SyncUntil:          syncUntil,
	})
}

func initSentry(sentryDsn string) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:   sentryDsn,
		Debug: true,
	})

	if err != nil {
		log.Fatalf("tried to initialize sentry, but got error: %s", err)
	}

}

func invariant(value string, errorMsg string) string {
	if reflect.ValueOf(value).IsZero() {
		panic(errorMsg)
	}

	return value
}
