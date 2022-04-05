package main

import (
	"log"

	"github.com/mauromamani/go-clean-architecture/config"
	"github.com/mauromamani/go-clean-architecture/internal/application"
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

	app := application.New(cfg)

	err = app.Run()
	if err != nil {
		panic(err)
	}
}
