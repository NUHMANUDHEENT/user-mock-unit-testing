package controller

import (
	"net/http"
	"userPage/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SignUp(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.Users
		if err := c.Bind(&user); err != nil {
			c.JSON(500, gin.H{
				"message": "binding error",
			})
			return
		}

		err := db.Create(&models.Users{
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		}).Error

		if err != nil {
			c.JSON(501, gin.H{
				"error": "failed created user",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "successfully created " + user.Name,
		})
	}
}
