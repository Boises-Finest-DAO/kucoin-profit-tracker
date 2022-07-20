package models

import "gorm.io/gorm"

type Bot struct {
	gorm.Model
	FundID      uint       `json:"fund_id"`
	ContainerID string     `json:"container_id"`
	Name        string     `json:"name"`
	Host        string     `json:"host"`
	Port        string     `json:"port"`
	Exchanges   []Exchange `json:"exchanges"`
}
