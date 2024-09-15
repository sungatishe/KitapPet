package db

import (
	"awesomeProject/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var Db *gorm.DB

func InitDb() {
	err1 := godotenv.Load()
	if err1 != nil {
		log.Fatalf("Error loading .env file")
	}
	dsn := os.Getenv("dsn")
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect db")
	}

	Db.AutoMigrate(&models.Book{})
	Db.AutoMigrate(&models.User{})

}
