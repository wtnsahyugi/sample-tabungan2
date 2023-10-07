package user

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"sample-tabungan2/entity"
	"sample-tabungan2/internal/repository"

	"github.com/jackc/pgconn"
)

// define the given charset for no rekening
const (
	charSet             = "1234567890"
	accountNumberLength = 10
	maxRetry            = 3
)

func (svc *UserService) Create(ctx context.Context, name, nik, noHp string) (string, error) {
	data, err := svc.repo.GetByNikAndPhoneNumber(ctx, nik, noHp)
	if err != nil {
		return "", errors.New("[GetByNikAndPhoneNumber] error: " + err.Error())
	}

	if data.ID != 0 {
		return "", entity.ErrUniqueHpNik
	}

	var accountNumber string
	i := 0
	for {
		// maximum retry generate account number
		if i > maxRetry {
			return "", errors.New("system error when generate account number")
		}

		accountNumber, errGenerate := generateAccountNumber(nik, noHp)
		if errGenerate != nil {
			return "", entity.ErrGenerateNoRekening
		}
		_, err := svc.repo.Insert(ctx, repository.InsertUserParam{
			Nama:       name,
			NIK:        nik,
			NoHP:       noHp,
			NoRekening: accountNumber,
		})
		if err != nil {
			// if other account has same account number, system will regenerate account number and try to insert again
			if pqErr := err.(*pgconn.PgError); pqErr != nil && pqErr.Code == entity.ErrUniqueViolationCode {
				i++
				continue
			}

			return "", err
		}

		break
	}
	return accountNumber, nil
}

func generateAccountNumber(nik, phoneNumber string) (string, error) {
	bytes := make([]byte, accountNumberLength)
	copy(bytes[:], fmt.Sprintf("%s%s", nik, phoneNumber))
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = charSet[b%byte(len(charSet))]
	}
	return string(bytes), nil
}
