package main

import (
	"log"
	"subscription-service/internal/app"
	"subscription-service/internal/config"
)

// @title Subscription Service API
// @version 1.0
// @description REST API для управления подписками
// @host localhost:8080
// @BasePath /api/v1
func main() {
	cfg := config.LoadConfig()

	application := app.New(cfg)

	if err := application.Run(); err != nil {
		log.Fatal("Failed to run application: ", err)
	}
}
