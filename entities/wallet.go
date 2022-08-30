package entities

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	UserId  int     `json:"user_id"`
	Balance float64 `json:"balance"`
	Status  string  `json:"status"`
}