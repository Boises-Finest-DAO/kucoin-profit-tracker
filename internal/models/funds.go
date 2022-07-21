package models

import "gorm.io/gorm"

type Fund struct {
	gorm.Model
	Name            string           `json:"name"`
	Description     string           `json:"description"`
	Exchanges       []Exchange       `json:"exchanges"`
	TradingBalances []TradingBalance `json:"trading_balances"`
	Users           []User           `json:"users" gorm:"many2many:user_funds"`
	BotContainerID  []byte           `json:"bot_container_id"`
	BotHost         []byte           `json:"bot_host"`
	BotPort         []byte           `json:"bot_port"`
}
