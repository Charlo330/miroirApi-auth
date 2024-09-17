package initializers

import (
	"miroirapiauth/models"

	"golang.org/x/crypto/bcrypt"
)

func SeedData() {

	users := []models.User{
		{Email: "max@test.com", Password: "12345678"},
		{Email: "charlo@test.com", Password: "12345678"},
		{Email: "oli@test.com", Password: "12345678"},
	}

	hash, err := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)

	for _, user := range users {

		if err != nil {
			panic("Error hashing pwd")
		}

		user.Password = string(hash)

		if err := DB.Create(&user).Error; err != nil {
			print("Error seeding users")
		}
	}
}
