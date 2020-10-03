package utils

import (
	"github.com/terra-project/mantle/test"
	"sync"
)

func AppendTxAsync(testBlock *test.TestkitBlock, sender, receiver test.TestAccount, wg *sync.WaitGroup) {
	testBlock.WithTx(test.NewTx().WithMsg(test.NewMsgSend(
		sender.GetAddress(),
		receiver.GetAddress(),
		test.Coins{
			{ Denom: "uluna", Amount: test.NewInt(20000) },
		},
	)).ToTx(sender))

	wg.Done()
}