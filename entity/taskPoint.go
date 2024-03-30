package entity

import "taskmaster/db"

type TaskPoint struct {
	Pid   uint32 `json: "pid" sql:"AUTO_INCREMENT" gorm:"primaryKey"`
	Title string `json: "title"`
	Ready bool   `json: "ready"`
	Tid   uint32 `json: "tid" gorm:"index:idx_taskpoint_tid`
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
