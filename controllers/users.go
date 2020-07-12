package controllers

import (
	"github.com/dgkatz/go-rest-api/models"
	"github.com/dgkatz/go-rest-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CreateUserSerializer struct {
	FirstName  string `json:"first_name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
}

type GetUserSerializer struct {
	ID uint `json:"id" binding:"required"`
}

type UpdateUserSerializer struct {
	FirstName  string `json:"first_name"`
	LastName string `json:"last_name"`
}

type DeleteUserSerializer struct {
	ID uint `json:"id" binding:"required"`
}


func ListUsers(c *gin.Context) {
	var users []models.User
	services.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUser(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	if err := services.DB.Where(&models.User{ID: uint(user_id)}).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func CreateUser(c *gin.Context) {
	var input CreateUserSerializer
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// create user
	user := models.User{FirstName: input.FirstName, LastName: input.LastName}
	services.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Param("id"))

	var input UpdateUserSerializer
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := services.DB.Where(&models.User{ID: uint(user_id)}).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	services.DB.Model(&user).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	if err := services.DB.Where(&models.User{ID: uint(user_id)}).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	services.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}