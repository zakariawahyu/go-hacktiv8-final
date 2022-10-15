package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/zakariawahyu/go-hacktiv8-final/entity"
	"github.com/zakariawahyu/go-hacktiv8-final/exception"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func DatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	exception.PanicIfNeeded(errEnv)

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUsername, dbPassword, dbPort, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	exception.PanicIfNeeded(err)

	db.AutoMigrate(
		&entity.User{},
		&entity.Photo{},
		&entity.Comment{},
		&entity.SocialMedia{})

	fmt.Println("Database connected")
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	exception.PanicIfNeeded(err)
	dbSQL.Close()
}
