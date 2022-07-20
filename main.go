package main

import (
	"log"

	"github.com/boises-finest-dao/investmentdao-backend/internal/database"
	"github.com/boises-finest-dao/investmentdao-backend/internal/exchanges/kucoin"
	"github.com/boises-finest-dao/investmentdao-backend/internal/models"
)

func main() {
	// Load Configurations using Viper
	LoadAppConfig()

	// Initialize Database
	database.Connect(AppConfig.GormConnection)
	database.Migrate()

	tradingBalance := models.TradingBalance{
		FundID: 1,
	}

	result := database.Instance.Create(&tradingBalance)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	ks := kucoin.ConfigureConnection(AppConfig.KuCoinApiKey, AppConfig.KuCoinApiSecret, AppConfig.KuCoinApiPass)

	balance := ks.GetTradingBalances(tradingBalance.ID)

	result = database.Instance.Create(&balance)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	tradingBalance.TotalTradingBalance += balance.TotalExchangeBalance

	database.Instance.Save(&tradingBalance)

	// bot := services.GetBotByID(1)

	// log.Println(services.EncryptString("123456789"))
}
