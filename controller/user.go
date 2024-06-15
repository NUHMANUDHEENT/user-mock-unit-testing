package controller

import (
	"net/http"
	"userPage/database"
	"userPage/models"

	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	var users []models.Users
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
