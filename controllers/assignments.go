package controllers

import (
	"github.com/dgkatz/go-rest-api/models"
	"github.com/dgkatz/go-rest-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type AssignmentSerializer struct {
	UserID  uint `json:"user_id" binding:"required"`
	Title string `json:"title"`
	Description string `json:"title"`
	Deadline time.Time `json:"deadline"`
	Status string `json:"status"`
}


func ListAssignments(c *gin.Context) {
	var assignments []models.Assignment
	services.DB.Find(&assignments)
	c.JSON(http.StatusOK, gin.H{"data": assignments})
}

func GetAssignment(c *gin.Context) {
	assignment_id, _ := strconv.Atoi(c.Param("id"))
	var assignment models.Assignment
	if err := services.DB.Where(&models.Assignment{ID: uint(assignment_id)}).First(&assignment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Assignment not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": assignment})
}

func CreateAssignment(c *gin.Context) {
	var input AssignmentSerializer
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := services.DB.Where(&models.User{ID: input.UserID}).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	if input.Status == "" {
		input.Status = "TODO"
	}
	// create assignment
	assignment := models.Assignment{
		User: user,
		Title: input.Title,
		Description: input.Description,
		Deadline: input.Deadline,
		Status: input.Status,
	}
	services.DB.Create(&assignment)
	c.JSON(http.StatusOK, gin.H{"data": assignment})
}

func UpdateAssignment(c *gin.Context) {
	assignment_id, _ := strconv.Atoi(c.Param("id"))

	var input AssignmentSerializer
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var assignment models.Assignment
	if err := services.DB.Where(&models.Assignment{ID: uint(assignment_id)}).First(&assignment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Assignment not found!"})
		return
	}

	services.DB.Model(&assignment).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": assignment})
}

func DeleteAssignment(c *gin.Context) {
	assignment_id, _ := strconv.Atoi(c.Param("id"))
	var assignment models.Assignment
	if err := services.DB.Where(&models.Assignment{ID: uint(assignment_id)}).First(&assignment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Assignment not found!"})
		return
	}
	services.DB.Delete(&assignment)

	c.JSON(http.StatusOK, gin.H{"data": true})
}