package txs

import (
	"fmt"
	"github.com/terra-project/core/x/auth"
	lutils "github.com/terra-project/mantle-official/utils"
	. "github.com/terra-project/mantle/types"
	"reflect"
	"time"
)

// Mantle-specific tx type
type MantleTx struct {
	Hash       string `model:"index"`
	TxString   string
	HeightAt   uint64 `model:"index"`
	Timestamp  int64  `model:"index"`
	Success    bool
	Fee        auth.StdFee
	Signatures []MantleSignature
	Msg        []MantleTxMsg
	Memo       string
	Result     ResponseDeliverTx
}

type MantleTxMsg struct {
	Route string
	Type  string
	Value JSONScalar
}

type MantleSignature struct {
	PubKey    MantleSignaturePubKey
	Signature string
}

type MantleSignaturePubKey struct {
	Type  string
	Value string
}

type Txs []MantleTx

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
					Time time.Time
				}
			}
			Txs                []Tx
			DeliverTxResponses []ResponseDeliverTx
		}
	})

	if queryErr := q(request, nil); queryErr != nil {
		return fmt.Errorf("fetching txs failed, err=%s", queryErr)
	}

	// transform txs into MantleTx
	txs := request.BaseState.Txs
	txResults := request.BaseState.DeliverTxResponses
	var commitTarget = make(Txs, len(txs))

	for txIndex, tx := range txs {
		txHash := tx.Hash()
		txdoc, err := TxDecoder(tx)
		if err != nil {
			return err
		}

		// recreate msgs
		mmsg := make([]MantleTxMsg, len(txdoc.Msgs))

		for j, msg := range txdoc.Msgs {
			mmsg[j].Type = msg.Type()
			mmsg[j].Route = msg.Route()
			mmsg[j].Value = NewJSONScalar(msg, lutils.DecodeWasm)
		}

		signatures := make([]MantleSignature, len(txdoc.Signatures))
		for j, signature := range txdoc.Signatures {
			signatures[j] = MantleSignature{
				PubKey: MantleSignaturePubKey{
					Type:  "tendermint/PubKeySecp256k1",
					Value: string(signature.PubKey.Bytes()),
				},
				Signature: string(signature.Signature),
			}
		}

		commitTarget[txIndex] = MantleTx{
			Hash:       fmt.Sprintf("%X", txHash),
			TxString:   tx.String(),
			Timestamp:  request.BaseState.Block.Header.Time.UnixNano(),
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
