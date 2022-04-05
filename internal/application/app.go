package application

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/mauromamani/go-clean-architecture/config"
)

type application struct {
	cfg *config.Config
	srv *http.Server
	db  *sql.DB
}

func New(cfg *config.Config, db *sql.DB) *application {
	return &application{
		cfg: cfg,
		db:  db,
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

	log.Println("starting server")
	if err := app.srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
