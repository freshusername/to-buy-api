package main

import (
	"fmt"
	"os"

	"github.com/freshusername/to-buy-api/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setupRoutes(app *fiber.App) {
	app.Get("/buyItems", api.HandleGetBuyItems)
	app.Post("/buyItems", api.HandleCreateBuyItem)
	app.Get("/", api.HealthCheck)
}

func setupCors(app *fiber.App) {
	hostUrl := os.Getenv("DB_HOST")
	originsList := fmt.Sprintf("http://localhost:5173, http://%s", hostUrl)

	app.Use(cors.New(cors.Config{
		AllowOrigins: originsList,
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
}
