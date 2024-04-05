package db

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
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
	// подключение env-файла
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	// подключение БД
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	database = db
}
