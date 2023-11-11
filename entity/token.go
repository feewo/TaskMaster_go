package entity

import (
	"taskmaster/db"
	"time"
)

type Token struct {
	Tokid   uint32    `json tokid gorm:"primaryKey;autoIncrement"`
	Iid     uint32    `json iid gorm:"index:idx_item_iid`
	Token   string    `json token`
	Expired time.Time `json expired`
}

func (i Token) TableName() string {
	return "token"
}

func MigrateToken() {
	db.DB().AutoMigrate(&Token{})
}

func init() {
	db.Add(MigrateToken)
}
