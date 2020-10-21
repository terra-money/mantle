package utils

import "github.com/terra-project/mantle-sdk/types"

func MustPass(result types.BlockState) types.BlockState {
	for _, result := range result.ResponseDeliverTx {
		if result.IsErr() || !result.IsOK() {
			panic(result.Log)
		}
	}

	return result
}