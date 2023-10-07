package entity

type User struct {
	ID         int64
	Name       string
	NIK        string
	NoHP       string
	NoRekening string
	Saldo      *float64
}
