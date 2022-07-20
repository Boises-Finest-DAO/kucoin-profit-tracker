package main

import (
	"fmt"

	"github.com/boises-finest-dao/investmentdao-backend/internal/controllers"
	"github.com/boises-finest-dao/investmentdao-backend/internal/database"
	"github.com/boises-finest-dao/investmentdao-backend/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load Configurations using Viper
	LoadAppConfig()

	// Initialize Database
	database.Connect(AppConfig.GormConnection)
	database.Migrate()

	// Initialize Router
	router := initRouter()
	router.Run(fmt.Sprintf(":%v", AppConfig.ServerPort))

	// tradingBalance := models.TradingBalance{
	// 	FundID: 1,
	// }

	// result := database.Instance.Create(&tradingBalance)
	// if result.Error != nil {
	// 	log.Fatal(result.Error)
	// }

	// ks := kucoin.ConfigureConnection(AppConfig.KuCoinApiKey, AppConfig.KuCoinApiSecret, AppConfig.KuCoinApiPass)

	// balance := ks.GetTradingBalances(tradingBalance.ID)

	// result = database.Instance.Create(&balance)
	// if result.Error != nil {
	// 	log.Fatal(result.Error)
	// }

	// tradingBalance.TotalTradingBalance += balance.TotalExchangeBalance

	// database.Instance.Save(&tradingBalance)

	// bot := services.GetBotByID(1)

	// log.Println(services.EncryptString("123456789"))

	// var fund models.Fund
	// result := database.Instance.Preload("Bot.Exchanges").Preload(clause.Associations).First(&fund)
	// if result.Error != nil {
	// 	log.Fatal(result.Error)
	// }

	// fmt.Printf("fund.Bot.Name: %v\n", fund.Bot.Exchanges)
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
		}

		// Other Secured Routes
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
