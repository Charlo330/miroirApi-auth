package initializers

import (
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	var dsn = os.Getenv("DB")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//DB.Migrator().DropTable(&models.User{})
	if err != nil {
		panic("Failed to connect to database!")
	}
}
