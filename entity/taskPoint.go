package entity

import (
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
	db.DB().AutoMigrate(&TaskPoint{})
}

func init() {
	db.Add(MigrateTaskPoint)
}
