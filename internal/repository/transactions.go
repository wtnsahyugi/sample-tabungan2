package repository

import (
	"context"
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
