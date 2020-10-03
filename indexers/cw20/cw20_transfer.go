package cw20

import (
	"encoding/base64"
	"github.com/terra-project/mantle-official/indexers"
	utils2 "github.com/terra-project/mantle-official/utils"
	"github.com/terra-project/mantle/types"
	"github.com/terra-project/mantle/utils"
	"reflect"
)

type CW20Transfer struct {
	Height uint64
	Sender string `model:"index"`
	Recipient string `model:"index"`
	Contract string `model:"index"`
	Address string `model:"index"`
	TransferType string `model:"index"` // send|receive
	Amount string `model:"index"`// wasm:uint128
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
		TxInfos indexers.TxInfos
	})

	if queryErr := query(txInfosQuery, nil); queryErr != nil {
		return queryErr
	}

	// iterate over all txInfos, and extract CW20-like transfers
	var commitTargets CW20Transfers
	for _, tx := range txInfosQuery.TxInfos {
		// skip if this tx is not successful
		if !tx.Success { continue }

		for _, msg := range tx.Tx.Msg {
			// skip if not wasm/MsgExecuteContract
			if msg.Type != "wasm/MsgExecuteContract" { continue }

			// check if MsgExecuteContract has "transfer" key
			msgPayload := new(struct {
				sender string `json:"sender"`
				contract string `json:"contract"`
				executeMsg string `json:"execute_msg"`
			})

			utils.MustUnmarshal([]byte(msg.Value), msgPayload)

			var isValidExecuteMsg = !utils2.IsZero(msgPayload.sender) && !utils2.IsZero(msgPayload.contract) && !tils2.IsZero(msgPayload.executeMsg)
			if !isValidExecuteMsg {
				continue
			}

			// skip if this executeMsg does NOT contain executeMsg
			if utils2.IsZero(msgPayload.executeMsg) {
				continue
			}

			msgExecuteContract := new(struct {
				transfer struct {
					amount string `json:"amount"`
					recipient string `json:"recipient"`
				} `json:"transfer"`
			})
			decodedMsgExecuteContract, decodeErr := base64.StdEncoding.DecodeString(msgPayload.executeMsg)
			if decodeErr != nil {
				return decodeErr
			}
			utils.MustUnmarshal([]byte(msgPayload.executeMsg), decodedMsgExecuteContract)

			// skip if this executeMsg does NOT contain transfer
			if utils2.IsZero(msgExecuteContract.transfer) {
				continue
			}

			// skip if this executeMsg has transfer,
			// but does not comply with CW20 spec
			var isValidCW20Transfer = !utils2.IsZero(msgExecuteContract.transfer.amount) && !utils2.IsZero(msgExecuteContract.transfer.recipient)
			if !isValidCW20Transfer {
				continue
			}

			// create payloads
			// note that we need to create 2 CW20Transfer payloads for each transfer,
			// one for send and one for receive
			commitTargets = append(commitTargets, CW20Transfer{
				Height:       tx.Height,
				Sender:       msgPayload.sender,
				Recipient:    msgExecuteContract.transfer.recipient,
				Contract:     msgPayload.contract,
				Address:      msgPayload.sender,
				TransferType: "send",
				Amount:       msgExecuteContract.transfer.amount,
			})

			commitTargets = append(commitTargets, CW20Transfer{
				Height:       tx.Height,
				Sender:       msgPayload.sender,
				Recipient:    msgExecuteContract.transfer.recipient,
				Contract:     msgPayload.contract,
				Address:      msgExecuteContract.transfer.recipient,
				TransferType: "receive",
				Amount:       msgExecuteContract.transfer.amount,
			})
		}
	}

	if commitErr := commit(commitTargets); commitErr != nil {
		return commitErr
	}

	return nil
}
