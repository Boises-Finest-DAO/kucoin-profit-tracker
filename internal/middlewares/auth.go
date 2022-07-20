package middlewares

import (
	"strconv"

	"github.com/boises-finest-dao/investmentdao-backend/internal/auth"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}

		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		context.Set("user", strconv.FormatUint(uint64(claims.ID), 10))

		context.Next()
	}
}
