package main

import (
	"log"
	"os"

	"github.com/bapply/biparty/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/users", controllers.CreateUser)
	}
	router.Run(os.Getenv("BIPARTY_PORT"))
}
