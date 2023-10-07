package entity

type User struct {
	ID         int64
	Name       string
	NIK        string
	NoHP       string
	NoRekening string
	Saldo      *float64
}

func (u *User) AddSaldo(nominal float64) {
	if u.Saldo == nil {
		u.Saldo = &nominal
		return
	}

	newSaldo := *u.Saldo + nominal
	u.Saldo = &newSaldo
}

func (u *User) SubtractSaldo(nominal float64) error {
	if u.Saldo == nil {
		return ErrSaldoZero
	}

	if *u.Saldo < nominal {
		return ErrSaldoLessThanWithrawNominal
	}

	newSaldo := *u.Saldo - nominal
	u.Saldo = &newSaldo

	return nil
}
