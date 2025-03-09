package config

import "time"

const (
	ProdEnvKey     = "prod"
	StagingEnvKey  = "staging"
	DevEnvKey      = "dev"
	MongoDBURI     = "mongodb://localhost:27017"
	MongoDBName    = "database"
	MongoDBTimeOut = 10 * time.Second
)
