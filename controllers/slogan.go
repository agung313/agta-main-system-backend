package controllers

import (
	"github.com/agung313/agta-main-system-backend/config"
	"github.com/agung313/agta-main-system-backend/models"
	"github.com/gofiber/fiber/v2"
)

func GetSlogan(c *fiber.Ctx) error {
	var slogan []models.Slogan
	result := config.DB.Find(&slogan)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not find slogan",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Get slogan success",
		"data": map[string]interface{}{
			"firstText":   slogan[0].FirstText,
			"secondText":  slogan[0].SecondText,
			"thirdText":   slogan[0].ThirdText,
			"description": slogan[0].Description,
		},
	})
}

func CreateOrUpdateSlogan(c *fiber.Ctx) error {
	slogan := new(models.Slogan)
	if err := c.BodyParser(slogan); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}

	var existingSlogan models.Slogan
	result := config.DB.First(&existingSlogan)

	if result.Error != nil {
		// No existing slogan found, create a new one
		result = config.DB.Create(&slogan)
		if result.Error != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Could not create slogan",
				"error":   result.Error.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"message": "Create slogan success",
		})
	}

	// Existing slogan found, update it
	result = config.DB.Model(&existingSlogan).Updates(slogan)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not update slogan",
			"error":   result.Error.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Update slogan success",
	})
}

func DeleteSlogan(c *fiber.Ctx) error {
	var slogan models.Slogan
	if err := config.DB.First(&slogan).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Slogan not found",
		})
	}

	if err := config.DB.Unscoped().Delete(&slogan).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not delete slogan",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Delete slogan success",
	})
}
