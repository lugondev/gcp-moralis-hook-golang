package config

import (
	"time"

	"github.com/spf13/viper"
)

const (
	issuerKey     = "auth.jwt.issuer"
	issuerEnv     = "AUTH_ISSUER"
	audienceKey   = "auth.jwt.audience"
	audienceEnv   = "AUTH_AUDIENCE"
	claimsPathKey = "auth.jwt.claimsPath"
	claimsPathEnv = "AUTH_CLAIMS_PATH"
)

const (
	domainAuth0Key       = "auth.auth0.domain"
	domainAuth0Env       = "AUTH0_DOMAIN"
	clientIdAuth0Key     = "auth.auth0.clientId"
	clientIdAuth0Env     = "AUTH0_CLIENT_ID"
	clientSecretAuth0Key = "auth.auth0.clientSecret"
	clientSecretAuth0Env = "AUTH0_CLIENT_SECRET"
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

func init() {
	_ = viper.BindEnv(issuerKey, issuerEnv)
	_ = viper.BindEnv(audienceKey, audienceEnv)
	_ = viper.BindEnv(claimsPathKey, claimsPathEnv)

	_ = viper.BindEnv(domainAuth0Key, domainAuth0Env)
	_ = viper.BindEnv(clientIdAuth0Key, clientIdAuth0Env)
	_ = viper.BindEnv(clientSecretAuth0Key, clientSecretAuth0Env)
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
