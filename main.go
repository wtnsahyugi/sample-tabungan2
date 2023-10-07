package main

import (
	"context"
	"fmt"
	"net/http"

	"sample-tabungan2/api"
	"sample-tabungan2/config"
	"sample-tabungan2/internal/repository"
	"sample-tabungan2/service/user"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// init config
	cfg, err := config.NewConfig(".env")
	checkError(err)

	// init db
	pgxPool, err := NewPgx(*cfg)
	checkError(err)

	// init repo
	userRepo := repository.NewUserRepository(pgxPool)

	//init service
	userSvc := user.NewUserService(userRepo)

	// init handler
	userHandler := api.NewUserHTTPHandler(userSvc)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World2!")
	})
	e.Add(http.MethodPost, "/daftar", userHandler.Create)
	e.Add(http.MethodPost, "/tabung", userHandler.Deposit)
	e.Add(http.MethodPost, "/tarik", userHandler.Withdraw)
	e.Logger.Fatal(e.Start(":1323"))

}

func NewPgx(cfg config.Config) (*pgxpool.Pool, error) {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
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

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
