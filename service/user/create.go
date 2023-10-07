package user

import (
	"context"
	"math/rand"
)

// define the given charset for no rekening
var charset = []byte("0123456789")

func (svc *UserService) Create(ctx context.Context, name, nik, noHp string) (string, error) {
	return "", nil
}

// n is the length of random string we want to generate
func randStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		// randomly select 1 character from given charset
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
