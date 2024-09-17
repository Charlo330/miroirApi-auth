package services

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GetToken(tokenString string) (token *jwt.Token, err error) {
	// Validate the token

	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	return
}

func GetUserIdByToken(tokenString string) (int, error) {

	token, err := GetToken(tokenString)

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, exists := claims["subs"].(float64)
		if !exists {
			return 0, fmt.Errorf("User ID not found in claims")
		}

		return int(userID), nil
	}

	return 0, fmt.Errorf("Invalid token")
}
