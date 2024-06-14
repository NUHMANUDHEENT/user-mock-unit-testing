package routers

import (
	"userPage/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/signup", controller.Signup)
}
