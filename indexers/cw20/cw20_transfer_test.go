package cw20

import (
	"fmt"
	"github.com/terra-project/mantle-official/indexers/tx_infos"
	"github.com/terra-project/mantle-official/test/fixtures"
	"github.com/terra-project/mantle/app"
	"github.com/terra-project/mantle/test"
	"github.com/terra-project/mantle/types"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestCW20(t *testing.T) {
	simapp, accounts := fixtures.NewTestChain(
		tx_infos.RegisterTxInfos,
		app.TrackIndexerResult(
			RegisterCW20Transfers,
			func(_ []app.QueryRecord, commitHistory []app.CommitRecord) {
				fmt.Println(commitHistory)
			},
		),
	)

	fmt.Println("list of all accounts")
	for _, account := range accounts {
		fmt.Println(account.GetAddress().String())
	}

	var owner = accounts[0]
	var recipient = accounts[1]

	// store code
	_ = mustPass(simapp.Inject(test.NewBlock().
		WithTx(test.NewTx().
			WithMsg(test.NewMsgStoreCode(
				accounts[0].GetAddress(),
				getWasmBytes("../../test/fixtures/terraswap_token.wasm"),
			),
		).ToTx(owner)).
	ToBlock()))

	// instantiate
	tokenAddress := getContractAddressFromInstantiate(simapp.Inject(test.NewBlock().WithTx(test.NewTx().WithMsg(test.NewMsgInstantiateContract(
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

}

func getWasmBytes(p string) (wasmBytes []byte) {
	filename, _ := filepath.Abs(p)
	var wasmBytesErr error
	if wasmBytes, wasmBytesErr = ioutil.ReadFile(filename); wasmBytesErr != nil {
		panic(wasmBytesErr)
	}

	return
}


func getContractAddressFromInstantiate(result types.BaseState) string {
	var addr string
	for _, event := range result.DeliverTxResponses[0].Events {
		switch event.Type {
		case "instantiate_contract":
			for _, attr := range event.Attributes {
				switch string(attr.Key) {
				case "contract_address":
					addr = string(attr.Value)
				}
			}
		}
	}

	return addr
}

func mustPass(result types.BaseState) types.BaseState {
	for _, result := range result.DeliverTxResponses {
		if result.IsErr() || !result.IsOK() {
			panic(result.Log)
		}
	}

	return result
}