package controllers

import (
	"github.com/agung313/agta-main-system-backend/config"
	"github.com/agung313/agta-main-system-backend/models"
	"github.com/gofiber/fiber/v2"
)

func GetDashboard(c *fiber.Ctx) error {
	var slogan []models.Slogan
	result := config.DB.Find(&slogan)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not find slogan",
		})
	}

	var technologyLists []models.TechnologyList
	resultTechnologyLists := config.DB.Find(&technologyLists)
	if resultTechnologyLists.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not find technology lists",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Get slogan success",
		"data": map[string]interface{}{
			"firstText":       slogan[0].FirstText,
			"secondText":      slogan[0].SecondText,
			"thirdText":       slogan[0].ThirdText,
			"description":     slogan[0].Description,
			"technologyLists": technologyLists,
		},
	})
}
