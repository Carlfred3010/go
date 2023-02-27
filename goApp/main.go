package main

import (
	"log"

	"github.com/Carlfred3010/goApp/database"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Hello, welcome to my awesome api!")

}

func main() {

	database.ConnectDb()
	app := fiber.New()

	app.Get("/api", welcome)

	log.Fatal(app.Listen(":3000"))

}
