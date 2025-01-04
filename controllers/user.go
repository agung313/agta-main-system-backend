package controllers

import (
	"time"

	"github.com/agung313/agta-main-system-backend/config"
	"github.com/agung313/agta-main-system-backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid input",
		})
	}

	var user models.User
	result := config.DB.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"token":   tokenString,
	})
}

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
