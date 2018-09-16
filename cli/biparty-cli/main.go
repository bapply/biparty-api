package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/bapply/biparty/helpers"
	"github.com/bapply/biparty/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("please provide a subcommand")
	}
	switch flag.Args()[0] {
	case "migrate":
		migrate()
	}
}

func migrate() {
	// Connect to the DB
	db := helpers.ConnectDB()

	// Models
	db.AutoMigrate(&models.User{})

	fmt.Println("All models migrated successfully")
}
