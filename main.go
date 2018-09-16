package main

import (
	"github.com/bapply/biparty/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/users", controllers.CreateUser)
	}
	router.Run()
}
