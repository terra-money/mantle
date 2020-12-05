package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	//	"github.com/terra-project/mantle/indexers/account_txs"
	"github.com/terra-project/mantle/indexers/blocks"
	"github.com/terra-project/mantle/indexers/cw20"
	"github.com/terra-project/mantle/indexers/tx_infos"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"

	_ "net/http/pprof"

	tmtypes "github.com/tendermint/tendermint/types"
	"github.com/terra-project/mantle-sdk/app"
	"github.com/terra-project/mantle-sdk/db/leveldb"
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
	var graphqlPort = func() int {
		graphqlPortString := os.Getenv("PORT")
		graphqlPort, err := strconv.Atoi(graphqlPortString)
		if err != nil {
			return 1337
		}

		return graphqlPort
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

	// Wasm keeper will read home variable from viper and create $home/data/wasm directory
	viper.Set("home", dbDir)
	db := leveldb.NewLevelDB(dbDir)
	defer func() {
		if closeErr := db.Close(); closeErr != nil {
			fmt.Println(closeErr)
		}
	}()

	genesis, genesisErr := tmtypes.GenesisDocFromFile(genesisPath)
	if genesisErr != nil {
		panic(genesisErr)
	}

	// check shasum
	//
	// file is definitely available once we come to this line,
	// skip error check
	jsonBlob, _ := ioutil.ReadFile(genesisPath)
	shasum := sha1.New()
	shasum.Write(jsonBlob)
	sum := hex.EncodeToString(shasum.Sum(nil))

	log.Printf("genesis shasum(sha1)=%v", sum)

	mantle := app.NewMantle(
		db,
		genesis,
		//		account_txs.RegisterAccountTxs,
		tx_infos.RegisterTxInfos,
		blocks.RegisterBlocks,
		cw20.RegisterCW20Transfers,
	)

	if debugPort != 0 {
		go func() {
			log.Println(http.ListenAndServe(fmt.Sprintf("localhost:%d", debugPort), nil))
		}()
	}

	mantle.Server(graphqlPort)
	mantle.Sync(app.SyncConfiguration{
		TendermintEndpoint: endpoint,
		SyncUntil:          syncUntil,
		Reconnect:          true,
		OnWSError: func(err error) {
			sentry.CaptureException(err)
			sentry.Flush(time.Second)
		},
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
