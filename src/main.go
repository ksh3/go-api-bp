package main // package

import (
	"os"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
	"github.com/ksh3/go-api/src/core"
	"github.com/ksh3/go-api/src/core/errors"
	"github.com/ksh3/go-api/src/core/i18n"
	"github.com/ksh3/go-api/src/core/logging"
	"github.com/ksh3/go-api/src/core/route"
)

func main() {
	appCtx := core.NewAppContext()

	logger, _ := logging.NewLogger()
	defer logger.Close()

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
			route.SetupRoutes(router, db, logger)
			logger.InfoLog("Starting server on :8080")
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
