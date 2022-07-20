package models

type Funds struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	BotID       uint   `json:"bot_id"`
	Name        string `json:"name"`
	Description string `json:"host"`
}
