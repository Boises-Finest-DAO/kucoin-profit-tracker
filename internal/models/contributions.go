package models

import "gorm.io/gorm"

type Contributions struct {
	gorm.Model
	ExchangeID uint    `json:"exchange_id"`
	Currency   string  `json:"currency"`
	Amount     float64 `json:"amount"`
	UsdValue   float64 `json:"usd_value"`
}
