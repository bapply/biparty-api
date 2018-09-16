package controllers

import (
	"errors"

	"github.com/bapply/biparty/helpers"
	"github.com/bapply/biparty/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	db := helpers.ConnectDB()
	defer db.Close()
	user := models.User{}
	if !db.NewRecord(&user) {
		c.AbortWithError(409, errors.New("user already exist"))
	}
	db.Create(&user)
	c.JSON(201, &user)
}
