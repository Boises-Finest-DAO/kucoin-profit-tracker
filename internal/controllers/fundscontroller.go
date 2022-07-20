package controllers

import (
	"net/http"

	"github.com/boises-finest-dao/investmentdao-backend/internal/database"
	"github.com/boises-finest-dao/investmentdao-backend/internal/models"
	"github.com/boises-finest-dao/investmentdao-backend/internal/services"
	"github.com/gin-gonic/gin"
)

type ExchangeAPI struct {
	ApiKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
	ApiPass   string `json:"api_pass"`
}

func AddExchange(context *gin.Context) {
	userId := context.MustGet("user").(string)
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"userId":  userId,
	})
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
