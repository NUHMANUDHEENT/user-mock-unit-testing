package routers

import (
	"userPage/controller"
	"userPage/database"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	db := database.InitDatabase()
	router.POST("/signup", controller.SignUp(db))
}
