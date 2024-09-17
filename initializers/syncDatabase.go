package initializers

import (
	"miroirapiauth/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
