package entity

import (
	"fmt"
	"taskmaster/db"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title  string `json: "title"`
	UserID uint   `json: "uid"`
	User   User   `gorm:"foreignkey:UserID"`
}

func (i Task) TableName() string {
	return "tasks"
}

func MigrateTask() {
	fmt.Println("Task migrate")
	db.DB().AutoMigrate(&Task{})
}
