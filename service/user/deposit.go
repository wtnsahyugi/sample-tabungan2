package user

import (
	"context"
	"sample-tabungan2/entity"

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

	return *data.Saldo, nil
}
