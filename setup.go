package main

import (
	"github.com/dgkatz/learn_go_api/controllers"
	"github.com/dgkatz/learn_go_api/services"
)


func initializeRoutes() {
	router.GET("/status", controllers.Status)
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", controllers.ListUsers)
		userRoutes.POST("/", controllers.CreateUser)
		userRoutes.GET("/:id", controllers.GetUser)
		userRoutes.PATCH("/:id", controllers.UpdateUser)
		userRoutes.DELETE("/:id", controllers.DeleteUser)
	}
}

func connectServices() {
	services.ConnectDB()
}