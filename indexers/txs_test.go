package indexers

import (
	"encoding/json"
	"fmt"
	"github.com/terra-project/core/x/wasm"
	"github.com/terra-project/mantle-official/test/fixtures"
	"github.com/terra-project/mantle-official/utils"
	"github.com/terra-project/mantle/test"
	"sync"
	"testing"
)

func TestTxs(t *testing.T) {
	simapp, accounts := fixtures.NewTestChain(
		RegisterTxs,
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
			go AppendTxAsync(testBlock, sender, receiver, &wg)
		}

		wg.Wait()

		_ = simapp.Inject(testBlock.ToBlock())
	}

	for{}

}


func TestDeocdeWasm(t *testing.T) {
	acc := test.NewAccount("test")
	testMsg := wasm.NewMsgInstantiateContract(
		acc.GetAddress(),
		0,
		[]byte(fmt.Sprintf(
			"{\"provide_liquidity\": {\"assets\": [{\"info\": {\"native_token\": {\"denom\": \"uusd\"}}, \"amount\": \"6000000000\"}, {\"info\": {\"token\": {\"contract_addr\": \"%s\"}}, \"amount\": \"6000000\"}]}}",
			"testing",
		)),
		nil,
		false,
	)
	marshaled, _ := json.Marshal(testMsg)

	target := utils.DecodeWasm(testMsg, marshaled)

	testMsg2 := wasm.NewMsgExecuteContract(
		acc.GetAddress(),
		acc.GetAddress(),
		[]byte(fmt.Sprintf(
			"{\"provide_liquidity\": {\"assets\": [{\"info\": {\"native_token\": {\"denom\": \"uusd\"}}, \"amount\": \"6000000000\"}, {\"info\": {\"token\": {\"contract_addr\": \"%s\"}}, \"amount\": \"6000000\"}]}}",
			"testing",
		)),
		nil,
	)
	marshaled, _ = json.Marshal(testMsg2)

	target = utils.DecodeWasm(testMsg2, marshaled)
	fmt.Println(string(target))
}