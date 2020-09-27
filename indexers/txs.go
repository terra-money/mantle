package indexers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"github.com/terra-project/core/x/auth"
	"github.com/terra-project/core/x/wasm"
	. "github.com/terra-project/mantle/types"
	"reflect"
	"time"
)

// mantle-specific tx type
type mantleTx struct {
	Hash       string `model:"index,primary"`
	TxString   string
	HeightAt   uint64 `model:"index"`
	Timestamp  int64  `model:"index"`
	Success    bool
	Fee        auth.StdFee
	Signatures []mantleSignature
	Msg        []mantleTxMsg
	Memo       string `model:"index"`
	Result     ResponseDeliverTx
}

type mantleTxMsg struct {
	Route string
	Type  string
	Value JSONScalar
}

type mantleSignature struct {
	PubKey    mantleSignaturePubKey
	Signature string
}

type mantleSignaturePubKey struct {
	Type  string
	Value string
}

type Txs []mantleTx

func RegisterTxs(register Register) {
	register(
		IndexTxs,
		reflect.TypeOf((*Txs)(nil)),
	)
}

func IndexTxs(q Query, c Commit) error {
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
	timeInUint64, timeErr := time.Parse(time.RFC3339, request.BaseState.Block.Header.Time)
	if timeErr != nil {
		return timeErr
	}
	var commitTarget = make(Txs, len(txs))

	for txIndex, tx := range txs {
		txBytes, txBytesErr := base64.StdEncoding.DecodeString(tx.TxString)
		if txBytesErr != nil {
			return fmt.Errorf("tx byte decode failed, err=%s", txBytesErr)
		}
		txHash := tmhash.Sum(txBytes)
		txdoc := tx.Decode()

		// recreate msgs
		mmsg := make([]mantleTxMsg, len(txdoc.Msgs))

		for j, msg := range txdoc.Msgs {
			mmsg[j].Type = msg.Type()
			mmsg[j].Route = msg.Route()
			mmsg[j].Value = NewJSONScalar(msg, decodeWasm)
		}

		signatures := make([]mantleSignature, len(txdoc.Signatures))
		for j, signature := range txdoc.Signatures {
			signatures[j] = mantleSignature{
				PubKey: mantleSignaturePubKey{
					Type:  "tendermint/PubKeySecp256k1",
					Value: string(signature.PubKey.Bytes()),
				},
				Signature: string(signature.Signature),
			}
		}

		commitTarget[txIndex] = mantleTx{
			Hash:       fmt.Sprintf("%X", txHash),
			TxString:   tx.TxString,
			Timestamp:  timeInUint64.Unix(),
			HeightAt:   uint64(request.BaseState.Height),
			Success:    txResults[0].IsOK() && !txResults[0].IsErr(),
			Fee:        txdoc.Fee,
			Signatures: signatures,
			Msg:        mmsg,
			Memo:       txdoc.Memo,
			Result:     txResults[txIndex],
		}
	}

	if commitErr := c(commitTarget); commitErr != nil {
		return commitErr
	}

	return nil
}

func decodeWasm(oMsg interface{}, data []byte) []byte {
	switch oMsg.(type) {
	case wasm.MsgExecuteContract:
		interimBuffer := make(map[string]interface{})
		if err := json.Unmarshal(data, &interimBuffer); err != nil {
			panic("could not unmarshal")
		}

		interimBuffer["execute_msg"], _ = base64.StdEncoding.DecodeString(interimBuffer["execute_msg"].(string))
		interimBuffer["execute_msg"] = string(interimBuffer["execute_msg"].([]byte))
		marshaled, err := json.Marshal(interimBuffer)
		if err != nil {
			panic(err)
		}

		return marshaled

	case wasm.MsgInstantiateContract:
		interimBuffer := make(map[string]interface{})
		if err := json.Unmarshal(data, &interimBuffer); err != nil {
			panic(err)
		}

		interimBuffer["init_msg"], _ = base64.StdEncoding.DecodeString(interimBuffer["init_msg"].(string))
		interimBuffer["init_msg"] = string(interimBuffer["init_msg"].([]byte))
		marshaled, err := json.Marshal(interimBuffer)
		if err != nil {
			panic(err)
		}

		return marshaled
	default:
		return data
	}


}