package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddExchange(context *gin.Context) {
	userId := context.MustGet("user").(string)
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"userId":  userId,
	})
}
