package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"moralis-webhook/db"
)

const (
	profileKey = "db.profile"
	profileEnv = "DB_PROFILE"
)

type PostgresConfig struct {
	Host     string
	Port     int
	DBName   string
	Username string
	Password string
	SSLMode  string
	TimeZone string

	ConnectionMaxIdleTime int
	ConnectionMaxLifetime int
	MaxIdleConnections    int
	MaxOpenConnections    int

	Logger struct {
		LogLevel string
	}
}

type DBConfig struct {
	Profile    string
	Postgresql PostgresConfig
}

func init() {
	_ = viper.BindEnv(profileKey, profileEnv)
}

func NewDB() (*db.SQLStore, error) {
	var config DBConfig
	var postgresConfig PostgresConfig
	if err := viper.UnmarshalKey("db", &config); err != nil {
		panic(err)
	}
	var dsn string
	if config.Profile != "" {
		if err := viper.UnmarshalKey(fmt.Sprintf("db.%s", config.Profile), &postgresConfig); err != nil {
			panic(err)
		}
		dsn = getDSN(postgresConfig)
	} else {
		dsn = getDSN(config.Postgresql)
	}

	openDB, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db.NewStore(openDB), nil
}

func getDSN(configuration PostgresConfig) string {
	return fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=%s TimeZone=%s",
		configuration.Host,
		configuration.Port,
		configuration.DBName,
		configuration.Username,
		configuration.Password,
		configuration.SSLMode,
		configuration.TimeZone,
	)
}
