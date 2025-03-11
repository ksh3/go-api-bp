package config

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type EnvMode string

const (
	ProdEnvKey     = "prod"
	StagingEnvKey  = "staging"
	DevEnvKey      = "dev"
	MongoDBURI     = "mongodb://localhost:27017"
	MongoDBName    = "database"
	MongoDBTimeOut = 10 * time.Second
)

var (
	DevTrustedProxies  = []string{"127.0.0.1", "localhost"}
	StgTrustedProxies  = getTrustedProxies("STAGING_TRUSTED_PROXIES")
	ProdTrustedProxies = getTrustedProxies("PROD_TRUSTED_PROXIES")
)

func getTrustedProxies(envVar string) []string {
	proxies := os.Getenv(envVar)
	if proxies == "" {
		return []string{}
	}
	return strings.Split(proxies, ",")
}

func GetAppEnv() EnvMode {
	env := os.Getenv("APP_ENV")
	switch env {
	case string(StagingEnvKey):
		return StagingEnvKey
	case string(ProdEnvKey):
		return ProdEnvKey
	case string(DevEnvKey):
		return DevEnvKey
	default:
		panic(fmt.Sprintf("Invalid environment mode: %s", env))
	}
}
