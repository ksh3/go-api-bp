package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ksh3/go-api/src/core"
	"github.com/ksh3/go-api/src/feature/user/domain/usecase"
	"github.com/ksh3/go-api/src/feature/user/infrastructure/repository"
	"github.com/ksh3/go-api/src/feature/user/infrastructure/service"
)

const (
	WebIndex         = "/"
	WebArticle       = "/article"
	WebArticleDetail = "/article/:id"
	WebAbout         = "/about"
	WebContact       = "/contact"

	// NOTE: API
	APIV1Base = "/v1"

	APIUserBase    = "/users"
	APIUserDetail  = "/users/:id"
	APIUserProfile = "/users/profile"

	APISystemBase = "/_system"

	APIHealthCheck  = "/health"
	APISystemStatus = "/status"
	APIMetrics      = "/metrics"

	APIWebhookBase       = "/webhook"
	APIWebhookEvent      = "/event"
	APIWebhookUserUpdate = "/user_update"
)

// NOTE: Setup all system routes
func SetupSystemRoutes(router *gin.Engine) {
	systemGroup := router.Group(APISystemBase)
	systemGroup.GET(APIHealthCheck, func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	systemGroup.GET(APISystemStatus, func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "running", "uptime": "48h"})
	})
}

// NOTE: Setup all web routes
func SetupWebRoutes(router *gin.Engine, logger *core.Logger) {
	router.GET(WebIndex, func(ctx *gin.Context) {
		logger.InfoLog("Hello")
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET(WebArticle, func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to the top!"})
	})

	router.GET(WebArticleDetail, func(ctx *gin.Context) {
		articleID := ctx.Param("id")
		ctx.JSON(http.StatusOK, gin.H{"message": "Article ID: " + articleID})
	})
}

// NOTE: Setup all API v1 routes
func SetupAPIv1Routes(router *gin.Engine, db *mongo.Database, logger *core.Logger) {
	apiV1 := router.Group(APIV1Base)

	_setupUserRoutes := func(router *gin.RouterGroup, db *mongo.Database) {
		router.GET(APIUserBase, func(ctx *gin.Context) {
			entities, err := usecase.NewDefaultUserUseCase(
				repository.NewDefaultUserRepository(db),
				service.NewGoogleCalendarService(),
			).GetSubscribeUsers()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"users": entities})
		})

		router.GET(APIUserDetail, func(ctx *gin.Context) {
			userID := ctx.Param("id")
			ctx.JSON(http.StatusOK, gin.H{"message": "User ID: " + userID})
		})

		router.GET(APIUserProfile, func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "User Profile"})
		})
	}

	_setupUserRoutes(apiV1, db)
}
