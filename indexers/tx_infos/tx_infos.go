package tx_infos

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	. "github.com/terra-project/mantle-sdk/types"
	lutils "github.com/terra-project/mantle/utils"
)

// mantle-specific tx type
type TxInfo struct {
	Height       uint64 // already indexed
	TxHash       string `model:"index,primary"`
	RawLog       string
	Logs         []TxInfoLog
	GasWanted    uint64
	GasUsed      uint64
	Timestamp    string
	TimestampUTC uint64 `model:"index"`
	Events       []TxInfoEvent
	Code         uint64
	Tx           TxInfoStdTx

	// custom
	Success bool `model:"index"`
}

type TxInfoStdTx struct {
	Msg        []TxInfoStdTxMsg
	Fee        TxInfoStdFee
	Signatures []TxInfoStdTxSignature
	Memo       string `model:"index"`
}

type TxInfoStdTxSignature struct {
	PubKey    TxInfoStdTxSignaturePubKey
	Signature []byte
}

type TxInfoStdTxSignaturePubKey struct {
	Type  string
	Value []byte
}

type TxInfoStdFee struct {
	Amount []TxInfoStdFeeAmount
	Gas    uint64
}

type TxInfoStdFeeAmount struct {
	Denom  string
	Amount string
}

type TxInfoLog struct {
	MsgIndex uint64        `json:"msg_index"`
	Log      string        `json:"log"`
	Events   []TxInfoEvent `json:"events"`
}

type TxInfoEvent struct {
	Type       string            `json:"type"`
	Attributes []TxInfoAttribute `json:"attributes"`
}

type TxInfoAttribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type TxInfoStdTxMsg struct {
	Type  string
	Value JSONScalar
}

type TxInfos []TxInfo

func RegisterTxInfos(register Register) {
	register(
		IndexTxInfos,
		reflect.TypeOf((*TxInfos)(nil)),
	)
}

func IndexTxInfos(q Query, c Commit) error {
	request := new(struct {
		BlockState struct {
			Height int64
			Block  struct {
				Header struct {
					Time time.Time
				}
				Data struct {
					Txs []Tx
				}
			}
			ResponseDeliverTx []ResponseDeliverTx
		}
	})

	if queryErr := q(request, nil); queryErr != nil {
		return fmt.Errorf("fetching txs failed, err=%s", queryErr)
	}

	// transform txs into mantleTx
	txs := request.BlockState.Block.Data.Txs
	txResults := request.BlockState.ResponseDeliverTx
	var commitTarget = make(TxInfos, len(txs))

	for txIndex, tx := range txs {
		txHash := tx.Hash()
		txResult := txResults[txIndex]
		txdoc, txdocErr := TxDecoder(tx)
		if txdocErr != nil {
			return txdocErr
		}
		timeInUint64 := uint64(request.BlockState.Block.Header.Time.Unix())

		// log -> TxxInfoLog
		rawLogParsed := new([]TxInfoLog)
		if unmarshalErr := json.Unmarshal([]byte(txResult.Log), rawLogParsed); unmarshalErr != nil {
			// this means that the log is plain string.
			// leave rawLogParsed as nil
			// noop
		}

		// txResult.Events -> Events
		eventsParsed := make([]TxInfoEvent, len(txResult.Events))
		for ei, e := range txResult.Events {
			attributes := make([]TxInfoAttribute, len(e.Attributes))
			for ai, a := range e.Attributes {
				attributes[ai] = TxInfoAttribute{
					Key:   string(a.Key),
					Value: string(a.Value),
				}
			}
			eventsParsed[ei] = TxInfoEvent{
				Type:       e.Type,
				Attributes: attributes,
			}
		}

		// txdoc.Msgs -> Msg json scalars
		var msgs = make([]TxInfoStdTxMsg, len(txdoc.Msgs))
		for mi, m := range txdoc.Msgs {
			msgs[mi] = TxInfoStdTxMsg{
				Type:  lutils.MsgRouteAndTypeToString(m.Route(), m.Type()),
				Value: NewJSONScalar(m, nil),
			}
		}

		// txdoc.Fee ->
		var feeAmount = make([]TxInfoStdFeeAmount, len(txdoc.Fee.Amount))
		for amountIndex, amount := range txdoc.Fee.Amount {
			feeAmount[amountIndex] = TxInfoStdFeeAmount{
				Denom:  amount.Denom,
				Amount: amount.Amount.String(),
			}
		}
		var fee = TxInfoStdFee{
			Amount: feeAmount,
			Gas:    txdoc.Fee.Gas,
		}

		// txdoc.Signatures -> []TxInfoStdTxSignature
		var signatures = make([]TxInfoStdTxSignature, len(txdoc.Signatures))
		for si, s := range txdoc.Signatures {
			signatures[si] = TxInfoStdTxSignature{
				PubKey: TxInfoStdTxSignaturePubKey{
					Type:  "tendermint/PubKeySecp256k1",
					Value: s.Bytes(),
				},
				Signature: s.Signature,
			}
		}

		commitTarget[txIndex] = TxInfo{
			Height:       uint64(request.BlockState.Height),
			TxHash:       fmt.Sprintf("%X", txHash),
			RawLog:       txResult.Log,
			Logs:         *rawLogParsed,
			GasWanted:    uint64(txResult.GasWanted),
			GasUsed:      uint64(txResult.GasUsed),
			Timestamp:    request.BlockState.Block.Header.Time.String(),
			TimestampUTC: timeInUint64,
			Events:       eventsParsed,
			Code:         uint64(txResult.Code),
			Tx: TxInfoStdTx{
				Msg:        msgs,
				Fee:        fee,
				Signatures: signatures,
				Memo:       txdoc.Memo,
			},
			Success: txResult.IsOK() && !txResult.IsErr(),
		}
	}

	if commitErr := c(commitTarget); commitErr != nil {
		return commitErr
	}

	return nil
}