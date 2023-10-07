package transaction

import (
	"context"
	"sample-tabungan2/entity"
)

func (svc *TransactionService) ListByRekening(ctx context.Context, noRekening string) ([]entity.Transaction, error) {
	data, err := svc.repo.ListByNoRekening(ctx, noRekening)
	if err != nil {

		return nil, err
	}

	if len(data) == 0 {
		return nil, entity.ErrNoRekeningNotFound
	}

	return data, nil
}
