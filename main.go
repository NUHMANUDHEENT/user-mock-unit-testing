package main

import (
	"userPage/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routers.UserRoutes(router)
	router.Run(":8081")
}
