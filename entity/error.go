package entity

import "errors"

const (
	ErrUniqueViolationCode = "23505"
)

var (
	ErrUniqueHpNik           = errors.New("user already registered")
	ErrGenerateNoRekening    = errors.New("error generate no rekening")
	ErrInvalidRequestPayload = errors.New("invalid request")
)
