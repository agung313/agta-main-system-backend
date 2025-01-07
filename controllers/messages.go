package controllers

import (
	"github.com/agung313/agta-main-system-backend/config"
	"github.com/agung313/agta-main-system-backend/models"
	"github.com/gofiber/fiber/v2"
)

func GetMessages(c *fiber.Ctx) error {
	var messages []models.Message
	result := config.DB.Find(&messages)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not retrieve messages",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success get all messages",
		"total":   result.RowsAffected,
		"data":    messages,
	})
}

func CreateMessage(c *fiber.Ctx) error {
	message := new(models.Message)
	if err := c.BodyParser(message); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid input",
		})
	}

	result := config.DB.Create(message)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not send message",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Message sent",
	})

}

func DeleteMessageById(c *fiber.Ctx) error {
	id := c.Params("id")
	var message models.Message
	if err := config.DB.First(&message, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Message not found",
		})
	}

	if err := config.DB.Delete(&message).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not delete message",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Message deleted",
	})
}

func PermanentDeleteAllMessages(c *fiber.Ctx) error {
	result := config.DB.Unscoped().Where("1 = 1").Delete(&models.Message{})
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": result.Error,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success permanently deleted all messages",
		"total":   result.RowsAffected,
	})
}
