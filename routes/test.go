package routes

import (
	"github.com/labstack/echo/v4"
	"moralis-webhook/notifier"
)

func (r *router) sendTelegram(_ echo.Context) error {
	return (*r.notifier).Notify(notifier.Message{Content: "Hello World"})
}
