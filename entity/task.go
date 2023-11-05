package entity

import "taskmaster/db"

type Task struct {
	Tid   uint32 `json: "tid" gorm:"primaryKey"`
	Title string `json: "title"`
	Ready bool   `json: "ready"`
	Iid   uint32 `json: "iid" gorm:"index:idx_item_iid`
}

func (i Task) TableName() string {
	return "task"
}

func MigrateTask() {
	db.DB().AutoMigrate(&Task{})
}

func init() {
	db.Add(MigrateTask)
}
