package utils

import (
	"encoding/base64"
	"encoding/json"
	"github.com/terra-project/core/x/wasm"
	"reflect"
)

func IsZero(data interface{}) bool {
	return reflect.Indirect(reflect.ValueOf(data)).IsZero()
}

func DecodeWasm(oMsg interface{}, data []byte) []byte {
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
