package main

import (
	"fmt"
	"github.com/dgkatz/go-rest-api/config"
	"github.com/dgkatz/go-rest-api/services"
	"github.com/gin-gonic/gin"
	"log"
)

var router *gin.Engine

func setup() {
	loadConfig()
	initializeRoutes()
	connectServices()
}

func main() {
	router = gin.Default()
	setup()
	bind := fmt.Sprintf("%v:%v",
		config.Config.SERVER.Host,
		config.Config.SERVER.Port)
	err := router.Run(bind)
	if err != nil{
		log.Fatal("Failed to start server")
	}
	defer services.DB.Close()
}