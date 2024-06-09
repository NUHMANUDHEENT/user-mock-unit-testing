package controller

import (
	"net/http"
	"test/database"
	"test/models"

	"github.com/gin-gonic/gin"
)

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var DB database.Database

func SignUpPost(c *gin.Context) {
	var request SignUpRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := DB.CreateUser(request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User signed up successfully"})
}
func SignUp(c *gin.Context) {
	var userData models.Users
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&userData).Error; err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"message": "User signed up successfully",
	})
}
