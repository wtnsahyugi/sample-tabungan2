package main

import (
	"net/http"

	"sample-tabungan2/api"
	"sample-tabungan2/common"
	"sample-tabungan2/config"
	"sample-tabungan2/internal/repository"
	"sample-tabungan2/service/user"

	"github.com/gocraft/work"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// init config
	cfg, err := config.NewConfig(".env")
	checkError(err)

	// init db
	pgxPool, err := common.NewPgx(*cfg)
	checkError(err)

	// init repo
	userRepo := repository.NewUserRepository(pgxPool)

	// init pool worker for async process
	var enqueuer = work.NewEnqueuer(cfg.WorkerNamespace, common.NewRedisPool(*cfg))

	//init service
	userSvc := user.NewUserService(userRepo, enqueuer)

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

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
