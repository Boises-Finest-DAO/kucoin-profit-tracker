package models

type Bot_Exchanges struct {
	ID         uint `json:"id" gorm:"primary_key"`
	BotID      uint `json:"bot_id"`
	ExchangeID uint `json:"exchange_id"`
	Enabled    bool `json:"enabled"`
}
