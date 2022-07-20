package middlewares

import (
	"strconv"

	"github.com/boises-finest-dao/investmentdao-backend/internal/database"
	"github.com/boises-finest-dao/investmentdao-backend/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func FundUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		fundId := context.Param("fundId")
		userId := context.MustGet("user")

		//Get Fund
		var fund *models.Fund
		database.Instance.Preload(clause.Associations).First(&fund, fundId)

		//Get User
		var user *models.User
		database.Instance.Preload(clause.Associations).First(&user, userId)

		//Check User is has access to Fund
		var hasAccess = false
		for _, fund := range user.Funds {
			fund_id, _ := strconv.ParseUint(fundId, 10, 8)
			if fund.ID == uint(fund_id) {
				hasAccess = true
			}
		}

		if !hasAccess {
			context.JSON(401, gin.H{"error": "user does not have permission to access this resource"})
			context.Abort()
			return
		}

		context.Next()
	}
}
