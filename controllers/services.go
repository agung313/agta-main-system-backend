package controllers

import (
	"github.com/agung313/agta-main-system-backend/config"
	"github.com/agung313/agta-main-system-backend/models"
	"github.com/gofiber/fiber/v2"
)

func GetServices(c *fiber.Ctx) error {
	var services []models.Service
	result := config.DB.Preload("TechnologyList").Find(&services)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not find services",
		})
	}
	if len(services) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "No services found",
		})
	}
	technologyLists := make([]map[string]interface{}, len(services[0].TechnologyList))
	for i, technology := range services[0].TechnologyList {
		technologyLists[i] = map[string]interface{}{
			"icont":           technology.Icont,
			"title":           technology.Title,
			"link":            technology.Link,
			"descriptionText": technology.Description,
		}
	}
	return c.JSON(fiber.Map{
		"message": "Get services success",
		"data": map[string]interface{}{
			"title": map[string]string{
				"id": "LAYANAN",
				"en": "SERVICES",
			},
			"description":     services[0].Description,
			"technologyLists": technologyLists,
		},
	})
}

func CreateOrUpdateServices(c *fiber.Ctx) error {
	services := new(models.Service)
	if err := c.BodyParser(services); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}

	var existingServices models.Service
	result := config.DB.Preload("TechnologyList").First(&existingServices)

	if result.Error != nil {
		// No existing services found, create a new one
		result = config.DB.Create(&services)
		if result.Error != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Could not create services",
				"error":   result.Error.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"message": "Create services success",
		})
	}

	// Existing services found, update it
	result = config.DB.Model(&existingServices).Updates(services)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not update services",
			"error":   result.Error.Error(),
		})
	}

	// Update TechnologyList
	if len(services.TechnologyList) > 0 {
		// Delete existing TechnologyList
		config.DB.Where("service_id = ?", existingServices.ID).Delete(&models.TechnologyList{})
		// Add new TechnologyList
		for _, comitment := range services.TechnologyList {
			comitment.ServiceId = existingServices.ID
			config.DB.Create(&comitment)
		}
	}

	return c.JSON(fiber.Map{
		"message": "Update services success",
	})
}

func DeleteAllServices(c *fiber.Ctx) error {
	// Delete all TechnologyLists first
	result := config.DB.Where("service_id > 0").Unscoped().Delete(&models.TechnologyList{})
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not delete technology lists",
			"error":   result.Error.Error(),
		})
	}

	// Delete all services
	result = config.DB.Where("id > 0").Unscoped().Delete(&models.Service{})
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not delete services",
			"error":   result.Error.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Delete services success",
	})
}
