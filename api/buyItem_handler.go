package api

import (
	"github.com/freshusername/to-buy-api/database"
	"github.com/freshusername/to-buy-api/models"
	"github.com/gofiber/fiber/v2"
)

func HandleCreateBuyItem(c *fiber.Ctx) error {
	buyItem := new(models.BuyItem)
	if err := c.BodyParser(buyItem); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&buyItem)

	return c.Status(200).JSON(buyItem)
}

func HandleGetBuyItems(c *fiber.Ctx) error {
	buyItems := []models.BuyItem{}
	database.DB.Db.Find(&buyItems)

	return c.Status(200).JSON(buyItems)
}
