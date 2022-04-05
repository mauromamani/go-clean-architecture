package main

import (
	"log"

	"github.com/mauromamani/go-clean-architecture/config"
	"github.com/mauromamani/go-clean-architecture/internal/application"
	"github.com/mauromamani/go-clean-architecture/pkg/database/postgres"
)

func main() {
	configPath := "./config/config-local"

	configFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	cfg, err := config.ParseConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	db, err := postgres.NewConnection(cfg)
	if err != nil {
		panic(err)
	}
	log.Println("Connection pool established!")

	app := application.New(cfg, db)

	err = app.Run()
	if err != nil {
		panic(err)
	}
}
