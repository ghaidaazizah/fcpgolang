package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {	
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, model.NewErrorResponse("authorization header is missing"))
			ctx.Abort()
			return
		}

		tokenString := strings.Split(authHeader, " ")
		if len(tokenString) != 2 || tokenString[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, model.NewErrorResponse("invalid token format"))
			ctx.Abort()
			return
		}

		userID := 1 

		ctx.Set("user_id", userID)

		ctx.Next()
	})
}
