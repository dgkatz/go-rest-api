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
	assignmentRoutes := router.Group("/assignments")
	{
		assignmentRoutes.GET("/", controllers.ListAssignments)
		assignmentRoutes.POST("/", controllers.CreateAssignment)
		assignmentRoutes.GET("/:id", controllers.GetAssignment)
		assignmentRoutes.PATCH("/:id", controllers.UpdateAssignment)
		assignmentRoutes.DELETE("/:id", controllers.DeleteAssignment)
	}
}

func connectServices() {
	services.ConnectDB()
}