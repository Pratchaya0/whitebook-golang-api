package middlewares

import (
	"net/http"
	"strings"

	"github.com/Pratchaya0/whitebook-golang-api/services"
	"github.com/gin-gonic/gin"
)

// validate token
func Authorizes() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientToken := ctx.Request.Header.Get("Authorization")
		if clientToken == "" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "No Authorization header provided"})
			return
		}

		extractedToken := strings.Split(clientToken, "Bearer ")
		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Incorrect Format of Authorization Token"})
			return
		}

		jwtWrapper := services.JwtWrapper{
			SecretKey: "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
			Issuer:    "AuthService",
		}

		_, err := jwtWrapper.ValidateToken(clientToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.Next()
	}
}
