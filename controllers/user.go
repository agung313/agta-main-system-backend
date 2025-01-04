package controllers

import (
	"github.com/agung313/agta-main-system-backend/config"
	"github.com/agung313/agta-main-system-backend/models"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	result := config.DB.Find(&users)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": result.Error,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success get all users",
		"total":   result.RowsAffected,
		"data":    users,
	})
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	result := config.DB.First(&user, "id = ?", id)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": result.Error,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success get user",
		"data":    user,
	})
}

func GetDeletedUsers(c *fiber.Ctx) error {
	var users []models.User
	result := config.DB.Unscoped().Where("deleted_at IS NOT NULL").Find(&users)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": result.Error,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success get all deleted users",
		"total":   result.RowsAffected,
		"data":    users,
	})
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if user.Role != "superAdmin" {
		user.Role = "visiting"
	}
	result := config.DB.Create(user)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success create new user",
		"data":    user,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err,
		})
	}
	result := config.DB.Model(&user).Where("id = ?", id).Updates(&user)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": result.Error,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success update user's data",
		"data":    user,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	result := config.DB.Delete(&user, "id = ?", id)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": result.Error,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success deleted user",
	})
}

func DeleteAllUsers(c *fiber.Ctx) error {
	result := config.DB.Where("1 = 1").Delete(&models.User{})
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": result.Error,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success deleted all users",
		"total":   result.RowsAffected,
	})
}

func RestoreUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	result := config.DB.Unscoped().Where("id = ?", id).First(&user)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": result.Error,
		})
	}
	result = config.DB.Unscoped().Model(&user).Update("deleted_at", nil)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": result.Error,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success restored user",
		"data":    user,
	})
}

func RestoreAllUsers(c *fiber.Ctx) error {
	result := config.DB.Unscoped().Model(&models.User{}).Where("deleted_at IS NOT NULL").Update("deleted_at", nil)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": result.Error,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success restored all users",
		"total":   result.RowsAffected,
	})
}

func PermanentDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	result := config.DB.Unscoped().Delete(&user, "id = ?", id)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": result.Error,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success permanently deleted user",
	})
}

func PermanentDeleteAllUsers(c *fiber.Ctx) error {
	result := config.DB.Unscoped().Where("1 = 1").Delete(&models.User{})
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": result.Error,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success permanently deleted all users",
		"total":   result.RowsAffected,
	})
}
