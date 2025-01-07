package controllers

import (
	"github.com/agung313/agta-main-system-backend/config"
	"github.com/agung313/agta-main-system-backend/models"
	"github.com/gofiber/fiber/v2"
)

func CreateVisitor(c *fiber.Ctx) error {
	visitor := new(models.Visitor)
	if err := c.BodyParser(visitor); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid input",
		})
	}

	result := config.DB.Create(visitor)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not send visitor",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Visitor sent",
	})
}

func DeleteAllVisitors(c *fiber.Ctx) error {
	result := config.DB.Unscoped().Where("1 = 1").Delete(&models.Visitor{})
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": result.Error,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success permanently deleted all visitors",
		"total":   result.RowsAffected,
	})
}
