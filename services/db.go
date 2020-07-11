package services

import (
	"github.com/dgkatz/learn_go_api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open("postgres", "host=localhost port=5432 user=go_user dbname=gorm password=123 sslmode=disable")

	if err != nil {
		log.Print(err)
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.User{}, &models.Assignment{})

	DB = database
}