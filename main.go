package main

import (
	"github.com/bapply/biparty/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/v1/user", controllers.CreateUser)
	router.Run()
}
