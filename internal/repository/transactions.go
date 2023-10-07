package repository

import (
	"context"
	"sample-tabungan2/entity"
	"time"
)

type TransactionRepository struct {
	db PgxPoolIface
}

func NewTransactionRepository(db PgxPoolIface) *TransactionRepository {
	return &TransactionRepository{db: db}
}

type InsertTransactionParam struct {
	TransactionCode string    `db:"transaction_code"`
	NoRekening      string    `db:"no_rekening"`
	CreatedDate     time.Time `db:"created_date"`
	Nominal         float64   `db:"nominal"`
}

func (q *TransactionRepository) Insert(ctx context.Context, args InsertTransactionParam) error {
	_, err := q.db.Exec(ctx, insertTransaction,
		args.TransactionCode,
		args.NoRekening,
		args.CreatedDate,
		args.Nominal,
	)
	return err
}

func (q *TransactionRepository) ListByNoRekening(ctx context.Context, noRekening string) ([]entity.Transaction, error) {
	rows, err := q.db.Query(ctx, getTransactionByRekening, noRekening)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []entity.Transaction
	for rows.Next() {
		var i entity.Transaction
		if err := rows.Scan(&i.ID, &i.TransactionCode, &i.NoRekening, &i.CreatedDate, &i.Nominal); err != nil {
			return nil, err
		}
		items = append(items, i)
	}

	return items, nil
}
