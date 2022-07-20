package models

import "gorm.io/gorm"

type Exchange struct {
	gorm.Model
	BotID         uint   `json:"bot_id"`
	Name          string `json:"name" gorm:"name"`
	ApiKey        []byte `json:"api_key"`
	APISecret     []byte `json:"api_secret"`
	APIPassPhrase []byte `json:"api_pass"`
}
