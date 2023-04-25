package routes

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"math/big"
	evm_signer "moralis-webhook/evm-signer"
	"net/http"
	"strings"
)

func (r *router) getChain(c echo.Context) error {
	chainParam := c.Param("chainId")
	chainId, ok := big.NewInt(0).SetString(chainParam, 10)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid chainId")
	}
	chain, err := r.SQLStore.GetChain(context.Background(), chainId.Int64())
	if err != nil || chain.ID == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "chain not found")
	}
	return c.JSON(http.StatusOK, chain)
}

func (r *router) listChains(c echo.Context) error {
	chains, err := r.SQLStore.ListChains(context.Background())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, chains)
}

func (r *router) listAddresses(c echo.Context) error {
	addresses, err := r.SQLStore.ListAddresses(context.Background())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, addresses)
}

func (r *router) signAndTransact(c echo.Context) error {
	fromAddress := c.Param("address")
	if !common.IsHexAddress(fromAddress) {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"message": "address is invalid",
		})
	}
	fromAddressData, err := r.SQLStore.GetPrivateKey(context.Background(), fromAddress)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var txRequest evm_signer.TxRequest
	if err := c.Bind(&txRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	fmt.Println("txRequest nonce:", txRequest.Nonce)
	fmt.Println("txRequest gasPrice:", txRequest.GasPrice)
	fmt.Println("txRequest gasLimit:", txRequest.GasLimit)
	fmt.Println("txRequest data:", txRequest.Data)
	fmt.Println("txRequest value:", txRequest.Value)

	if txRequest.GasLimit == nil {
		txRequest.GasLimit = big.NewInt(500000)
	}
	if len(txRequest.Data) > 0 {
		if strings.Contains(txRequest.Data, "0x") {
			txRequest.Data = strings.ReplaceAll(txRequest.Data, "0x", "")
		}
		_, err := hex.DecodeString(txRequest.Data)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "data is invalid hex")
		}
	}

	if txRequest.Value != "0" && txRequest.Value != "0x0" && txRequest.Value != "0x" {
		v := big.NewInt(0).SetBytes(common.FromHex(txRequest.Value))
		if v.Uint64() == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "value is invalid")
		}
		txRequest.Value = common.Bytes2Hex(v.Bytes())
	}

	if txRequest.To == common.HexToAddress("0x0000000000000000000000000000000000000000") {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid to address")
	}

	chain, err := r.SQLStore.GetChain(context.Background(), txRequest.ChainId)
	if err != nil || chain.ID == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "chain not found")
	}

	transaction, err := evm_signer.SignAndTransact(txRequest, fromAddressData, chain)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, transaction)
}
