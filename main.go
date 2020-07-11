package main

import (
	"github.com/dgkatz/learn_go_api/services"
	"github.com/gin-gonic/gin"
	"log"
)

var router *gin.Engine

func setup() {
	initializeRoutes()
	connectServices()
}

func main() {
	router = gin.Default()
	setup()
	err := router.Run("0.0.0.0:8001")
	if err != nil{
		log.Fatal("Failed to start server")
	}
	defer services.DB.Close()
}