package models

import "gorm.io/gorm"

type Fund struct {
	gorm.Model
	BotID           uint             `json:"bot_id"`
	Name            string           `json:"name"`
	Description     string           `json:"host"`
	Bot             Bot              `json:"bots"`
	TradingBalances []TradingBalance `json:"trading_balances"`
	Users           []User           `json:"users" gorm:"many2many:user_funds"`
}
