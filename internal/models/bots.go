package models

type Bots struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	ContainerID string `json:"container_id"`
	Name        string `json:"name"`
	Host        string `json:"host"`
	Port        string `json:"port"`
}
