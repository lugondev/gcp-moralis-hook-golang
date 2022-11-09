package routes

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	sqlc "moralis-webhook/db/sqlc"
	"net/http"
)

func (r *router) addEmail(c echo.Context) error {
	var data sqlc.AddEmailParams
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	addedEmail, err := r.SQLStore.AddEmail(context.Background(), data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, addedEmail)
}

func (r *router) listEmail(c echo.Context) error {
	var data sqlc.ListEmailsParams
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if data.Limit == 0 {
		data.Limit = 10
	}

	listEmails, err := r.SQLStore.ListEmails(context.Background(), data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, listEmails)
}

func (r *router) emailSubscriber(c echo.Context) error {
	data := AddressRequest{
		Address: c.QueryParam("address"),
	}

	if err := c.Validate(data); err != nil {
		fmt.Println("error:", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	listEmails, err := r.SQLStore.GetEmailForContract(context.Background(), 1)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, listEmails)
}
