package repository

const (
	insertUser = `
	INSERT INTO users VALUES($1, $2, $3, $4);
	`
)
