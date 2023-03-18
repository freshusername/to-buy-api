package main

import (
	"github.com/freshusername/to-buy-api/api"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/buyItems", api.HandleGetBuyItems)
	app.Post("/buyItems", api.HandleCreateBuyItem)
}
