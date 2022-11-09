package config

import (
	"github.com/spf13/viper"
)

var (
	appConfig *AppConfig
)

const (
	serverPort = "server.port"
)

type AppConfig struct {
	auth0 *Auth0Config
	port  int
}

func NewConfig() *AppConfig {
	port := viper.GetInt(serverPort)
	auth0Config := NewAuth0Config()

	return &AppConfig{
		auth0: auth0Config,
		port:  port,
	}
}

func SetAppConfig(config *AppConfig) {
	appConfig = config
}

func GetAppConfig() *AppConfig {
	return appConfig
}

func (a *AppConfig) GetAuth0Config() *Auth0Config {
	return a.auth0
}

func (a *AppConfig) GetServerPort() int {
	return a.port
}
