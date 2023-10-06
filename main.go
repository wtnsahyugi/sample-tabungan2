package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"sample-tabungan2/config"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World2!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func NewPgx(cfg config.Database) (*pgxpool.Pool, error) {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
	)
	poolcfg, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		return nil, err
	}
	poolcfg.MaxConns = int32(cfg.MaxOpenConns)
	poolcfg.MaxConnIdleTime = time.Duration(cfg.MaxIdleLifetime) * time.Minute
	poolcfg.MaxConnLifetime = time.Duration(cfg.MaxConnLifetime) * time.Minute

	conn, err := pgxpool.ConnectConfig(context.Background(), poolcfg)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
