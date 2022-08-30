package entities

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	WalletId  int  `json:"wallet_id"`
	Amount float64 `json:"amount"`
	TransactionType string `json:"enums"`
}