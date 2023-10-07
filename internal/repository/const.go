package repository

const (
	insertUser = `
	INSERT INTO users (nama, nik, no_hp, no_rekening) VALUES($1, $2, $3, $4);
	`

	getUserByNikAndNoHP = `
		SELECT id, nama, nik, no_hp, no_rekening, latest_saldo FROM users where nik = $1 AND no_hp = $2;
	`
)
