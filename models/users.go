package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `gorm:"type:varchar(100);omitempty" json:"-"`
}
