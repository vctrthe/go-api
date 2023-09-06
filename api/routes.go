package api

import (
	"go-api/auth"
	"go-api/campaign"
	"go-api/handler"
	"go-api/middleware"
	"go-api/transaction"
	"go-api/user"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, authService auth.Service, userService user.Service, campaignService campaign.Service, transactionService transaction.Service) {
	api := router.Group("/api/v1")

	// User-related endpoints
	userRoutes := api.Group("/users")
	userHandler := handler.NewUserHandler(userService, authService)
	userRoutes.POST("/", userHandler.RegisterUser)
	userRoutes.POST("/sessions", userHandler.LoginUser)
	userRoutes.POST("/email_check", userHandler.CheckEmail)
	userRoutes.POST("/avatars", middleware.AuthMiddleware(authService, userService), userHandler.AvatarUpload)

	// Campaign-related endpoints
	campaignRoutes := api.Group("/campaigns")
	campaignHandler := handler.NewCampaignHandler(campaignService)
	campaignRoutes.GET("/", campaignHandler.GetCampaigns)
	campaignRoutes.GET("/:id", campaignHandler.GetCampaign)
	campaignRoutes.POST("/", middleware.AuthMiddleware(authService, userService), campaignHandler.CreateCampaign)
	campaignRoutes.PUT("/:id", middleware.AuthMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	campaignRoutes.POST("/campaign-images", middleware.AuthMiddleware(authService, userService), campaignHandler.UploadImage)

	// Transaction-related endpoints
	transactionRoutes := api.Group("/transactions")
	transactionHandler := handler.NewTransactionHandler(transactionService)
	transactionRoutes.GET("/", middleware.AuthMiddleware(authService, userService), transactionHandler.GetUserTransaction)
	transactionRoutes.GET("/campaigns/:id", middleware.AuthMiddleware(authService, userService), transactionHandler.GetCampaignTransactions)
}
