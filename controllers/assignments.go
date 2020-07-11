package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListAssignments(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "running"})
}

func GetAssignment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "running"})
}

func CreateAssignment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "running"})
}

func UpdateAssignment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "running"})
}

func DeleteAssignment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "running"})
}