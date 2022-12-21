package models

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to connect .ENV")
	}
	DSN_DB := os.Getenv("DSN_DB")
	db, err := gorm.Open(mysql.Open(DSN_DB))
	if err != nil {
		log.Fatal("Failed to connect DB")
	}
	log.Println("DB sucessfull migrate")

	db.AutoMigrate(&Book{})

	DB = db
}
