package blocks

import (
	"github.com/terra-money/mantle-sdk/types"
	"github.com/terra-money/mantle/indexers/tx_infos"
	"reflect"
)

// Block is a handy alias to BlockState.Block,
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
		BlockState struct {
			Block types.RawBlock
		}
		TxInfos tx_infos.TxInfos
	})

	if queryErr := query(queryBlock, nil); queryErr != nil {
		return queryErr
	}

	//
	var commitTarget = Block{}
	commitTarget.Height = uint64(queryBlock.BlockState.Block.Header.Height)
	commitTarget.Header = queryBlock.BlockState.Block.Header
	commitTarget.Txs = queryBlock.TxInfos

	var asSlice = Blocks{commitTarget}
	if commitErr := commit(asSlice); commitErr != nil {
		return commitErr
	}

	return nil
}
