package services

import (
	"log"

	"github.com/boises-finest-dao/investmentdao-backend/internal/database"
	"github.com/boises-finest-dao/investmentdao-backend/internal/models"
)

func GetBotDetails(containerID string) models.Bots {
	var bot models.Bots
	if err := database.Instance.Where("container_id = ?", containerID).First(&bot); err != nil {
		// TODO: handle error
		log.Println(err)
	}

	return bot
}

func GetBotByID(botID uint) models.Bots {
	var bot models.Bots
	if err := database.Instance.Where("id = ?", botID).First(&bot); err != nil {
		// TODO: handle error
		log.Println(err)
	}

	return bot
}
