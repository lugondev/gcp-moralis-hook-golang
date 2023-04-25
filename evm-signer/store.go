package evm_signer

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"math/rand"
	db "moralis-webhook/db/sqlc"
	"moralis-webhook/eth"
)

type TxRequest struct {
	Nonce    *big.Int       `json:"nonce"`
	GasPrice *big.Int       `json:"gasPrice"`
	GasLimit *big.Int       `json:"gasLimit"`
	Value    string         `json:"value"`
	ChainId  int64          `json:"chainId"`
	To       common.Address `json:"to"`
	Data     string         `json:"data"`
}

func random(rpcs []string) string {
	if len(rpcs) == 1 {
		return rpcs[0]
	}

	randomIndex := rand.Intn(len(rpcs))
	return rpcs[randomIndex]
}

func GetEthClient(chain db.Chain) (*ethclient.Client, error) {
	rpc := random(chain.Rpcs)
	return ethclient.Dial(rpc)
}

func SignAndTransact(txRequest TxRequest, data db.PrivateKey, chain db.Chain) (*types.Transaction, error) {
	client, err := GetEthClient(chain)
	if err != nil {
		return nil, err
	}
	chainId, _ := client.ChainID(context.Background())
	fmt.Println("chainId", chainId)
	privateKey, address := eth.WalletFromPrivate(data.PrivateKey)
	if *address != common.HexToAddress(data.Address) {
		return nil, errors.New("private key and address do not match")
	}

	opt := &bind.TransactOpts{
		From: common.HexToAddress(data.Address),
		Signer: func(addr common.Address, txn *types.Transaction) (*types.Transaction, error) {
			signer := types.NewEIP155Signer(chainId)
			txHash := signer.Hash(txn)
			fmt.Println("txHash", common.Bytes2Hex(txHash.Bytes()))
			sig, err := crypto.Sign(txHash.Bytes()[:], privateKey)
			if err != nil {
				fmt.Println("sign error", err)
				return nil, err
			}

			for j := 0; j < 2; j++ {
				signedTxn, err := txn.WithSignature(signer, sig)

				if err != nil {
					fmt.Println("Error with signature txn", "detail", err)
					return nil, err
				}
				sender, err := types.Sender(signer, signedTxn)
				if sender.String() == addr.String() {
					return signedTxn, nil
				}

				fmt.Println("sender", "addr", sender, "require", addr)
				vPos := crypto.SignatureLength - 1
				sig[vPos] ^= 0x1
			}
			return nil, errors.New("wrong sender address")
		},

		Nonce:    txRequest.Nonce,
		GasPrice: txRequest.GasPrice,
		GasLimit: txRequest.GasLimit.Uint64(),
		Value:    big.NewInt(0).SetBytes(common.FromHex(txRequest.Value)),
	}
	bound := bind.NewBoundContract(txRequest.To, abi.ABI{}, client, client, client)
	transaction, err := bound.RawTransact(opt, common.Hex2Bytes(txRequest.Data))

	if err != nil {
		return nil, err
	}

	return transaction, nil
}
