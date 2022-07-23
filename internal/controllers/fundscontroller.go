package controllers

import (
	"net/http"
	"strconv"

	"github.com/boises-finest-dao/investmentdao-backend/internal/database"
	"github.com/boises-finest-dao/investmentdao-backend/internal/models"
	"github.com/boises-finest-dao/investmentdao-backend/internal/services"
	"github.com/gin-gonic/gin"
)

type ExchangeAPI struct {
	MasterExchange uint   `json:"master_exchange"`
	ExchangeType   uint   `json:"exchange_type"`
	ApiKey         string `json:"api_key"`
	ApiSecret      string `json:"api_secret"`
	ApiPass        string `json:"api_pass"`
}

type BotContainer struct {
	ContainerID string `json:"container_id"`
	Host        string `json:"host"`
	Port        string `json:"port"`
}

func CreateFund(context *gin.Context) {
	var fund models.Fund

	if context.ShouldBind(&fund) == nil {
		result := database.Instance.Create(&fund)
		if result.Error != nil {
			context.Status(http.StatusBadRequest)
			return
		}

		context.JSON(http.StatusOK, fund)
	} else {
		context.Status(http.StatusBadRequest)
		return
	}
}

func AttachBot(context *gin.Context) {
	var bot BotContainer
	var fund *models.Fund
	fundId := context.Param("fundId")

	if err := context.ShouldBindJSON(&bot); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	database.Instance.First(&fund, fundId)

	fund.BotContainerID = services.EncryptString(bot.ContainerID)
	fund.BotHost = services.EncryptString(bot.Host)
	fund.BotPort = services.EncryptString(bot.Port)

	record := database.Instance.Save(&fund)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
}

func AddExchange(context *gin.Context) {
	var body ExchangeAPI
	fundId := context.Param("fundId")
	fund_id, _ := strconv.ParseUint(fundId, 10, 8)

	if context.ShouldBind(&body) == nil {
		var exchange *models.Exchange

		exchange.MasterexchangeID = body.MasterExchange
		exchange.FundID = uint(fund_id)
		exchange.ExchangelistID = body.ExchangeType
		exchange.ApiKey = services.EncryptString(body.ApiKey)
		exchange.APISecret = services.EncryptString(body.ApiSecret)
		exchange.APIPassPhrase = services.EncryptString(body.ApiPass)

		result := database.Instance.Create(&exchange)
		if result.Error != nil {
			context.Status(http.StatusBadRequest)
			return
		}

		context.JSON(http.StatusOK, exchange)
	} else {
		context.Status(http.StatusBadRequest)
		return
	}
}

func UpdateApiKey(context *gin.Context) {
	var exchangeApi ExchangeAPI

	if context.ShouldBind(&exchangeApi) == nil {
		exchangeId := context.Param("exchangeId")

		var exchange *models.Exchange
		database.Instance.First(&exchange, exchangeId)

		exchange.ApiKey = services.EncryptString(exchangeApi.ApiKey)
		exchange.APISecret = services.EncryptString(exchangeApi.ApiSecret)
		exchange.APIPassPhrase = services.EncryptString(exchangeApi.ApiPass)

		database.Instance.Save(&exchange)

		context.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	} else {
		context.Status(http.StatusBadRequest)
		return
	}
}
