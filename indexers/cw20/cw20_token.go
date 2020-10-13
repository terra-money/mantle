package cw20

import (
	"github.com/terra-project/mantle/types"
	"reflect"
)

type CW20Token struct {
	Contract    string `model:"index"`
	Denom       string `model:"index"`
	TotalSupply string
	Holders     uint64
	Deicimals   uint8
	Transfers   uint64
}

type CW20Tokens []CW20Token

func RegisterCW20Token(register types.Register) {
	register(
		IndexCW20Token,
		reflect.TypeOf((*CW20Token)(nil)),
	)
}

func IndexCW20Token(query types.Query, commit types.Commit) error {
	//var commitTargets CW20Tokens
	return nil
}

func indexCW20Instantiate() CW20Tokens {
	return nil
}


func indexCW20Transfers() CW20Tokens {
	return nil
}