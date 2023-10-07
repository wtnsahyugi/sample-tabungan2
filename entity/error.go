package entity

import "errors"

const (
	ErrUniqueViolationCode = "23505"
)

var (
	ErrUniqueHpNik                 = errors.New("user already registered")
	ErrGenerateNoRekening          = errors.New("error generate no rekening")
	ErrNoRekeningNotFound          = errors.New("no rekening not found")
	ErrSaldoZero                   = errors.New("no rekening has zero saldo")
	ErrSaldoLessThanWithrawNominal = errors.New("cannot withdraw more than saldo")
	ErrInvalidRequestPayload       = errors.New("invalid request")
)
