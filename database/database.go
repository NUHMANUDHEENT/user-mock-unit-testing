package database

import (
	"log"
	"userPage/models"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Mock sqlmock.Sqlmock

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Users{})
}

func CreateDB() {
	DSN := "host=localhost user=postgres password=Nuhman@456 dbname=postgres port=5432"
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	autoMigrate(db)
	DB = db
}

func SetDB(database *gorm.DB) {
	DB = database
}
