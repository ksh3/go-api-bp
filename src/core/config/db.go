package config

import (
	"context"

	"github.com/dgraph-io/badger/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewBadgerDB() (*badger.DB, error) {
	opts := badger.DefaultOptions("./.badger_data")
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewMongoDB() (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(MongoDBURI)
	ctx, cancel := context.WithTimeout(
		context.Background(), MongoDBTimeOut,
	)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	db := client.Database(MongoDBName)

	return db, nil
}
