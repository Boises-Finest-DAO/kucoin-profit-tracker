package models

import "gorm.io/gorm"

type Exchange struct {
	gorm.Model
	MasterexchangeID uint         `json:"masterexchange_id"`
	FundID           uint         `json:"fund_id"`
	ExchangelistID   uint         `json:"exchange_type"`
	Exchangelist     Exchangelist `json:"exchange_info"`
	ApiKey           []byte       `json:"api_key"`
	APISecret        []byte       `json:"api_secret"`
	APIPassPhrase    []byte       `json:"api_pass"`
}

type Masterexchange struct {
	gorm.Model
	ExchangelistID uint         `json:"exchange_type"`
	Exchangelist   Exchangelist `json:"exchange_info"`
	Email          string       `json:"email"`
	ApiKey         []byte       `json:"api_key"`
	APISecret      []byte       `json:"api_secret"`
	APIPassPhrase  []byte       `json:"api_pass"`
	SubAccounts    []Exchange   `json:"sub_accounts"`
}

type Exchangelist struct {
	gorm.Model
	Name string `json:"name"`
}
