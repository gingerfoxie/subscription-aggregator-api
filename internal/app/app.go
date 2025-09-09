package app

import (
	"fmt"
	"log"
	"subscription-service/internal/config"
	"subscription-service/internal/handlers"
	"subscription-service/internal/repository"
	"subscription-service/internal/routes"
	"subscription-service/internal/service"
	"subscription-service/pkg/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "subscription-service/docs"
)

type App struct {
	config     *config.Config
	db         *gorm.DB
	router     *gin.Engine
	httpServer *gin.Engine
}

func New(cfg *config.Config) *App {
	return &App{
		config: cfg,
	}
}

func (a *App) Run() error {
	if err := a.initDB(); err != nil {
		return err
	}

	if err := a.initServices(); err != nil {
		return err
	}

	a.initRouter()

	log.Printf("Server starting on port %s", a.config.ServerPort)
	return a.router.Run(":" + a.config.ServerPort)
}

func (a *App) initDB() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		a.config.DBHost, a.config.DBPort, a.config.DBUser, a.config.DBPassword, a.config.DBName, a.config.DBSslMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	a.db = db
	return nil
}

func (a *App) initServices() error {
	// Initialize repositories
	subscriptionRepo := repository.NewSubscriptionRepository(a.db)

	// Initialize services
	subscriptionService := service.NewSubscriptionService(subscriptionRepo)

	// Initialize handlers
	subscriptionHandler := handlers.NewSubscriptionHandler(subscriptionService)

	// Initialize router
	a.router = routes.SetupRouter(subscriptionHandler)

	return nil
}

func (a *App) initRouter() {
	// Additional router configuration can be done here
	logger.InitLogger(a.config.LogOutput, a.config.LogLevel)
}
