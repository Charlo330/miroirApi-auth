package controllers

import (
	"net/http"
	"os"
	"time"

	"miroirapiauth/repository"
	"miroirapiauth/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	// Get the email/pass
	var body struct {
		Email    string `validate:"required" json:"email"`
		Password string `validate:"required,max=200,min=6" json:"password"`
	}

	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	// Hash the pass

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})

		return
	}

	// Create the user in db
	user, err := repository.CreateUser(body.Email, string(hash))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "The email is already in use",
		})

		return
	}
	// Send response
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func Login(c *gin.Context) {
	// Get the email/pass

	var body struct {
		Email    string `validate:"required" json:"email"`
		Password string `validate:"required,max=200,min=8" json:"password"`
	}

	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	// Get the user from db
	user, err := repository.GetUserByEmail(body.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	// Compare the pass
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	// Generate JWT token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subs": user.ID,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign the token with our secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})

		return
	}

	// Send response

	c.JSON(http.StatusOK, gin.H{
		"user_id": user.ID,
		"token":   tokenString,
	})
}

func Validate(c *gin.Context) {

	tokenString := string(c.Request.Header.Get("Authorization")[7:])

	id, err := services.GetUserIdByToken(tokenString)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
