package entity

import (
	"taskmaster/db"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Surname    string `gorm:"not null;type:varchar(100)"`
	Name       string `gorm:"not null;type:varchar(100)"`
	Patronymic string `gorm:"type:varchar(100)"`
	Login      string `gorm:"not null;type:varchar(100)"`
	Email      string `gorm:"not null;type:varchar(100)"`
	Role       string `gorm:"not null;type:varchar(100)"`
	Password   string `gorm:"not null;type:varchar(100)"`
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
