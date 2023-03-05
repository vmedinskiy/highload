package database

import (
	"github.com/go-faster/errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/vmedinskiy/highload/internal/pkg/log"
)

type Config interface {
	DbUri() string
	DbOpenConns() int
	DbIdleConns() int
}

func New(cfg Config, logger pgx.Logger) (*sqlx.DB, error) {
	if cfg.DbUri() == "" {
		return nil, errors.New("no DatabaseURI provided")
	}

	log.L().Infof("connecting to db: %s", cfg.DbUri())
	connConfig, err := pgx.ParseConfig(cfg.DbUri())
	if err != nil {
		return nil, errors.Wrap(err, "parse db config")
	}
	connConfig.Logger = logger
	connConfig.LogLevel = pgx.LogLevelDebug
	connStr := stdlib.RegisterConnConfig(connConfig)
	db, err := sqlx.Connect("pgx", connStr)
	if err != nil {
		return nil, errors.Wrap(err, "preparing db connection")
	}
	db.SetMaxOpenConns(cfg.DbOpenConns())
	db.SetMaxIdleConns(cfg.DbIdleConns())
	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "connecting to database")
	}
	log.L().Info("connected successfully")
	return db, nil

}
