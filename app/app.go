package app

import (
	"go-api/api"
	"go-api/auth"
	"go-api/campaign"
	"go-api/database"
	"go-api/transaction"
	"go-api/user"
	"log"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router      *gin.Engine
	AuthService auth.Service
}

func InitApp() *App {
	db := database.Init()

	// Repositories
	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	// Services
	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	transactionService := transaction.NewService(transactionRepository, campaignRepository)

	// JWT Service
	authService, err := auth.NewService(".jwt_secret")
	if err != nil {
		log.Fatal("Error initializing JWT Service", err)
	}

	// Creating Gin Router
	router := api.InitRouter()

	// setup routes
	api.SetupRoutes(router, authService, userService, campaignService, transactionService)

	return &App{
		Router:      router,
		AuthService: authService,
	}
}
