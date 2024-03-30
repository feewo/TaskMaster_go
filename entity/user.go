package entity

import "taskmaster/db"

type User struct {
	Iid      uint32 `json: "iid" sql:"AUTO_INCREMENT" gorm:"primaryKey"`
	Login    string `json: "login" gorm:"uniqueIndex:idx_user_login.1`
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
