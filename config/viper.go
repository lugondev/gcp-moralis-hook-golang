package config

import (
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"strings"
)

func LoadConfiguration() {
	readFile()

	appConfig := NewConfig()
	SetAppConfig(appConfig)

	readEnvironment()
	overrideConfiguration()
}

func readFile() {
	viper.SetConfigFile("configuration.yml")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		log.Error("Cannot read configuration file", err)
		panic(err)
	}
}

func readEnvironment() {
	viper.SetEnvPrefix("WAAS")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func overrideConfiguration() {
	for _, key := range viper.AllKeys() {
		val := viper.Get(key)
		viper.Set(key, val)
	}
}
