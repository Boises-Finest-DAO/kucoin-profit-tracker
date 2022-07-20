package main

import (
	// "log"

	// "github.com/boises-finest-dao/investmentdao-backend/internal/database"
	// "github.com/boises-finest-dao/investmentdao-backend/internal/exchanges/kucoin"

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

	// kucoin.ConfigureConnection(AppConfig.KuCoinApiKey, AppConfig.KuCoinApiSecret, AppConfig.KuCoinApiPass)

	// kucoin.GetTradingBalances()

	// bot := helpers.GetBotDetails("c42a00d704fda9a62c54e15012f0dd0a994b4ab3c90c9185aa99e80edc931fd5")

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

	// log.Println(balance)

	// log.Println(services.EncryptString("123456789"))
}
