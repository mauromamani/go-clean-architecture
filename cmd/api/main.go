package main

import (
	"context"
	"log"

	"github.com/mauromamani/go-clean-architecture/config"
	"github.com/mauromamani/go-clean-architecture/ent"
	"github.com/mauromamani/go-clean-architecture/internal/server"

	_ "github.com/mattn/go-sqlite3"
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

	// database client
	client, err := ent.Open("sqlite3", cfg.SQLite.FileName+"?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// Run auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	s := server.NewServer(cfg, client)

	s.Run()
}
