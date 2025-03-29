package main // package

import (
	"os"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
	"github.com/ksh3/go-api/src/core"
	"github.com/ksh3/go-api/src/core/contract"
	"github.com/ksh3/go-api/src/core/i18n"
	"github.com/ksh3/go-api/src/server"
)

func main() {
	logger, _ := core.NewLogger()
	defer logger.Close()

	appCtx := core.NewAppContext()
	translator := i18n.NewTranslator()

	if err := translator.LoadTranslations(
		"./src/core/i18n",
		"./src/feature/user/i18n",
	); err != nil {
		logger.ErrorLog(
			contract.InternalError(err.Error(), err),
		)
		os.Exit(1)
	}

	err := appCtx.Container.Invoke(
		func(router *gin.Engine, db *mongo.Database) {
			server.Run(router, db, logger)
		})
	if err != nil {
		appCtx.Close()
		logger.ErrorLog(
			contract.InternalError(err.Error(), err),
		)
		os.Exit(1)
	}
}
