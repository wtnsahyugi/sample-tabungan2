package repository

const (
	insertUser = `
		INSERT INTO users (nama, nik, no_hp, no_rekening) VALUES($1, $2, $3, $4);
	`

	getUserByNikAndNoHP = `
		SELECT id, nama, nik, no_hp, no_rekening, latest_saldo FROM users where nik = $1 AND no_hp = $2;
	`

	getUserByNoRekening = `
		SELECT id, nama, nik, no_hp, no_rekening, latest_saldo FROM users where no_rekening = $1;
	`

	updateUserSaldoByRekening = `
		UPDATE users SET latest_saldo = $1 WHERE no_rekening = $2;
	`

	insertTransaction = `
		INSERT INTO transactions (transaction_code, no_rekening, created_date, nominal) VALUES($1, $2, $3, $4);
	`

	getTransactionByRekening = `
		SELECT id, transaction_code, no_rekening, created_date, nominal FROM transactions WHERE no_rekening = $1;
	`
)
