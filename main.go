package main

import (
	"go-api/auth"
	"go-api/campaign"
	"go-api/database"
	"go-api/handler"
	"go-api/middleware"
	"go-api/transaction"
	"go-api/user"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Reading database config file and database init
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := database.Init()

	// Repositories
	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)
	// Services
	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	transactionService := transaction.NewService(transactionRepository, campaignRepository)
	authService, err := auth.NewService(".jwt_secret")
	if err != nil {
		log.Fatal("Error initializing JWT Service", err)
	}
	// Handlers
	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	// API Endpoint Routes
	router := gin.Default()
	router.Static("/images", "./images")
	api := router.Group("/api/v1")
	// User-related endpoints
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.LoginUser)
	api.POST("/email_check", userHandler.CheckEmail)
	api.POST("/avatars", middleware.AuthMiddleware(authService, userService), userHandler.AvatarUpload)
	// Campaign-related endpoints
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns", middleware.AuthMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", middleware.AuthMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", middleware.AuthMiddleware(authService, userService), campaignHandler.UploadImage)
	// Transaction-related endpoints
	api.GET("/campaigns/:id/transactions", middleware.AuthMiddleware(authService, userService), transactionHandler.GetCampaignTransactions)
	api.GET("/transactions", middleware.AuthMiddleware(authService, userService), transactionHandler.GetUserTransaction)
	router.Run()
}
