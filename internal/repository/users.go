package repository

import "context"

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
