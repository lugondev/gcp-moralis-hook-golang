package routes

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"moralis-webhook/db"
	"moralis-webhook/eth"
	"moralis-webhook/notifier"
	"net/http"
)

type router struct {
	*db.SQLStore
	notifier *notifier.Notifier
}

type (
	AddressRequest struct {
		Address string `json:"address" validate:"required,eth_addr"`
	}

	NewContractRequest struct {
		Name    string `json:"name" validate:"required"`
		Address string `json:"address" validate:"required,eth_addr"`
		ChainId int64  `json:"chainId" validate:"required"`
	}

	RemoveTokenFromContractRequest struct {
		Address string `json:"address" validate:"required,eth_addr"`
		Token   string `json:"token" validate:"required,eth_addr"`
	}

	NewMapEmailContract struct {
		EmailId    int64 `json:"emailId" validate:"required"`
		ContractId int64 `json:"contractId" validate:"required"`
	}
)

func NewRouter(e *echo.Echo, store *db.SQLStore, notifier *notifier.Notifier) {
	router := &router{
		store,
		notifier,
	}

	e.POST("/chain/:chainId", func(c echo.Context) error {
		chainId := c.Param("chainId")
		chain, err := eth.GetChainInfoByChainId(chainId)
		if err != nil {
			fmt.Println("cannot get chain info: ", err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, chain)
	})

	testRouter := e.Group("/test")
	testRouter.POST("/telegram", router.sendTelegram)

	emailRouter := e.Group("/email")
	emailRouter.POST("/add", router.addEmail)
	emailRouter.GET("/list", router.listEmail)
	emailRouter.GET("/subscription", router.emailSubscriber)

	contractRouter := e.Group("/contract")
	contractRouter.POST("/add", router.addContract)
	contractRouter.GET("/list", router.listContracts)

	contractRouter.POST("/map-email", router.mapEmailContract)
	contractRouter.GET("/subscriptions/:contractId", router.listEmailsSubscription)
	contractRouter.GET("/subscriptions-by-address/:address", router.listEmailsSubscriptionByAddress)
	contractRouter.GET("/contract-by-address/:address", router.getContractByAddress)
	contractRouter.GET("/tokens-by-address/:address", router.listTokensInContractByAddress)
	contractRouter.DELETE("/remove-token", router.removeTokenFromContractByAddress)
}
