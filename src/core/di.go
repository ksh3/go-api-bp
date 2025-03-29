package core

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"

	"github.com/ksh3/go-api/src/core/config"
	"github.com/ksh3/go-api/src/feature/user/infrastructure/service"
)

type AppContext struct {
	Container *dig.Container
}

func NewAppContext() *AppContext {
	container := dig.New()

	if err := container.Provide(func() *gin.Engine {
		env := config.GetAppEnv()
		switch env {
		case config.ProdEnvKey:
			gin.SetMode(gin.ReleaseMode)
		case config.StagingEnvKey:
			gin.SetMode(gin.TestMode)
		case config.DevEnvKey:
			gin.SetMode(gin.DebugMode)
		default:
			log.Fatalf("Invalid environment mode")
		}
		return gin.Default()
	}); err != nil {
		log.Fatalf("failed to provide gin router: %v", err)
	}

	if err := container.Provide(service.NewElasticsearchService); err != nil {
		log.Fatalf("failed to provide user service: %v", err)
	}

	if err := container.Provide(config.NewMongoDB); err != nil {
		log.Fatalf("failed to provide mongo")
	}

	// NOTE: Local database
	// if err := container.Provide(config.NewBadgerDB); err != nil {
	// 	log.Fatalf("failed to provide badger db: %v", err)
	// }

	return &AppContext{
		Container: container,
	}
}

func (a *AppContext) Close() error {
	return a.Container.Invoke(func(db *mongo.Database) error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := db.Client().Disconnect(ctx); err != nil {
			return err
		}
		return nil
	})
}
