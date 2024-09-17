package repository

import (
	"miroirapiauth/initializers"
	"miroirapiauth/models"
)

func CreateUser(email string, password string) (user models.User, err error) {
	user = models.User{Email: email, Password: password}
	if err := initializers.DB.Create(&user).Error; err != nil {
		panic("Error creating user")
	}

	return
}

func GetUserByEmail(email string) (user models.User, err error) {
	user = models.User{}
	if err := initializers.DB.Where("email = ?", email).First(&user).Error; err != nil {
		panic("Error getting user")
	}

	return
}
