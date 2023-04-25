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
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	addedEmail, err := r.SQLStore.AddEmail(context.Background(), data)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, addedEmail)
}

func (r *router) listEmail(c echo.Context) error {
	var data sqlc.ListEmailsParams
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if data.Limit == 0 {
		data.Limit = 10
	}

	listEmails, err := r.SQLStore.ListEmails(context.Background(), data)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, listEmails)
}

func (r *router) emailSubscriber(c echo.Context) error {
	data := AddressRequest{
		Address: c.QueryParam("address"),
	}

	if err := c.Validate(data); err != nil {
		fmt.Println("error:", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	listEmails, err := r.SQLStore.GetEmailForContract(context.Background(), 1)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, listEmails)
}
