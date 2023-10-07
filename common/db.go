package common

import (
	"context"
	"fmt"
	"sample-tabungan2/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPgx(cfg config.Config) (*pgxpool.Pool, error) {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DatabaseUsername, cfg.DatabasePassword, cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseName,
	)
	poolcfg, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		return nil, err
	}

	conn, err := pgxpool.ConnectConfig(context.Background(), poolcfg)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
