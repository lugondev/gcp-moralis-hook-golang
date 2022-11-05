package eth

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Method string

const (
	SubmitTx  Method = "0xbe616434"
	ExecuteTx Method = "0xee22610b"
	CancelTx  Method = "0x3380c0d8"
	RevokeTx  Method = "0x735631ad"
	ConfirmTx Method = "0xc01a8c84"
)

func ByteToMethod(input []byte) Method {
	return Method(hexutil.Encode(input))
}

type (
	MultiSigCall struct {
		ID           int64                  `json:"id"`
		Caller       string                 `json:"caller"`
		MethodId     Method                 `json:"methodId"`
		MethodName   string                 `json:"methodName"`
		InputDecoder map[string]interface{} `json:"input"`
	}
)
