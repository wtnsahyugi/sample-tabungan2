package user

import (
	"context"
	"sample-tabungan2/entity"

	"github.com/jackc/pgx/v4"
)

type WithdrawRequest struct {
	Nominal    float64 `json:"nominal"`
	NoRekening string  `json:"no_rekening"`
}

func (svc *UserService) Withdraw(ctx context.Context, req WithdrawRequest) (float64, error) {
	data, err := svc.repo.GetByNoRekening(ctx, req.NoRekening)
	if err != nil {
		// ignore if data not found
		if err == pgx.ErrNoRows {
			return float64(0), entity.ErrNoRekeningNotFound
		}

		return float64(0), err
	}

	if errSubtract := data.SubtractSaldo(req.Nominal); errSubtract != nil {
		return float64(0), errSubtract
	}

	if err := svc.repo.UpdateSaldo(ctx, req.NoRekening, *data.Saldo); err != nil {
		return float64(0), err
	}

	return *data.Saldo, nil
}
