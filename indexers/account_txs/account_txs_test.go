package account_txs

import (
	"fmt"
	"github.com/terra-project/mantle-official/indexers/tx_infos"
	"github.com/terra-project/mantle-official/test/fixtures"
	"github.com/terra-project/mantle-official/utils"
	"github.com/terra-project/mantle-sdk/test"
	"sync"
	"testing"
)

func TestAccountTxs(t *testing.T) {
	simapp, accounts := fixtures.NewTestChain(
		RegisterAccountTxs,
		tx_infos.RegisterTxInfos,
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
			go utils.AppendTxAsync(testBlock, sender, receiver, &wg)
		}

		wg.Wait()

		_ = simapp.Inject(testBlock.ToBlock())
	}

	// inject cw20 token
	var owner = accounts[0]
	var recipient = accounts[1]

	// store code
	_ = utils.MustPass(simapp.Inject(test.NewBlock().
		WithTx(test.NewTx().
			WithMsg(test.NewMsgStoreCode(
				accounts[0].GetAddress(),
				utils.GetWasmBytes("../../test/fixtures/terraswap_token.wasm"),
			),
		).ToTx(owner)).
	ToBlock()))

	// instantiate
	tokenAddress := utils.GetContractAddressFromInstantiate(simapp.Inject(test.NewBlock().WithTx(test.NewTx().WithMsg(test.NewMsgInstantiateContract(
		owner.GetAddress(),
		1,
		[]byte(fmt.Sprintf(
			"{\"initial_balances\":[{\"address\":\"%s\",\"amount\":\"100000000000\"}],\"symbol\":\"TestToken\",\"name\":\"TestTokenName\",\"mint\":{\"cap\":\"100000000000\",\"minter\":\"%s\"},\"decimals\":6}",
			owner.GetAddress().String(),
			owner.GetAddress().String(),
		)),
		nil,
		false,
	)).ToTx(owner)).ToBlock()))

	// make transfer
	_ = simapp.Inject(test.NewBlock().WithTx(test.NewTx().WithMsg(test.NewMsgExecuteContract(
		owner.GetAddress(),
		test.AccAddressFromBech32(tokenAddress),
		[]byte(fmt.Sprintf(
			"{\"transfer\":{\"amount\":\"1000000\",\"recipient\":\"%s\"}}",
			recipient.GetAddress().String(),
		)),
		nil,
	)).ToTx(owner)).ToBlock())

	// check indexed result
	for{}

	// test cw20 receive too

}
