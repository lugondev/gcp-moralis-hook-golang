package routes

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	sqlc "moralis-webhook/db/sqlc"
	"moralis-webhook/eth"
	"net/http"
	"strconv"
)

func (r *router) addContract(c echo.Context) error {
	var data NewContractRequest
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(data); err != nil {
		fmt.Println("error:", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	chain, err := eth.GetChainInfoByChainId(strconv.FormatInt(data.ChainId, 10))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	isContract := eth.IsContract(lo.Sample[string](chain.Rpc), common.HexToAddress(data.Address))

	addContract, err := r.SQLStore.AddContract(context.Background(), sqlc.AddContractParams{
		Name:       data.Name,
		Network:    chain.Name,
		Address:    data.Address,
		ChainID:    strconv.FormatInt(data.ChainId, 10),
		IsContract: isContract,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, addContract)
}

func (r *router) removeTokenFromContractByAddress(c echo.Context) error {
	var data RemoveTokenFromContractRequest
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(data); err != nil {
		fmt.Println("error:", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := r.SQLStore.RemoveTokenFromContractByAddress(context.Background(), sqlc.RemoveTokenFromContractByAddressParams{
		Address: data.Address,
		Token:   data.Token,
	}); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}

func (r *router) mapEmailContract(c echo.Context) error {
	var data NewMapEmailContract
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(data); err != nil {
		fmt.Println("error:", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	mappedEmailContract, err := r.SQLStore.MapEmailContract(context.Background(), sqlc.MapEmailContractParams{
		EmailID:    data.EmailId,
		ContractID: data.ContractId,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, mappedEmailContract)
}

func (r *router) getContractByAddress(c echo.Context) error {
	contractAddress := c.Param("address")
	if !common.IsHexAddress(contractAddress) {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"message": "address is invalid",
		})
	}

	contractByAddress, err := r.SQLStore.GetContractByAddress(context.Background(), contractAddress)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, contractByAddress)
}

func (r *router) listTokensInContractByAddress(c echo.Context) error {
	contractAddress := c.Param("address")
	if !common.IsHexAddress(contractAddress) {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"message": "address is not a valid address",
		})
	}

	contractByAddress, err := r.SQLStore.ListTokenInContractByAddress(context.Background(), contractAddress)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, contractByAddress)
}

func (r *router) listEmailsSubscription(c echo.Context) error {
	contractId := c.Param("contractId")
	parseInt, err := strconv.Atoi(contractId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"message": "contract id is not a number",
		})
	}
	mappedEmailContract, err := r.SQLStore.ListEmailsSubscription(context.Background(), int64(parseInt))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, mappedEmailContract)
}

func (r *router) listEmailsSubscriptionByAddress(c echo.Context) error {
	contractAddress := c.Param("address")
	if !common.IsHexAddress(contractAddress) {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"message": "address is not a valid address",
		})
	}
	mappedEmailContract, err := r.SQLStore.ListEmailsSubscriptionByAddress(context.Background(), contractAddress)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, mappedEmailContract)
}

func (r *router) listContracts(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	if limit <= 0 {
		limit = 10
	}
	listContracts, err := r.SQLStore.ListContracts(context.Background(), sqlc.ListContractsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, listContracts)
}
