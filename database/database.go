package database

import (
	"test/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// connecting to database for access table
func InitMethod() {
	dsn := "host=localhost user=postgres  password=Nuhman@456 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect daatabase")
	}
	DB = db
	//create table using struct
	DB.AutoMigrate(&models.Users{})
}
