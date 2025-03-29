package server

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ksh3/go-api/src/core"
	"github.com/ksh3/go-api/src/core/config"
	"github.com/ksh3/go-api/src/core/contract"
	"github.com/ksh3/go-api/src/server/middleware"
	"github.com/ksh3/go-api/src/server/routes"
)

// NOTE: Register all routes
func Run(router *gin.Engine, db *mongo.Database, logger *core.Logger) {
	envMode := config.GetAppEnv()

	routes.SetupSystemRoutes(router)
	routes.SetupWebRoutes(router, logger)
	routes.SetupAPIv1Routes(router, db, logger)

	switch envMode {
	case config.ProdEnvKey:
		router.SetTrustedProxies(config.ProdTrustedProxies)
	case config.StagingEnvKey:
		router.SetTrustedProxies(config.StgTrustedProxies)
	case config.DevEnvKey:
		router.SetTrustedProxies(config.DevTrustedProxies)
	default:
		logger.ErrorLog(
			contract.InternalError("Invalid environment mode", nil),
		)
		os.Exit(1)
	}
	router.Use(middleware.CORS())
	router.LoadHTMLGlob("src/ui/web/public/*.html")
	router.Static("/static", "./src/ui/web/static")

	logger.InfoLog(fmt.Sprintf("[%s] Starting server on :8080", envMode))
	router.Run(":8080")
}
