package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"moralis-webhook/email/factory"
	"moralis-webhook/email/sendgrid"
	"moralis-webhook/moralis"
	"net/http"
)

// GitCommitLog is set at build-time
var GitCommitLog string

func init() {
	fmt.Printf("GIT log: %s\n", GitCommitLog)
}

func main() {
	// Echo instance
	e := echo.New()
	emailClient, err := factory.NewEmailClient(factory.Config{
		Adapter: "sendgrid",
		Sendgrid: sendgrid.Config{
			Debug:  true,
			ApiKey: "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",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	// Routes
	e.GET("/", hello)
	e.GET("/git", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("GIT log: %s\n", GitCommitLog))
	})
	e.POST("/webhook-multisig/:contract", moralis.Hook(emailClient))

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello, World!",
	})
}
