package entity

import (
	"fmt"
	"taskmaster/db"

	"gorm.io/gorm"
)

type TaskPoint struct {
	gorm.Model
	Title  string `json: "title"`
	Ready  *bool  `json: "ready"`
	TaskID uint   `json: "tid"`
	Task   Task   `gorm:"foreignkey:TaskID"`
}

func (i TaskPoint) TableName() string {
	return "taskpoints"
}

func MigrateTaskPoint() {
	fmt.Println("TaskPoint migrate")
	db.DB().AutoMigrate(&TaskPoint{})
}
