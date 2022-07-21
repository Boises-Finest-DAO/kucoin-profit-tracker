package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/boises-finest-dao/investmentdao-backend/internal/controllers"
	"github.com/boises-finest-dao/investmentdao-backend/internal/database"
	"github.com/boises-finest-dao/investmentdao-backend/internal/exchanges/kucoin"
	"github.com/boises-finest-dao/investmentdao-backend/internal/middlewares"
	"github.com/boises-finest-dao/investmentdao-backend/internal/models"
	"github.com/boises-finest-dao/investmentdao-backend/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/procyon-projects/chrono"
	"gorm.io/gorm/clause"
)

func main() {
	// Load Configurations using Viper
	LoadAppConfig()

	// Initialize Database
	database.Connect(AppConfig.GormConnection)
	database.Migrate()

	// Start Balance Tacker
	startBalanceTracker()

	// Initialize Router
	router := initRouter()
	router.Run(fmt.Sprintf(":%v", AppConfig.ServerPort))

}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		// Get Token Route
		api.POST("/token", controllers.GenerateToken)

		// User Routers
		user := api.Group("/user")
		user.POST("/register", controllers.RegisterUser)
		userFunds := user.Group("/funds/:fundId").Use(middlewares.Auth()).Use((middlewares.FundUser()))
		{

			userFunds.POST("/bot/exchanges/add", controllers.AddExchange)
			userFunds.POST("/bot/exchanges/:exchangeId/update", controllers.UpdateApiKey)
		}

		// Other Secured Routes
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}

func startBalanceTracker() {
	taskScheduler := chrono.NewDefaultTaskScheduler()

	_, err := taskScheduler.ScheduleAtFixedRate(func(ctx context.Context) {
		var funds *[]models.Fund
		fundsResult := database.Instance.Preload("Bot.Exchanges").Preload(clause.Associations).Find(&funds)
		for _, fund := range *funds {
			tradingBalance := models.TradingBalance{
				FundID: fund.ID,
			}

			result := database.Instance.Create(&tradingBalance)
			if result.Error != nil {
				log.Fatal(result.Error)
			}

			for _, exchange := range fund.Bot.Exchanges {
				KuCoinApiKey := services.DecryptString(exchange.ApiKey)
				KuCoinApiSecret := services.DecryptString(exchange.APISecret)
				KuCoinApiPass := services.DecryptString(exchange.APIPassPhrase)

				ks := kucoin.ConfigureConnection(KuCoinApiKey, KuCoinApiSecret, KuCoinApiPass)

				balance := ks.GetTradingBalances(tradingBalance.ID)

				result = database.Instance.Create(&balance)
				if result.Error != nil {
					log.Fatal(result.Error)
				}

				tradingBalance.TotalTradingBalance += balance.TotalExchangeBalance
			}

			database.Instance.Save(&tradingBalance)
		}

		if fundsResult.Error != nil {
			log.Fatalln(fundsResult.Error.Error())
		}
	}, 5*time.Minute)

	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Print("Balance Tracker has been scheduled successfully.")
}
