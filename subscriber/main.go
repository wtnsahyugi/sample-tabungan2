package main

import (
	"os"
	"os/signal"
	"sample-tabungan2/common"
	"sample-tabungan2/config"
	"sample-tabungan2/internal/repository"
	"sample-tabungan2/service/transaction"
	"syscall"

	"github.com/gocraft/work"
)

func main() {
	// init config
	cfg, err := config.NewConfig(".env")
	checkError(err)

	// init db
	pgxPool, err := common.NewPgx(*cfg)
	checkError(err)

	// init redis
	redisPool := common.NewRedisPool(*cfg)

	// init repository
	transactionRepo := repository.NewTransactionRepository(pgxPool)

	// init service
	trxSvc := transaction.NewTransactionService(transactionRepo)

	// init pool worker
	pool := work.NewWorkerPool(transaction.TransactionService{}, cfg.WorkerConcurrency, cfg.WorkerNamespace, redisPool)
	pool.JobWithOptions("transaction_settlement", work.JobOptions{Priority: 1, MaxFails: 3}, trxSvc.Create)

	// Start processing jobs
	pool.Start()

	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	// Stop the pool
	pool.Stop()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
