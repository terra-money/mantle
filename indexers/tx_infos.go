package indexers

import (
	"encoding/base64"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
	lutils "github.com/terra-project/mantle-official/utils"
	. "github.com/terra-project/mantle/types"
	"github.com/terra-project/mantle/utils"
	"reflect"
	"time"
)

// mantle-specific tx type
type TxInfo struct {
	Height       uint64 // already indexed
	TxHash       string `model:"index"`
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
	Success bool
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
	Amount sdk.Coins `json:"amount" yaml:"amount"`
	Gas    uint64
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
		BaseState struct {
			Height int64
			Block  struct {
				Header struct {
					Time string
				}
			}
			Txs                []LazyTx
			DeliverTxResponses []ResponseDeliverTx
		}
	})

	if queryErr := q(request, nil); queryErr != nil {
		return fmt.Errorf("fetching txs failed, err=%s", queryErr)
	}

	// transform txs into mantleTx
	txs := request.BaseState.Txs
	txResults := request.BaseState.DeliverTxResponses
	var commitTarget = make(TxInfos, len(txs))

	for txIndex, tx := range txs {
		txBytes, txBytesErr := base64.StdEncoding.DecodeString(tx.TxString)
		if txBytesErr != nil {
			return fmt.Errorf("tx byte decode failed, err=%s", txBytesErr)
		}
		txHash := tmhash.Sum(txBytes)
		txResult := txResults[txIndex]
		txdoc := tx.Decode()
		timeInUint64, _ := time.Parse(time.RFC3339, request.BaseState.Block.Header.Time)

		// log -> TxxInfoLog
		rawLogParsed := new([]TxInfoLog)
		utils.MustUnmarshal([]byte(txResult.Log), rawLogParsed)

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
				Value: NewJSONScalar(m, lutils.DecodeWasm),
			}
		}

		// txdoc.Fee ->
		var fee = TxInfoStdFee{
			Amount: txdoc.Fee.Amount,
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
			Height:       uint64(request.BaseState.Height),
			TxHash:       fmt.Sprintf("%X", txHash),
			RawLog:       txResult.Log,
			Logs:         *rawLogParsed,
			GasWanted:    uint64(txResult.GasWanted),
			GasUsed:      uint64(txResult.GasUsed),
			Timestamp:    request.BaseState.Block.Header.Time,
			TimestampUTC: uint64(timeInUint64.UnixNano()),
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
