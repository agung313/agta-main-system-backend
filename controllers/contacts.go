package controllers

import (
	"github.com/agung313/agta-main-system-backend/config"
	"github.com/agung313/agta-main-system-backend/models"
	"github.com/gofiber/fiber/v2"
)

func GetContacts(c *fiber.Ctx) error {
	var contacts []models.Contacts
	result := config.DB.Find(&contacts)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not find contacts",
		})
	}
	if len(contacts) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "No contacts found",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Get contacts success",
		"data": map[string]interface{}{
			"email":          contacts[0].Email,
			"instagram":      contacts[0].Instagram,
			"linkedin":       contacts[0].Linkedin,
			"address":        contacts[0].Address,
			"googleMapsLink": contacts[0].GoogleMapsLink,
		},
	})
}

func CreateOrUpdateContacts(c *fiber.Ctx) error {
	contacts := new(models.Contacts)
	if err := c.BodyParser(contacts); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}
	var existingContacts models.Contacts
	result := config.DB.First(&existingContacts)
	if result.Error != nil {
		// No existing contacts found, create a new one
		result = config.DB.Create(&contacts)
		if result.Error != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Could not create contacts",
				"error":   result.Error.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"message": "Create contacts success",
		})
	}
	// Existing contacts found, update it
	result = config.DB.Model(&existingContacts).Updates(contacts)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not update contacts",
			"error":   result.Error.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Update contacts success",
	})
}

func DeleteContacts(c *fiber.Ctx) error {
	var contacts models.Contacts
	result := config.DB.First(&contacts)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "No contacts found",
		})
	}
	result = config.DB.Unscoped().Delete(&contacts)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not delete contacts",
			"error":   result.Error.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Delete contacts success",
	})
}
