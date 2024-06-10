package database

import (
	"log"
	"userPage/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() *gorm.DB {
	db := createDB()
	autoMigrate(db)
	return db
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Users{})
}

func createDB() *gorm.DB {
	DSN := "host=localhost user=postgres password=Nuhman@456 dbname=postgres port=5432"
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}


func SetDB(mockDB *gorm.DB) {
	DB = mockDB
}
