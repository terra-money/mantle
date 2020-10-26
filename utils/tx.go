package utils

import "github.com/terra-project/mantle-sdk/types"

func MustPass(blockState *types.BlockState, e error) *types.BlockState {
	if e != nil { panic(e) }
	return blockState
}
