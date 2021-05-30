package cw20

import (
	"encoding/base64"
	"github.com/terra-money/mantle-sdk/types"
	"github.com/terra-money/mantle/indexers/tx_infos"
	utils2 "github.com/terra-money/mantle/utils"
	"reflect"
)

type CW20Transfer struct {
	Height       uint64
	Sender       string `model:"index"`
	Recipient    string `model:"index"`
	Contract     string `model:"index"`
	Address      string `model:"index"`
	TransferType string `model:"index"` // send|receive
	Amount       string `model:"index"` // wasm:uint128
}

type CW20Transfers []CW20Transfer

func RegisterCW20Transfers(register types.Register) {
	register(
		IndexCW20Transfers,
		reflect.TypeOf((*CW20Transfers)(nil)),
	)
}

func IndexCW20Transfers(query types.Query, commit types.Commit) error {
	txInfosQuery := new(struct {
		TxInfos tx_infos.TxInfos
	})

	if queryErr := query(txInfosQuery, nil); queryErr != nil {
		return queryErr
	}

	// iterate over all txInfos, and extract CW20-like transfers
	var commitTargets CW20Transfers
	for _, tx := range txInfosQuery.TxInfos {
		// skip if this tx is not successful
		if !tx.Success {
			continue
		}

		for _, msg := range tx.Tx.Msg {
			// skip if not wasm/MsgExecuteContract
			if msg.Type != "wasm/MsgExecuteContract" {
				continue
			}

			// check if MsgExecuteContract has "transfer" key
			msgPayload := new(struct {
				Sender     string `json:"sender"`
				Contract   string `json:"contract"`
				ExecuteMsg string `json:"execute_msg"`
			})

			utils2.MustUnmarshal([]byte(msg.Value), msgPayload)

			var isValidExecuteMsg = !utils2.IsZero(msgPayload.Sender) && !utils2.IsZero(msgPayload.Contract) && !utils2.IsZero(msgPayload.ExecuteMsg)
			if !isValidExecuteMsg {
				continue
			}

			msgExecuteContract := new(struct {
				Transfer struct {
					Amount    string `json:"amount"`
					Recipient string `json:"recipient"`
				} `json:"transfer"`
			})

			decodedMsgExecuteContract, decodeErr := base64.StdEncoding.DecodeString(msgPayload.ExecuteMsg)
			if decodeErr != nil {
				return decodeErr
			}

			utils2.MustUnmarshal(decodedMsgExecuteContract, msgExecuteContract)

			// skip if this executeMsg does NOT contain transfer
			if utils2.IsZero(msgExecuteContract.Transfer) {
				continue
			}

			// skip if this executeMsg has transfer,
			// but does not comply with CW20 spec
			var isValidCW20Transfer = !utils2.IsZero(msgExecuteContract.Transfer.Amount) && !utils2.IsZero(msgExecuteContract.Transfer.Recipient)
			if !isValidCW20Transfer {
				continue
			}

			// create payloads
			// note that we need to create 2 CW20Transfer payloads for each transfer,
			// one for send and one for receive
			commitTargets = append(commitTargets, CW20Transfer{
				Height:       tx.Height,
				Sender:       msgPayload.Sender,
				Recipient:    msgExecuteContract.Transfer.Recipient,
				Contract:     msgPayload.Contract,
				Address:      msgPayload.Sender,
				TransferType: "send",
				Amount:       msgExecuteContract.Transfer.Amount,
			})

			commitTargets = append(commitTargets, CW20Transfer{
				Height:       tx.Height,
				Sender:       msgPayload.Sender,
				Recipient:    msgExecuteContract.Transfer.Recipient,
				Contract:     msgPayload.Contract,
				Address:      msgExecuteContract.Transfer.Recipient,
				TransferType: "receive",
				Amount:       msgExecuteContract.Transfer.Amount,
			})
		}
	}

	if commitErr := commit(commitTargets); commitErr != nil {
		return commitErr
	}

	return nil
}
