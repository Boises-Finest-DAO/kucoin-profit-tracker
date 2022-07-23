package controllers

import (
	"net/http"

	"github.com/boises-finest-dao/investmentdao-backend/internal/database"
	"github.com/boises-finest-dao/investmentdao-backend/internal/models"
	"github.com/boises-finest-dao/investmentdao-backend/internal/services"

	"github.com/gin-gonic/gin"
)

type MasterExchange struct {
	ExchangeType  uint   `json:"exchange_type"`
	Email         string `json:"email"`
	ApiKey        string `json:"api_key"`
	APISecret     string `json:"api_secret"`
	APIPassPhrase string `json:"api_pass"`
}

func RegisterUser(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}

func AddSupportedExchange(context *gin.Context) {
	var exchange *models.Exchangelist

	if err := context.ShouldBindJSON(&exchange); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record := database.Instance.Create(&exchange)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, exchange)
}

func AddMasterExchange(context *gin.Context) {
	var body MasterExchange

	if context.ShouldBind(&body) == nil {
		var masterExchange *models.Masterexchange

		masterExchange.Email = body.Email
		masterExchange.ExchangelistID = body.ExchangeType
		masterExchange.ApiKey = services.EncryptString(body.ApiKey)
		masterExchange.APISecret = services.EncryptString(body.APISecret)
		masterExchange.APIPassPhrase = services.EncryptString(body.APIPassPhrase)

		result := database.Instance.Create(&masterExchange)
		if result.Error != nil {
			context.Status(http.StatusBadRequest)
			return
		}

		context.JSON(http.StatusOK, masterExchange)
	} else {
		context.Status(http.StatusBadRequest)
		return
	}
}
