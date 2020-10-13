package token

import (
	"github.com/terra-project/mantle-official/indexers/cw20"
	"github.com/terra-project/mantle-official/indexers/tx_infos"
	"github.com/terra-project/mantle/types"
	"reflect"
)

type TokenTransfer struct {
	Height       uint64
	Sender       string `model:"index"`
	Recipient    string `model:"index"`
	Contract     string `model:"index"`
	Address      string `model:"index"`
	TransferType string `model:"index"` // send|receive
	Amount       string `model:"index"` // wasm:uint128
	TxHash       string `model:"index"`
	IsCW20       bool   `model:"index"`
	Denom        string `model:"index"`
}

type TokenTransfers []TokenTransfer

func RegisterTokenTransfers(register types.Register) {
	register(
		IndexTokenTransfers,
		reflect.TypeOf((*TokenTransfers)(nil)),
	)
}

func IndexTokenTransfers(query types.Query, commit types.Commit) error {
	gatherAllTrasnfers := new(struct {
		CW20Transfers cw20.CW20Transfers
		Txs           tx_infos.TxInfos
	})

	if queryErr := query(gatherAllTrasnfers, nil); queryErr != nil {
		return queryErr
	}

	var commitTargets TokenTransfers

	return nil
}
