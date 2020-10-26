package utils

import (
	"github.com/terra-project/mantle-sdk/types"
	"io/ioutil"
	"path/filepath"
)

func GetWasmBytes(p string) (wasmBytes []byte) {
	filename, _ := filepath.Abs(p)
	var wasmBytesErr error
	if wasmBytes, wasmBytesErr = ioutil.ReadFile(filename); wasmBytesErr != nil {
		panic(wasmBytesErr)
	}

	return
}


func GetContractAddressFromInstantiate(result *types.BlockState, _ error) string {
	var addr string
	for _, event := range result.ResponseDeliverTx[0].Events {
		switch event.Type {
		case "instantiate_contract":
			for _, attr := range event.Attributes {
				switch string(attr.Key) {
				case "contract_address":
					addr = string(attr.Value)
				}
			}
		}
	}

	return addr
}
