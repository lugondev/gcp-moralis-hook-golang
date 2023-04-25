package moralis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo/v4"
	"github.com/thoas/go-funk"
	"log"
	"math"
	"math/big"
	"moralis-webhook/notifier"
	"net/http"
	"strings"
)

//const privateKeyHoney = "3d153b43d2b05ed7cbdd4262ec4600bb8def570421d97d73dd59d00b4584be0c"
//const honeyAddress = "0xd453def0b97911be60d0899c62c445e6c4096582"

const privateKeyHoney = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
const honeyAddress = "0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266"
const toAddress = "0x01504761F5Ec308Fc0BAf3e705f31F2466535d94"
const honeyToken = "0x912ce59144191c1204e64559fe8253a0e49e6548"

var initBalance = big.NewInt(0)
var initNonce = uint64(0)

var (
	gasLimit = big.NewInt(250000)
	gasPrice = big.NewInt(100000000)
	fee      = big.NewInt(25000000000000)
)

func Hook(client *ethclient.Client, notifierClient notifier.Notifier) func(c echo.Context) error {
	fmt.Println("hook called")

	initBalance, _ = client.BalanceAt(context.Background(), common.HexToAddress(honeyAddress), nil)
	initNonce, _ = client.NonceAt(context.Background(), common.HexToAddress(honeyAddress), nil)
	return func(c echo.Context) error {
		//contract := strings.ToLower(c.Param("contract"))
		//if !common.IsHexAddress(contract) {
		//	return echo.NewHTTPError(http.StatusBadRequest, "Invalid contract address")
		//}

		//contractByAddress, err := store.GetContractByAddress(context.Background(), contract)
		//if err != nil {
		//	log.Println("unable to get contract by address", err)
		//	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		//}

		payload := new(Payload)
		if err := c.Bind(payload); err != nil {
			log.Printf("unable to bind payload: %v\n", err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, payload)
		if payload.ChainId == "" {
			log.Println("chainId is empty")
			return c.JSON(http.StatusOK, payload)
		}

		//chain, err := eth.GetChainInfoByHexId(payload.ChainId)
		//if err != nil {
		//	log.Printf("unable to get chain data: %v\n", err)
		//	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		//}
		//listRecipients := []email.Recipient{
		//	{Name: "Lugon", Email: "lugon@alphatrue.com"},
		//	//{Name: "Dev", Email: "lugondev@gmail.com"},
		//}

		if !payload.Confirmed {
			//marshalIndent, _ := json.MarshalIndent(payload, "", "    ")
			//if err := notifierClient.Notify(notifier.Message{
			//	Content: string(marshalIndent),
			//}); err != nil {
			//	log.Println("unable to notify", err)
			//}
			//
			//if contractByAddress.Notification == "disable" {
			//	if len(payload.Erc20Transfers) > 0 {
			//		for _, transferEvent := range payload.Erc20Transfers {
			//			if err := addTokenToContract(store, contract, transferEvent); err != nil {
			//				log.Println("unable to add token to contract", err)
			//			}
			//		}
			//	}
			//	return c.JSON(http.StatusOK, payload)
			//}

			//i, _ := strconv.ParseInt(payload.Block.Timestamp, 10, 64)
			//tm := time.Unix(i, 0)

			//message := email.Message{
			//	Subject:    "Transaction Notification",
			//	Sender:     email.Sender{Name: "Notifier", Email: "notifier@alphatrue.dev"},
			//	Recipients: &listRecipients,
			//}
			//emailInfo := template.AppInfo{
			//	Contract: contract,
			//	Network:  chain.Name,
			//	User:     "Multisig User",
			//	App:      payload.Tag,
			//	Date:     tm.Format(time.UnixDate),
			//}
			//if len(chain.Explorers) > 0 {
			//	emailInfo.ExplorerUrl += chain.Explorers[0].Url + "/tx/"
			//}
			//
			//_ = handleErc20Transfers(store, emailClient, contract, payload, &emailInfo, message)
			//_ = handleTxs(emailClient, payload, &emailInfo, message)
		} else {
			//marshalIndent, _ := json.MarshalIndent(payload, "", "    ")
			//if err := notifierClient.Notify(notifier.Message{
			//	Content: string(marshalIndent),
			//}); err != nil {
			//	log.Println("unable to notify", err)
			//}

			return c.JSON(http.StatusOK, payload)
		}

		//_ = handleErc20Transfers(notifierClient, payload)
		err := handleTxs(context.Background(), client, notifierClient, payload)
		if err != nil {
			notifierClient.Notify(notifier.Message{
				Content: "handleTxs error: " + err.Error(),
			})
		}

		//marshalIndent, _ := json.MarshalIndent(payload, "", "    ")
		//if err := notifierClient.Notify(notifier.Message{
		//	Content: string(marshalIndent),
		//}); err != nil {
		//	log.Println("unable to notify", err)
		//}
		return c.JSON(http.StatusOK, payload)
	}
}

func handleTxs(ctx context.Context, client *ethclient.Client, notifierClient notifier.Notifier, payload *Payload) error {
	if len(payload.Txs) > 0 {
		for _, tx := range payload.Txs {
			if tx.FromAddress != honeyAddress && tx.ToAddress != honeyAddress {
				continue
			}
			fbalance, _, _ := big.ParseFloat(tx.Value, 10, 64, big.ToNearestEven)
			ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
			txInfo := map[string]interface{}{
				"Tx":    tx.Hash,
				"From":  tx.FromAddress,
				"To":    tx.ToAddress,
				"Value": ethValue.String(),
				"Nonce": tx.Nonce,
			}

			balance, _ := new(big.Int).SetString(tx.Value, 10)
			if strings.HasPrefix(tx.Input, "0x4e71d92d") {
				txInfo["Action"] = "claim"
			}
			if strings.HasPrefix(tx.Input, "0xa9059cbb") {
				txInfo["Action"] = "transfer"
			}
			if tx.FromAddress == honeyAddress {
				txInfo["Direction"] = "out"
				initBalance = initBalance.Sub(initBalance, balance)
				txGas, _ := new(big.Int).SetString(tx.Gas, 10)
				txGasPrice, _ := new(big.Int).SetString(tx.GasPrice, 10)
				txFee := new(big.Int).Mul(txGas, txGasPrice)
				initBalance = initBalance.Sub(initBalance, txFee)
				initNonce++
			}

			initBalanceFloat, _ := new(big.Float).SetString(initBalance.String())
			initValue := new(big.Float).Quo(initBalanceFloat, big.NewFloat(math.Pow10(18)))
			txInfo["InitBalance"] = initValue.String()
			txInfo["InitNonce"] = initNonce

			if tx.ToAddress == honeyAddress && tx.Value != "0" {
				txInfo["Direction"] = "in"
				initBalance = initBalance.Add(initBalance, balance)
				// convert string to big.Int
				//value, _ := new(big.Int).SetString(tx.Value, 10)

				// routine golang run 3 times in background
				go func() {
					//value, err := client.BalanceAt(ctx, common.HexToAddress(honeyAddress), nil)
					//if err != nil {
					//	_ = notifierClient.Notify(notifier.Message{
					//		Content: fmt.Sprintf("unable to get balance: %v", err),
					//	})
					//	return
					//}
					value := new(big.Int).Set(initBalance)
					fmt.Println("value BalanceAt:", value.String())
					value = value.Sub(value, fee)
					fmt.Println("value remaining", value.String())

					//currentNonce, err := client.NonceAt(ctx, common.HexToAddress(honeyAddress), nil)
					//if err != nil {
					//	_ = notifierClient.Notify(notifier.Message{
					//		Content: fmt.Sprintf("unable to get nonce: %v", err),
					//	})
					//	return
					//}
					//fmt.Println("currentNonce", currentNonce)
					fmt.Println("initNonce", initNonce)
					arr := []uint64{0, 1, 2}
					funk.ForEach(arr, func(i uint64) {
						go func(i uint64) {
							err := transfer(ctx, client, common.HexToAddress(toAddress), initNonce+i, value)
							if err != nil {
								_ = notifierClient.Notify(notifier.Message{
									Content: err.Error(),
								})
							}
						}(i)
					})
				}()
			}

			marshalIndent, _ := json.MarshalIndent(txInfo, "", "    ")
			if err := notifierClient.Notify(notifier.Message{
				Content: string(marshalIndent),
			}); err != nil {
				log.Println("unable to notify", err)
			}
		}
	}
	return nil
}

//func handleErc20Transfers(notifierClient notifier.Notifier, payload *Payload) error {
//	if len(payload.Erc20Transfers) > 0 {
//		for _, transferEvent := range payload.Erc20Transfers {
//			if transferEvent.Contract != honeyToken {
//				continue
//			}
//			transferInfo := map[string]interface{}{
//				"Tx":          transferEvent.TransactionHash,
//				"From":        transferEvent.From,
//				"To":          transferEvent.To,
//				"Amount":      transferEvent.ValueWithDecimals,
//				"Contract":    transferEvent.Contract,
//				"TokenName":   transferEvent.TokenName,
//				"TokenSymbol": transferEvent.TokenSymbol,
//			}
//
//			marshalIndent, _ := json.MarshalIndent(transferInfo, "", "    ")
//			if err := notifierClient.Notify(notifier.Message{
//				Content: string(marshalIndent),
//			}); err != nil {
//				log.Println("unable to notify", err)
//			}
//		}
//	}
//	return nil
//}

//func handleTxs(emailClient email.Client, payload *Payload, emailInfo *template.AppInfo, message email.Message) error {
//	if len(payload.Txs) > 0 {
//		for _, tx := range payload.Txs {
//
//			txInfo := template.ContractCall{
//				Tx:   tx.Hash,
//				From: tx.FromAddress,
//				To:   tx.ToAddress,
//				Data: tx.Input,
//			}
//
//			inputDecoded, err := eth.InputDataDecoder(tx.Input)
//			if err != nil {
//				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
//			}
//
//			message.HtmlContent = template.GenerateEmail(template.TransactionDetail, emailInfo, txInfo, inputDecoded)
//			message.TextContent = template.GenerateEmail(template.TransactionDetail, emailInfo, txInfo, inputDecoded)
//
//			if err := emailClient.Send(message); err != nil {
//				return err
//			}
//		}
//	}
//	return nil
//}
//
//func handleErc20Transfers(store *db.SQLStore, emailClient email.Client, contract string, payload *Payload, emailInfo *template.AppInfo, message email.Message) error {
//	if len(payload.Erc20Transfers) > 0 {
//		for _, transferEvent := range payload.Erc20Transfers {
//			transferInfo := template.Erc20Transfer{
//				Tx:          transferEvent.TransactionHash,
//				From:        transferEvent.From,
//				To:          transferEvent.To,
//				Amount:      transferEvent.ValueWithDecimals,
//				Contract:    transferEvent.Contract,
//				TokenName:   transferEvent.TokenName,
//				TokenSymbol: transferEvent.TokenSymbol,
//			}
//			if err := addTokenToContract(store, contract, transferEvent); err != nil {
//				log.Println("unable to add token to contract", err)
//			}
//			message.HtmlContent = template.GenerateEmail(template.TransactionToken, emailInfo, transferInfo)
//			message.TextContent = template.GenerateEmail(template.TransactionToken, emailInfo, transferInfo)
//
//			if err := emailClient.Send(message); err != nil {
//				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
//			}
//		}
//	}
//	return nil
//}

//func addTokenToContract(store *db.SQLStore, contract string, transfer Erc20Transfer) error {
//	decimals, _ := strconv.Atoi(transfer.TokenDecimals)
//	params := sqlc.AddTokenToContractByAddressParams{
//		Address:  contract,
//		Name:     transfer.TokenName,
//		Symbol:   transfer.TokenSymbol,
//		Token:    transfer.Contract,
//		Decimals: int64(decimals),
//	}
//
//	if _, err := store.AddTokenToContractByAddress(context.Background(), params); err != nil {
//		return err
//	}
//	return nil
//}

func transfer(ctx context.Context, client *ethclient.Client, to common.Address, currentNonce uint64, value *big.Int) error {

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    currentNonce,
		GasPrice: gasPrice,
		Gas:      gasLimit.Uint64(),
		To:       &to,
		Value:    value,
	})

	signer := types.LatestSignerForChainID(big.NewInt(42161))
	// ecdsa private key from string
	privateKey, err := crypto.HexToECDSA(privateKeyHoney)
	if err != nil {
		return err
	}
	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		return err
	}
	return client.SendTransaction(ctx, signedTx)
}
