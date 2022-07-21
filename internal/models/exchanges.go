package models

import "gorm.io/gorm"

type Exchange struct {
	gorm.Model
	FundID         uint         `json:"bot_id"`
	ExchangelistID uint         `json:"exchange_type"`
	Exchangelist   Exchangelist `json:"exchange_info"`
	ApiKey         []byte       `json:"api_key"`
	APISecret      []byte       `json:"api_secret"`
	APIPassPhrase  []byte       `json:"api_pass"`
}

type Exchangelist struct {
	gorm.Model
	Name string `json:"name"`
}
