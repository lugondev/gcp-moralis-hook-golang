package routes

import (
	"github.com/labstack/echo/v4"
	"moralis-webhook/db"
)

type router struct {
	*db.SQLStore
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

	NewMapEmailContract struct {
		EmailId    int64 `json:"emailId" validate:"required"`
		ContractId int64 `json:"contractId" validate:"required"`
	}
)

func NewRouter(e *echo.Echo, store *db.SQLStore) {
	router := &router{
		store,
	}

	emailRouter := e.Group("/email")
	emailRouter.POST("/add", router.addEmail)
	emailRouter.GET("/list", router.listEmail)
	emailRouter.GET("/subscription", router.emailSubscriber)

	contractRouter := e.Group("/contract")
	contractRouter.POST("/add", router.addContract)
	contractRouter.POST("/map-email", router.mapEmailContract)
	contractRouter.GET("/subscriptions/:contractId", router.listEmailsSubscription)
	contractRouter.GET("/subscriptions-by-address/:address", router.listEmailsSubscriptionByAddress)
}
