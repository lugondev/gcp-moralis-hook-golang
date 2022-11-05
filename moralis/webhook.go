package moralis

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"moralis-webhook/chains"
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
			fmt.Printf("unable to bind payload: %v\n", err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		chain, err := chains.GetChainInfo(payload.ChainId)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		recipient := email.Recipient{Name: "Test Receiver", Email: "lugon@alphatrue.com"}

		if !payload.Confirmed {
			i, _ := strconv.ParseInt(payload.Block.Timestamp, 10, 64)
			tm := time.Unix(i, 0)

			emailInfo := template.AppInfo{
				Contract: contract,
				Network:  chain.Name,
				User:     recipient.Name,
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

					var message = email.Message{
						Subject:     "Transaction Notification",
						Sender:      email.Sender{Name: "Notifier", Email: "notifier@alphatrue.dev"},
						Recipient:   recipient,
						HtmlContent: template.GenerateEmail(template.TransactionToken, emailInfo, transferInfo),
						TextContent: template.GenerateEmail(template.TransactionToken, emailInfo, transferInfo),
						CCs:         &[]email.Recipient{
							//{Name: "Quyn", Email: "quyen@alphatrue.com"},
							//{Name: "Toan", Email: "toan@alphatrue.com"},
						},
					}

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
					fmt.Println(inputDecoded)

					var message = email.Message{
						Subject:     "Transaction Notification",
						Sender:      email.Sender{Name: "Notifier", Email: "notifier@alphatrue.dev"},
						Recipient:   recipient,
						HtmlContent: template.GenerateEmail(template.TransactionDetail, emailInfo, txInfo, inputDecoded),
						TextContent: template.GenerateEmail(template.TransactionDetail, emailInfo, txInfo, inputDecoded),
						CCs:         &[]email.Recipient{
							//{Name: "Quyn", Email: "quyen@alphatrue.com"},
							//{Name: "Toan", Email: "toan@alphatrue.com"},
						},
					}

					if err := emailClient.Send(message); err != nil {
						return echo.NewHTTPError(http.StatusBadRequest, err.Error())
					}
				}
			}
		}
		return c.JSON(http.StatusOK, payload)
	}
}
