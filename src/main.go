package main // package

import (
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
	"github.com/ksh3/go-api/src/core"
	"github.com/ksh3/go-api/src/core/config"
	"github.com/ksh3/go-api/src/core/errors"
	"github.com/ksh3/go-api/src/core/i18n"
	"github.com/ksh3/go-api/src/core/logging"
	"github.com/ksh3/go-api/src/core/route"
)

func main() {
	logger, _ := logging.NewLogger()
	defer logger.Close()

	envMode := config.GetAppEnv()

	appCtx := core.NewAppContext()
	translator := i18n.NewTranslator()

	if err := translator.LoadTranslations(
		"./src/core/i18n",
		"./src/feature/user/i18n",
	); err != nil {
		logger.ErrorLog(
			errors.InternalError(err.Error(), err),
		)
		os.Exit(1)
	}

	err := appCtx.Container.Invoke(
		func(router *gin.Engine, db *mongo.Database) {
			route.Register(router, db, logger)
			logger.InfoLog(fmt.Sprintf("[%s] Starting server on :8080", envMode))
			switch envMode {
			case config.ProdEnvKey:
				router.SetTrustedProxies(config.ProdTrustedProxies)
			case config.StagingEnvKey:
				router.SetTrustedProxies(config.StgTrustedProxies)
			case config.DevEnvKey:
				router.SetTrustedProxies(config.DevTrustedProxies)
			default:
				logger.ErrorLog(
					errors.InternalError("Invalid environment mode", nil),
				)
				os.Exit(1)
			}
			router.LoadHTMLGlob("src/ui/web/public/*.html")
			router.Static("/static", "./src/ui/web/static")
			router.Run(":8080")
		})
	if err != nil {
		appCtx.Close()
		logger.ErrorLog(
			errors.InternalError(err.Error(), err),
		)
		os.Exit(1)
	}
}
