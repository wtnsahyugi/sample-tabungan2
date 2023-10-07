package user

import (
	"context"
	"log"
	"sample-tabungan2/entity"
	"time"

	"github.com/gocraft/work"
	"github.com/jackc/pgx/v4"
)

type DepositRequest struct {
	Nominal    float64 `json:"nominal"`
	NoRekening string  `json:"no_rekening"`
}

func (svc *UserService) Deposit(ctx context.Context, req DepositRequest) (float64, error) {
	data, err := svc.repo.GetByNoRekening(ctx, req.NoRekening)
	if err != nil {
		// ignore if data not found
		if err == pgx.ErrNoRows {
			return float64(0), entity.ErrNoRekeningNotFound
		}

		return float64(0), err
	}

	data.AddSaldo(req.Nominal)
	if err := svc.repo.UpdateSaldo(ctx, req.NoRekening, *data.Saldo); err != nil {
		return float64(0), err
	}

	_, err = svc.publisher.Enqueue("transaction_settlement", work.Q{
		"no_rekening":      req.NoRekening,
		"nominal":          req.Nominal,
		"published_at":     time.Now().Format(time.RFC3339),
		"transaction_type": entity.TransactionTabung,
	})
	if err != nil {
		// should persist error log when fail to publish data
		log.Println("error when publishing data: " + err.Error())
	}

	return *data.Saldo, nil
}
