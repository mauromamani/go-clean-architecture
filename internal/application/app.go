package application

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mauromamani/go-clean-architecture/config"
)

type application struct {
	cfg *config.Config
	srv *http.Server
}

func New(cfg *config.Config) *application {
	return &application{
		cfg: cfg,
		srv: &http.Server{
			Addr:         cfg.Server.Port,
			ErrorLog:     log.Default(),
			IdleTimeout:  cfg.Server.IdleTimeout * time.Minute,
			ReadTimeout:  cfg.Server.ReadTimeout * time.Second,
			WriteTimeout: cfg.Server.WriteTimeout * time.Second,
		},
	}
}

// Run: run server on port 3000
func (app *application) Run() error {
	app.srv.Handler = app.mapHandlers()

	fmt.Println("Starting server")
	if err := app.srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
