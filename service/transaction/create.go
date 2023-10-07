package transaction

import (
	"context"
	"sample-tabungan2/internal/repository"
	"time"

	"github.com/gocraft/work"
)

func (svc *TransactionService) Create(job *work.Job) error {
	noRekening := job.ArgString("no_rekening")
	nominal := job.ArgFloat64("nominal")
	publishedAt := job.ArgString("published_at")
	transactionType := job.ArgString("transaction_type")
	createdDate, _ := time.Parse(time.RFC3339, publishedAt)

	err := svc.repo.Insert(context.Background(), repository.InsertTransactionParam{
		TransactionCode: transactionType,
		NoRekening:      noRekening,
		CreatedDate:     createdDate,
		Nominal:         nominal,
	})
	return err
}
