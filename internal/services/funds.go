package services

import (
	"github.com/boises-finest-dao/investmentdao-backend/internal/models"
)

type Fund struct {
	ID          uint              `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Bot         *models.Bots      `json:"bot"`
	Exchanges   *models.Exchanges `json:"exchanges"`
}

func GetFundDetails(containerID string) *Fund {

	return nil
}
