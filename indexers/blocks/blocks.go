package blocks

import (
	"github.com/terra-project/mantle-official/indexers/tx_infos"
	"github.com/terra-project/mantle-sdk/types"
	"reflect"
)

// Block is a handy alias to BaseState.Block,
// where it only contains Header and Data.Txs
type Block struct {
	Height uint64
	Header types.Header
	Txs    tx_infos.TxInfos
}

type Blocks []Block

func RegisterBlocks(register types.Register) {
	register(
		IndexBlocks,
		reflect.TypeOf((*Blocks)(nil)),
	)
}

func IndexBlocks(query types.Query, commit types.Commit) error {
	queryBlock := new(struct {
		BaseState struct {
			Block types.Block
		}
		TxInfos tx_infos.TxInfos
	})

	if queryErr := query(queryBlock, nil); queryErr != nil {
		return queryErr
	}

	//
	var commitTarget = Block{}
	commitTarget.Height = uint64(queryBlock.BaseState.Block.Header.Height)
	commitTarget.Header = queryBlock.BaseState.Block.Header
	commitTarget.Txs = queryBlock.TxInfos

	var asSlice = Blocks{commitTarget}
	if commitErr := commit(asSlice); commitErr != nil {
		return commitErr
	}

	return nil
}