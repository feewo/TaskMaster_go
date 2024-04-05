package entity

import (
	"taskmaster/db"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Login    string `json: "login" gorm:"unique"`
	Email    string `json: "email"`
	Role     string `json: "role"`
	Password string `json: "password"`
}

func (i User) TableName() string {
	return "users"
}

func MigrateUser() {
	db.DB().AutoMigrate(&User{})
}

func init() {
	db.Add(MigrateUser)
}
