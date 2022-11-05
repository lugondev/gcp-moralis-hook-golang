package eth

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

func unpackValues(input string) ([]interface{}, Method, error) {
	mutlisigInterface, _ := MultisigMetaData.GetAbi()
	decodedMethod, _ := hexutil.Decode(input[:10])
	methodById, err := mutlisigInterface.MethodById(decodedMethod)
	if err != nil {
		return nil, "", err
	}
	decodedInput, _ := hexutil.Decode(input)
	unpackValues, err := methodById.Inputs.UnpackValues(decodedInput[4:])
	if err != nil {
		return nil, "", err
	}

	return unpackValues, ByteToMethod(methodById.ID), nil
}

func InputDataDecoder(input string) (interface{}, error) {
	unpackValues, methodId, err := unpackValues(input)
	if err != nil {
		return nil, err
	}

	dataTransaction := &MultiSigCall{
		MethodId: methodId,
	}
	if methodId != SubmitTx {
		dataTransaction.ID = unpackValues[0].(*big.Int).Int64()
		dataTransaction.InputDecoder = map[string]interface{}{
			"id": unpackValues[0].(*big.Int).Int64(),
		}
	}
	switch methodId {
	case CancelTx:
		dataTransaction.MethodName = "Cancel Transaction"
	case ConfirmTx:
		dataTransaction.MethodName = "Confirm Transaction"
	case ExecuteTx:
		dataTransaction.MethodName = "Execute Transaction"
	case RevokeTx:
		dataTransaction.MethodName = "Revoke Transaction"
	default:
		dataTransaction.MethodName = "Submit Transaction"
		dataTransaction.InputDecoder = map[string]interface{}{
			"target":   unpackValues[0].(common.Address).String(),
			"data":     hexutil.Encode(unpackValues[1].([]byte)),
			"deadline": unpackValues[2].(*big.Int).Int64(),
			"note":     unpackValues[3].(string),
		}
	}

	return dataTransaction, nil
}
