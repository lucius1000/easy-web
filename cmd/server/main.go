package main

import (
	"github.com/user/go-gin-gorm-starter/internal/config"
	"github.com/user/go-gin-gorm-starter/internal/handler"
	"github.com/user/go-gin-gorm-starter/internal/repository"
	"github.com/user/go-gin-gorm-starter/internal/router"
	"github.com/user/go-gin-gorm-starter/internal/service"
	"github.com/user/go-gin-gorm-starter/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	// 1. Initialize Logger
	logger.InitLogger()
	defer logger.Log.Sync()

	logger.Log.Info("Starting server...")

	// Load configuration
	cfg := config.LoadConfig()

	// 2. Initialize Database
	db, err := config.InitDatabase(cfg)
	if err != nil {
		logger.Log.Fatal("Failed to connect database", zap.Error(err))
	}

	// 3. Setup Dependencies (Dependency Injection)
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// 4. Setup Router
	r := router.SetupRouter(userHandler)

	// 5. Start Server
	serverAddr := ":" + cfg.AppPort
	logger.Log.Info("Server listening on " + serverAddr)
	if err := r.Run(serverAddr); err != nil {
		logger.Log.Fatal("Failed to start server", zap.Error(err))
	}
}
