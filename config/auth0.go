package config

import (
	"time"

	"github.com/spf13/viper"
)

const (
	issuerKey     = "auth.jwt.issuer"
	audienceKey   = "auth.jwt.audience"
	claimsPathKey = "auth.jwt.claimsPath"
)

const (
	domainAuth0Key       = "auth.auth0.domain"
	clientIdAuth0Key     = "auth.auth0.clientId"
	clientSecretAuth0Key = "auth.auth0.clientSecret"
)

type OIDCConfig struct {
	Issuer     string
	CacheTTL   time.Duration
	Audience   []string
	ClaimsPath string
}

type Auth0Config struct {
	Domain       string
	ClientId     string
	ClientSecret string
	oidc         *OIDCConfig
}

func NewAuth0Config() *Auth0Config {
	issuer := viper.GetString(issuerKey)
	audience := viper.GetStringSlice(audienceKey)
	claimsPath := viper.GetString(claimsPathKey)
	domainAuth0 := viper.GetString(domainAuth0Key)
	clientIdAuth0 := viper.GetString(clientIdAuth0Key)
	clientSecretAuth0 := viper.GetString(clientSecretAuth0Key)

	if issuer == "" || audience == nil || claimsPath == "" || domainAuth0 == "" || clientIdAuth0 == "" || clientSecretAuth0 == "" {
		panic("cannot load configuration")
		return nil
	}

	auth0Config := &Auth0Config{
		Domain:       domainAuth0,
		ClientId:     clientIdAuth0,
		ClientSecret: clientSecretAuth0,
	}
	auth0Config.oidc = &OIDCConfig{
		Issuer:   issuer,
		Audience: audience,
		CacheTTL: 5 * time.Minute,
	}

	if claimsPath != "" {
		auth0Config.oidc.ClaimsPath = claimsPath
	}

	return auth0Config
}
