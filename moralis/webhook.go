package moralis

import (
	"github.com/labstack/echo/v4"
	"log"
	"moralis-webhook/email"
	"moralis-webhook/email/template"
	"moralis-webhook/eth"
	"net/http"
	"strconv"
	"time"
)

func Hook(emailClient email.Client) func(c echo.Context) error {
	return func(c echo.Context) error {
		contract := c.Param("contract")

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
			{Name: "Dev", Email: "lugondev@gmail.com"},
		}

		if !payload.Confirmed {
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

			if len(payload.Erc20Transfers) > 0 {
				for _, transferEvent := range payload.Erc20Transfers {
					transferInfo := template.Erc20Transfer{
						Tx:          transferEvent.TransactionHash,
						From:        transferEvent.From,
						To:          transferEvent.To,
						Amount:      transferEvent.ValueWithDecimals,
						TokenName:   transferEvent.TokenName,
						TokenSymbol: transferEvent.TokenSymbol,
					}

					message.HtmlContent = template.GenerateEmail(template.TransactionToken, emailInfo, transferInfo)
					message.TextContent = template.GenerateEmail(template.TransactionToken, emailInfo, transferInfo)

					if err := emailClient.Send(message); err != nil {
						return echo.NewHTTPError(http.StatusBadRequest, err.Error())
					}
				}
			}

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
						return echo.NewHTTPError(http.StatusBadRequest, err.Error())
					}
				}
			}
		}
		return c.JSON(http.StatusOK, payload)
	}
}
