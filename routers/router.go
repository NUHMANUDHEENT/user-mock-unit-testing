package routers

import (
	"test/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	router.POST("/signup", controller.SignUpPost)
	// router.POST("/signup", controller.SignUp)
}
