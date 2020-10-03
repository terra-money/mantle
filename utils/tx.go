package utils

import "github.com/terra-project/mantle/types"

func MustPass(result types.BaseState) types.BaseState {
	for _, result := range result.DeliverTxResponses {
		if result.IsErr() || !result.IsOK() {
			panic(result.Log)
		}
	}

	return result
}