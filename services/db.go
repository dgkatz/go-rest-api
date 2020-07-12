package services

import (
	"fmt"
	"github.com/dgkatz/go-rest-api/config"
	"github.com/dgkatz/go-rest-api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

const dbConnectionStringTemplate string = "host=%v port=%v user=%v dbname=%v password=%v sslmode=disable"

var DB *gorm.DB


func constructConnectionString() string {
	return fmt.Sprintf(
		dbConnectionStringTemplate,
		config.Config.DB.Host,
		config.Config.DB.Port,
		config.Config.DB.User,
		config.Config.DB.Name,
		config.Config.DB.Password)
}

func ConnectDB() {
	database, err := gorm.Open("postgres", constructConnectionString())

	if err != nil {
		log.Print(err)
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.User{}, &models.Assignment{})

	DB = database
}