package entity

import "taskmaster/db"

type User struct {
	Iid      uint32 `json: "iid" gorm:"primaryKey"`
	Login    string `json: "login" gorm:"uniqueIndex:idx_user_login.1`
	Email    string `json: "email"`
	Password string `json: "password"`
}

func (i User) TableName() string {
	return "user"
}

func MigrateUser() {
	db.DB().AutoMigrate(&User{})
}

func init() {
	db.Add(MigrateUser)
}
