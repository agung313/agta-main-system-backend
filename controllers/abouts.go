package controllers

import (
	"github.com/agung313/agta-main-system-backend/config"
	"github.com/agung313/agta-main-system-backend/models"
	"github.com/gofiber/fiber/v2"
)

func GetAbouts(c *fiber.Ctx) error {
	var abouts []models.About
	result := config.DB.Preload("ComitmentLists").Find(&abouts)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not find abouts",
		})
	}
	if len(abouts) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "No abouts found",
		})
	}
	comitmentLists := make([]map[string]interface{}, len(abouts[0].ComitmentLists))
	for i, comitment := range abouts[0].ComitmentLists {
		comitmentLists[i] = map[string]interface{}{
			"titleText":       comitment.TitleText,
			"descriptionText": comitment.DescriptionText,
		}
	}
	return c.JSON(fiber.Map{
		"message": "Get abouts success",
		"data": map[string]interface{}{
			"openingText":    abouts[0].OpeningText,
			"closingText":    abouts[0].ClosingText,
			"comitmentLists": comitmentLists,
		},
	})
}

func CreateOrUpdateAbouts(c *fiber.Ctx) error {
	abouts := new(models.About)
	if err := c.BodyParser(abouts); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}

	var existingAbouts models.About
	result := config.DB.Preload("ComitmentLists").First(&existingAbouts)

	if result.Error != nil {
		// No existing abouts found, create a new one
		result = config.DB.Create(&abouts)
		if result.Error != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Could not create abouts",
				"error":   result.Error.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"message": "Create abouts success",
		})
	}

	// Existing abouts found, update it
	result = config.DB.Model(&existingAbouts).Updates(abouts)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not update abouts",
			"error":   result.Error.Error(),
		})
	}

	// Update ComitmentLists
	if len(abouts.ComitmentLists) > 0 {
		// Delete existing ComitmentLists
		config.DB.Where("about_id = ?", existingAbouts.ID).Delete(&models.ComitmentList{})
		// Add new ComitmentLists
		for _, comitment := range abouts.ComitmentLists {
			comitment.AboutID = existingAbouts.ID
			config.DB.Create(&comitment)
		}
	}

	return c.JSON(fiber.Map{
		"message": "Update abouts success",
	})
}

func DeleteAllAbouts(c *fiber.Ctx) error {
	// Hapus semua ComitmentLists terlebih dahulu
	result := config.DB.Unscoped().Where("about_id > 0").Delete(&models.ComitmentList{})
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not delete comitment lists",
			"error":   result.Error.Error(),
		})
	}

	// Hapus semua Abouts
	result = config.DB.Unscoped().Where("id > 0").Delete(&models.About{})
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not delete abouts",
			"error":   result.Error.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Delete abouts success",
	})
}
