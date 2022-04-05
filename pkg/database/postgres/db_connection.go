package postgres

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/mauromamani/go-clean-architecture/config"
)

func NewConnection(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open(cfg.Postgres.PgDriver, cfg.Postgres.DSN)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.Postgres.PgMaxOpenConns)
	db.SetMaxIdleConns(cfg.Postgres.PgMaxIdleConns)

	duration, err := time.ParseDuration(cfg.Postgres.PgMaxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
