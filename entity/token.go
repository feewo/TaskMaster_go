package entity

import (
	"taskmaster/db"
	"time"
)

type Token struct {
	Tokid   uint32    `json tokid sql:"AUTO_INCREMENT" gorm:"primaryKey;autoIncrement"`
	Iid     uint32    `json iid gorm:"index:idx_item_iid`
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
