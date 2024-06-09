package main

import (
	"test/database"
	"test/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	database.InitMethod()
}
func main() {
	router := gin.Default()
	routers.SetupRouter(router)
	router.Run(":8081")
}
