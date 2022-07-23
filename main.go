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
	ginserver "github.com/go-oauth2/gin-server"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	oauth2models "github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/procyon-projects/chrono"
	"gorm.io/gorm/clause"
)

var (
	dumpvar   bool
	idvar     string
	secretvar string
	domainvar string
	portvar   int
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
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// generate jwt access token
	// manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))
	manager.MapAccessGenerate(generates.NewAccessGenerate())

	clientStore := store.NewClientStore()
	clientStore.Set(idvar, &oauth2models.Client{
		ID:     idvar,
		Secret: secretvar,
		Domain: domainvar,
	})
	manager.MapClientStorage(clientStore)

	// Initialize the oauth2 service
	ginserver.InitServer(manager)
	ginserver.SetPasswordAuthorizationHandler(func(ctx context.Context, clientID, username, password string) (userID string, err error) {
		if username == "test" && password == "test" {
			userID = "test"
		}
		return
	})
	ginserver.SetUserAuthorizationHandler(userAuthorizeHandler)
	ginserver.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})
	ginserver.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})
	ginserver.SetAllowGetAccessRequest(true)
	ginserver.SetClientInfoHandler(server.ClientFormHandler)

	router := gin.Default()
	api := router.Group("/api")
	{
		// Get Token Route
		api.POST("/token", controllers.GenerateToken)

		// Admin Routes
		admin := api.Group("/admin").Use(middlewares.Auth()).Use(middlewares.IsAdmin())
		{
			admin.POST("/register-user", controllers.RegisterUser)
			admin.POST("/add-exchange", controllers.AddSupportedExchange)
		}

		//Funds Routes
		fund1 := api.Group("/funds").Use(middlewares.Auth())
		{
			fund1.GET("/", controllers.Ping)
			fund1.POST("/create", controllers.CreateFund)
		}

		fund2 := api.Group("/funds/:fundId").Use(middlewares.Auth()).Use((middlewares.FundUser()))
		{
			fund2.POST("/bot/exchanges/add", controllers.AddExchange)
			fund2.POST("/bot/exchanges/:exchangeId/update", controllers.UpdateApiKey)
			fund2.POST("/bot/attach", controllers.AttachBot).Use(middlewares.IsAdmin())
		}

		// User Routers
		user := api.Group("/user").Use(middlewares.Auth())
		{
			user.GET("/", controllers.Ping)
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
		fundsResult := database.Instance.Preload(clause.Associations).Find(&funds)
		for _, fund := range *funds {
			tradingBalance := models.TradingBalance{
				FundID: fund.ID,
			}

			result := database.Instance.Create(&tradingBalance)
			if result.Error != nil {
				log.Fatal(result.Error)
			}

			for _, exchange := range fund.Exchanges {
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
