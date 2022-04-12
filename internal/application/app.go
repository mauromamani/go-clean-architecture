package application

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/mauromamani/go-clean-architecture/config"
	"github.com/mauromamani/go-clean-architecture/pkg/logger"
)

type application struct {
	cfg    *config.Config
	srv    *http.Server
	db     *sql.DB
	logger logger.Logger
}

func New(cfg *config.Config, db *sql.DB, handler *httprouter.Router, logger logger.Logger) *application {
	return &application{
		cfg: cfg,
		db:  db,
		srv: &http.Server{
			Handler:      handler,
			Addr:         cfg.Server.Port,
			ErrorLog:     log.Default(),
			IdleTimeout:  cfg.Server.IdleTimeout * time.Minute,
			ReadTimeout:  cfg.Server.ReadTimeout * time.Second,
			WriteTimeout: cfg.Server.WriteTimeout * time.Second,
		},
		logger: logger,
	}
}

// Run: run server on port 3000
func (app *application) Run() error {
	app.mapHandlers()

	app.logger.Infof("starting server on port %s", app.cfg.Server.Port)
	if err := app.srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
