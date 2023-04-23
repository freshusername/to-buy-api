package api

import (
	"github.com/freshusername/to-buy-api/database"
	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {
	table := "buy_items"
	check := database.DB.Db.Raw("select * from " + table + " limit 1;").Row()

	if check != nil {
		return c.Status(200).JSON("Healthy")
	}
	return c.Status(500).JSON("Unhealthy")
}
