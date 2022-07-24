package middleware

import (
	"registration/services"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(403, gin.H{"error":"request does not contain an access token"})
			ctx.Abort()
			return
		}

		if !strings.Contains(tokenString, "Bearer") {
			ctx.JSON(403, gin.H{"error":"request does not contain an access token"})
			ctx.Abort()
			return
		}

		tokenSigned := strings.Replace(tokenString, "Bearer ", "", -1)
		err := services.ValidateToken(tokenSigned)
		if err != nil {
			ctx.JSON(403, gin.H{"error":err.Error()})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}