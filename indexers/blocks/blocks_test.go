package blocks

import (
	"fmt"
	"github.com/terra-project/mantle-official/indexers"
	"github.com/terra-project/mantle-official/indexers/tx_infos"
	"github.com/terra-project/mantle-official/indexers/txs"
	"github.com/terra-project/mantle-official/test/fixtures"
	"github.com/terra-project/mantle/test"
	"sync"
	"testing"
)

func TestIndexBlocks(t *testing.T) {
	simapp, accounts := fixtures.NewTestChain(
		txs.RegisterTxs,
		tx_infos.RegisterTxInfos,
		RegisterBlocks,
	)

	fmt.Println("list of all accounts")
	for _, account := range accounts {
		fmt.Println(account.GetAddress().String())
	}

	// create random blocks
	for i := 0; i < 20; i++ {
		testBlock := test.NewBlock()
		wg := sync.WaitGroup{}
		wg.Add(len(accounts)-1)
		for j := 0; j < len(accounts)-1; j++ {
			sender := accounts[j]
			receiver := accounts[j+1]
			go indexers.AppendTxAsync(testBlock, sender, receiver, &wg)
		}

		wg.Wait()

		_ = simapp.Inject(testBlock.ToBlock())
	}

	for{}
}
