package main

import (
	"github.com/freshusername/to-buy-api/database"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.ConnectDb()

	app := fiber.New()
	setupCors(app)
	setupRoutes(app)

	app.Listen(":3000")
}
