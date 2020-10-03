package txs

import (
	"encoding/base64"
	"fmt"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"github.com/terra-project/core/x/auth"
	lutils "github.com/terra-project/mantle-official/utils"
	. "github.com/terra-project/mantle/types"
	"reflect"
	"time"
)

// mantle-specific tx type
type mantleTx struct {
	Hash       string `model:"index"`
	TxString   string
	HeightAt   uint64 `model:"index"`
	Timestamp  int64  `model:"index"`
	Success    bool
	Fee        auth.StdFee
	Signatures []mantleSignature
	Msg        []mantleTxMsg
	Memo       string
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
			mmsg[j].Value = NewJSONScalar(msg, lutils.DecodeWasm)
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
