package controllers

import (
	"github.com/agung313/agta-main-system-backend/config"
	"github.com/agung313/agta-main-system-backend/models"
	"github.com/gofiber/fiber/v2"
)

func GetDashboardAdmin(c *fiber.Ctx) error {
	var visitors []models.Visitor
	var messages []models.Message

	result := config.DB.Find(&visitors)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not retrieve visitors",
		})
	}

	resultMessages := config.DB.Find(&messages)
	if resultMessages.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not retrieve messages",
		})
	}

	totalVisitors := result.RowsAffected
	totalMessages := resultMessages.RowsAffected

	// Calculate totalCountries
	var Countries []string
	config.DB.Model(&models.Visitor{}).Distinct("countries").Pluck("countries", &Countries)
	totalCountries := len(Countries)

	// Filter data by year and month
	data := make(map[int][]map[string]interface{})
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	for _, visitor := range visitors {
		year := visitor.CreatedAt.Year()
		month := visitor.CreatedAt.Month()
		if _, ok := data[year]; !ok {
			data[year] = make([]map[string]interface{}, 12)
			for i := 0; i < 12; i++ {
				data[year][i] = map[string]interface{}{
					"name":      months[i],
					"Visits":    0,
					"Messages":  0,
					"Countries": 0,
				}
			}
		}
		data[year][month-1]["Visits"] = data[year][month-1]["Visits"].(int) + 1
	}

	for _, message := range messages {
		year := message.CreatedAt.Year()
		month := message.CreatedAt.Month()
		if _, ok := data[year]; ok {
			data[year][month-1]["Messages"] = data[year][month-1]["Messages"].(int) + 1
		}
	}

	uniqueCountries := make(map[int]map[int]map[string]bool)
	for _, visitor := range visitors {
		year := visitor.CreatedAt.Year()
		month := int(visitor.CreatedAt.Month())
		country := visitor.Countries
		if _, ok := uniqueCountries[year]; !ok {
			uniqueCountries[year] = make(map[int]map[string]bool)
		}
		if _, ok := uniqueCountries[year][month]; !ok {
			uniqueCountries[year][month] = make(map[string]bool)
		}
		if _, ok := uniqueCountries[year][month][country]; !ok {
			uniqueCountries[year][month][country] = true
			if _, ok := data[year]; ok {
				data[year][month-1]["Countries"] = data[year][month-1]["Countries"].(int) + 1
			}
		}
	}

	return c.JSON(fiber.Map{
		"message":        "Success get all visitors",
		"totalVisitors":  totalVisitors,
		"totalMessages":  totalMessages,
		"totalCountries": totalCountries,
		"data":           data,
	})
}
