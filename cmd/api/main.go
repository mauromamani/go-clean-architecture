package main

import (
	"log"

	"github.com/mauromamani/go-clean-architecture/config"
	"github.com/mauromamani/go-clean-architecture/internal/application"
	"github.com/mauromamani/go-clean-architecture/pkg/database/postgres"
	"github.com/mauromamani/go-clean-architecture/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	configPath := "./config/config-local"

	configFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	cfg, err := config.ParseConfig(configFile)
	if err != nil {
		log.Fatal("failed to parse config")
	}

	appLogger := logger.New(cfg)
	appLogger.InitLogger()

	db, err := postgres.NewConnection(cfg)
	if err != nil {
		appLogger.Fatal("failed to connect postgres database", zap.String("error", err.Error()))
	} else {
		appLogger.Info("connection pool establieshed")
	}

	app := application.New(cfg, db, appLogger)

	err = app.Run()
	if err != nil {
		appLogger.Fatal("failed to run server", zap.String("error", err.Error()))
	}
}
