package moralis

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"log"
	"moralis-webhook/db"
	sqlc "moralis-webhook/db/sqlc"
	"moralis-webhook/email"
	"moralis-webhook/email/template"
	"moralis-webhook/eth"
	"moralis-webhook/notifier"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Hook(store *db.SQLStore, emailClient email.Client, notifierClient notifier.Notifier) func(c echo.Context) error {
	return func(c echo.Context) error {
		contract := strings.ToLower(c.Param("contract"))
		if !common.IsHexAddress(contract) {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid contract address")
		}

		contractByAddress, err := store.GetContractByAddress(context.Background(), contract)
		if err != nil {
			log.Println("unable to get contract by address", err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		payload := new(Payload)
		if err := c.Bind(payload); err != nil {
			log.Printf("unable to bind payload: %v\n", err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if payload.ChainId == "" {
			log.Println("chainId is empty")
			return c.JSON(http.StatusOK, payload)
		}

		chain, err := eth.GetChainInfoByHexId(payload.ChainId)
		if err != nil {
			log.Printf("unable to get chain data: %v\n", err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		listRecipients := []email.Recipient{
			{Name: "Lugon", Email: "lugon@alphatrue.com"},
			//{Name: "Dev", Email: "lugondev@gmail.com"},
		}

		if !payload.Confirmed {
			marshalIndent, _ := json.MarshalIndent(payload, "", "    ")
			if err := notifierClient.Notify(notifier.Message{
				Content: string(marshalIndent),
			}); err != nil {
				log.Println("unable to notify", err)
			}

			if contractByAddress.Notification == "disable" {
				if len(payload.Erc20Transfers) > 0 {
					for _, transferEvent := range payload.Erc20Transfers {
						if err := addTokenToContract(store, contract, transferEvent); err != nil {
							log.Println("unable to add token to contract", err)
						}
					}
				}
				return c.JSON(http.StatusOK, payload)
			}

			i, _ := strconv.ParseInt(payload.Block.Timestamp, 10, 64)
			tm := time.Unix(i, 0)

			message := email.Message{
				Subject:    "Transaction Notification",
				Sender:     email.Sender{Name: "Notifier", Email: "notifier@alphatrue.dev"},
				Recipients: &listRecipients,
			}
			emailInfo := template.AppInfo{
				Contract: contract,
				Network:  chain.Name,
				User:     "Multisig User",
				App:      payload.Tag,
				Date:     tm.Format(time.UnixDate),
			}
			if len(chain.Explorers) > 0 {
				emailInfo.ExplorerUrl += chain.Explorers[0].Url + "/tx/"
			}

			_ = handleErc20Transfers(store, emailClient, contract, payload, &emailInfo, message)
			_ = handleTxs(emailClient, payload, &emailInfo, message)
		}
		return c.JSON(http.StatusOK, payload)
	}
}

func handleTxs(emailClient email.Client, payload *Payload, emailInfo *template.AppInfo, message email.Message) error {
	if len(payload.Txs) > 0 {
		for _, tx := range payload.Txs {

			txInfo := template.ContractCall{
				Tx:   tx.Hash,
				From: tx.FromAddress,
				To:   tx.ToAddress,
				Data: tx.Input,
			}

			inputDecoded, err := eth.InputDataDecoder(tx.Input)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			message.HtmlContent = template.GenerateEmail(template.TransactionDetail, emailInfo, txInfo, inputDecoded)
			message.TextContent = template.GenerateEmail(template.TransactionDetail, emailInfo, txInfo, inputDecoded)

			if err := emailClient.Send(message); err != nil {
				return err
			}
		}
	}
	return nil
}

func handleErc20Transfers(store *db.SQLStore, emailClient email.Client, contract string, payload *Payload, emailInfo *template.AppInfo, message email.Message) error {
	if len(payload.Erc20Transfers) > 0 {
		for _, transferEvent := range payload.Erc20Transfers {
			transferInfo := template.Erc20Transfer{
				Tx:          transferEvent.TransactionHash,
				From:        transferEvent.From,
				To:          transferEvent.To,
				Amount:      transferEvent.ValueWithDecimals,
				Contract:    transferEvent.Contract,
				TokenName:   transferEvent.TokenName,
				TokenSymbol: transferEvent.TokenSymbol,
			}
			if err := addTokenToContract(store, contract, transferEvent); err != nil {
				log.Println("unable to add token to contract", err)
			}
			message.HtmlContent = template.GenerateEmail(template.TransactionToken, emailInfo, transferInfo)
			message.TextContent = template.GenerateEmail(template.TransactionToken, emailInfo, transferInfo)

			if err := emailClient.Send(message); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
		}
	}
	return nil
}

func addTokenToContract(store *db.SQLStore, contract string, transfer Erc20Transfer) error {
	decimals, _ := strconv.Atoi(transfer.TokenDecimals)
	params := sqlc.AddTokenToContractByAddressParams{
		Address:  contract,
		Name:     transfer.TokenName,
		Symbol:   transfer.TokenSymbol,
		Token:    transfer.Contract,
		Decimals: int64(decimals),
	}

	if _, err := store.AddTokenToContractByAddress(context.Background(), params); err != nil {
		return err
	}
	return nil
}
