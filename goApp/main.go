package main

import (
	"fmt"
	"log"

	"github.com/Carlfred3010/goApp/database"
	"github.com/Carlfred3010/goApp/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Hello, welcome to my awesome api!")

}

func initDatabase() {
	var err error

	database.DBConn, err = gorm.Open("sqlite3", "users.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database successfully connected")
	database.DBConn.AutoMigrate(&models.User{})
	fmt.Println("Database migrated")

}

func main() {

	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()
	app.Get("/", welcome)

	log.Fatal(app.Listen(":3000"))

}
