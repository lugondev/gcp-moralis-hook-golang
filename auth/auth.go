package auth

import (
	"github.com/labstack/echo/v4"
	"moralis-webhook/dto"
)

type Context struct {
	echo.Context
	Claims *dto.UserClaims
}
