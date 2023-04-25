package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"moralis-webhook/config"
	"moralis-webhook/moralis"
	"moralis-webhook/notifier"
	"moralis-webhook/routes"
	"net/http"
)

// GitCommitLog is set at build-time
var GitCommitLog string

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func init() {
	log.Printf("GIT log: %s\n", GitCommitLog)
	config.LoadConfiguration()
}

const RpcUrl = "https://arb-mainnet.g.alchemy.com/v2/5yyYIuYGR_6hYQ6YVoThlQoabPZyimN9"

//const RpcUrl = "https://nameless-multi-tent.arbitrum-mainnet.discover.quiknode.pro/26e310d6e5af40e43af753866a314e779353cc8d/"

//const RpcUrl = "https://arbitrum-mainnet.infura.io/v3/26a2d6b5e2d2446c95ccf258f1f81f24"

func main() {
	dbStore, err := config.NewDB()
	if err != nil {
		log.Fatalf("failed to create DB store: %v", err)
	}

	appConfig := config.GetAppConfig()

	newNotifier, err := notifier.NewNotifier()
	if err != nil {
		log.Fatalf("failed to create Notifier client: %v", err)
	}

	//emailClient, err := email.NewEmailClient(&newNotifier)
	//if err != nil {
	//	log.Fatalf("failed to create Email client: %v", err)
	//}

	// Echo instance
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

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
	client, err := ethclient.Dial(RpcUrl)
	if err != nil {
		return
	}
	e.POST("/webhook-multisig/:contract", moralis.Hook(client, newNotifier))

	// Init routers
	routes.NewRouter(e, dbStore, &newNotifier)
	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appConfig.GetServerPort())))
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello, World!",
	})
}
