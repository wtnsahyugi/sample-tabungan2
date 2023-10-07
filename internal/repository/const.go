package repository

const (
	insertUser = `
	INSERT INTO users VALUES($1, $2, $3, $4);
	`

	getUserByNikAndNoHP = `
		SELECT id, name, nik, no_hp, no_rekening, last_saldo FROM users where nik = $1 AND no_hp = $2;
	`
)
