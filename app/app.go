package app

import (
	"go-api/api"
	"go-api/auth"
	"go-api/campaign"
	"go-api/config"
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
	dbConfig := config.DbConfig{
		Username:  config.DbUsername,
		Password:  config.DbPassword,
		Host:      config.DbHost,
		Port:      config.DbPort,
		DBName:    config.DbName,
		Charset:   config.DbCharset,
		ParseTime: config.DbParseTime,
		Loc:       config.DbLoc,
	}
	db := database.InitDatabase(dbConfig)

	// Repositories
	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	// Services
	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	transactionService := transaction.NewService(transactionRepository, campaignRepository)

	// JWT Service
	authService, err := auth.NewService()
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
