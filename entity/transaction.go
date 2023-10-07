package entity

import "time"

type TransactionType string

var (
	TransactionTabung TransactionType = "C"
	TransactionTarik  TransactionType = "D"
)

type Transaction struct {
	ID              int64
	NoRekening      string          `json:"no_rekening"`
	CreatedDate     time.Time       `json:"waktu"`
	TransactionCode TransactionType `json:"kode_transaksi"`
	Nominal         float64         `json:"nominal"`
}
