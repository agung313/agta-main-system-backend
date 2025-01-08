package controllers

import (
	"time"

	"github.com/agung313/agta-main-system-backend/config"
	"github.com/agung313/agta-main-system-backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// auth controller >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

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
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	if user.Role == "superAdmin" {
		listTokenAdmin := models.TokenAdmin{Token: tokenString}
		config.DB.Create(&listTokenAdmin)
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"token":   tokenString,
	})
}

func SignUp(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var existingUser models.User
	if err := config.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email already exists",
		})
	}

	if user.Role != "superAdmin" {
		user.Role = "visitor"
	}
	result := config.DB.Create(user)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not create user",
		})
	}

	if user.Role == "superAdmin" {
		listTokenAdmin := models.TokenAdmin{Token: tokenString}
		config.DB.Create(&listTokenAdmin)
	}

	return c.JSON(fiber.Map{
		"message": "Success create new user",
		"data":    user,
		"token":   tokenString,
	})
}

func Logout(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "No token provided",
		})
	}

	// Remove "Bearer " prefix from token
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	var existingToken models.Blacklist
	result := config.DB.Where("token = ?", token).First(&existingToken)
	if result.RowsAffected > 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Token already blacklisted",
		})
	}

	blacklistToken := models.Blacklist{Token: token}
	result = config.DB.Create(&blacklistToken)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not blacklist token",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Logout successful",
	})
}

func GetBlacklistTokens(c *fiber.Ctx) error {
	var blacklistTokens []models.Blacklist
	result := config.DB.Find(&blacklistTokens)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": result.Error,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success get all blacklist tokens",
		"total":   result.RowsAffected,
		"data":    blacklistTokens,
	})
}

func DeleteAllBlacklistTokens(c *fiber.Ctx) error {
	result := config.DB.Unscoped().Where("1 = 1").Delete(&models.Blacklist{})
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": result.Error,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success permanently deleted all blacklist tokens",
		"total":   result.RowsAffected,
	})
}

func GetListTokenAdmin(c *fiber.Ctx) error {
	var listTokenAdmin []models.TokenAdmin
	result := config.DB.Find(&listTokenAdmin)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": result.Error,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success get all list token admin",
		"total":   result.RowsAffected,
		"data":    listTokenAdmin,
	})
}

func DeleteAllListTokenAdmin(c *fiber.Ctx) error {
	result := config.DB.Unscoped().Where("1 = 1").Delete(&models.TokenAdmin{})
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": result.Error,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success permanently deleted all list token admin",
		"total":   result.RowsAffected,
	})
}

// user controller >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

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

	// Check if the email already exists
	var existingUser models.User
	if err := config.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email already exists",
		})
	}

	if user.Role != "superAdmin" {
		user.Role = "visitor"
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
