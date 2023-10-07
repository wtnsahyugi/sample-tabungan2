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

type CreateRequest struct {
	Name string `json:"nama"`
	NIK  string `json:"nik"`
	NoHP string `json:"no_hp"`
}

func (svc *UserService) Create(ctx context.Context, req CreateRequest) (string, error) {
	data, err := svc.repo.GetByNikAndPhoneNumber(ctx, req.NIK, req.NoHP)
	if err != nil {
		return "", errors.New("[GetByNikAndPhoneNumber] error: " + err.Error())
	}

	if data.ID != 0 {
		return "", entity.ErrUniqueHpNik
	}

	accountNumber := ""
	i := 0
	for {
		// maximum retry generate account number
		if i > maxRetry {
			return "", errors.New("system error when generate account number")
		}

		accountNumber, err = generateAccountNumber(req.NIK, req.NoHP)
		if err != nil {
			return "", entity.ErrGenerateNoRekening
		}
		err := svc.repo.Insert(ctx, repository.InsertUserParam{
			Nama:       req.Name,
			NIK:        req.NIK,
			NoHP:       req.NoHP,
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
