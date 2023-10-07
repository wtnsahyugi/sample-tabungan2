package repository

import (
	"context"
	"sample-tabungan2/entity"

	"github.com/jackc/pgx/v4"
)

type UserRepository struct {
	db PgxPoolIface
}

func NewUserRepository(db PgxPoolIface) *UserRepository {
	return &UserRepository{db: db}
}

type InsertUserParam struct {
	Nama       string `db:"nama"`
	NIK        string `db:"nik"`
	NoHP       string `db:"no_hp"`
	NoRekening string `db:"no_rekening"`
}

func (q *UserRepository) Insert(ctx context.Context, args InsertUserParam) (int64, error) {
	row := q.db.QueryRow(ctx, insertUser,
		args.Nama,
		args.NIK,
		args.NoHP,
		args.NoRekening,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

func (q *UserRepository) GetByNikAndPhoneNumber(ctx context.Context, nik, phoneNumber string) (entity.User, error) {
	var result entity.User
	row := q.db.QueryRow(ctx, getUserByNikAndNoHP, nik, phoneNumber)
	if err := row.Scan(&result.ID, &result.Name, &result.NIK, &result.NoHP, &result.NoRekening, &result.Saldo); err != nil {
		// ignore if data not found
		if err == pgx.ErrNoRows {
			return entity.User{}, nil
		}

		return entity.User{}, err
	}

	return result, nil
}
