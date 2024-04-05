package entity

import (
	"taskmaster/db"
	"time"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	UserID  uint      `json: "uid"`
	User    User      `gorm:"foreignkey:UserID"`
	Token   string    `json token`
	Expired time.Time `json expired`
}

func (i Token) TableName() string {
	return "tokens"
}

func MigrateToken() {
	db.DB().AutoMigrate(&Token{})
}

func init() {
	db.Add(MigrateToken)
}
