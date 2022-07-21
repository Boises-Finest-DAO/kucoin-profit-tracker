package middlewares

import (
	"github.com/boises-finest-dao/investmentdao-backend/internal/database"
	"github.com/boises-finest-dao/investmentdao-backend/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func IsAdmin() gin.HandlerFunc {
	return func(context *gin.Context) {
		userId := context.MustGet("user")

		//Get User
		var user *models.User
		database.Instance.Preload(clause.Associations).First(&user, userId)

		//Check User is an admin
		if user.IsAdmin {
			context.Next()
		} else {
			context.JSON(401, gin.H{"error": "user does not have permission to access this resource"})
			context.Abort()
			return
		}

	}
}
