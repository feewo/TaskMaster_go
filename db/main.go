package db

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB
var migrate = make([]func(), 0)

func Add(mF func()) {
	migrate = append(migrate, mF)
}
func DB() *gorm.DB {
	return database
}

func Migrate() {
	for _, f := range migrate {
		f()
	}
}

func init() {
	db, err := gorm.Open(sqlite.Open("TaskMaster.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	database = db
}
