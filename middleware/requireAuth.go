package middleware

import (
	"fmt"
	"miroirapiauth/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// Get the token

	tokenString := string(c.Request.Header.Get("Authorization")[7:])

	// Validate the token

	token, err := services.GetToken(tokenString)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		return
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		return
	}

	c.Next()
}
