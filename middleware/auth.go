package middleware

import (
	"log"

	"github.com/agung313/agta-main-system-backend/config"
	"github.com/agung313/agta-main-system-backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"message": "Missing or malformed JWT",
			})
		}

		tokenString := authHeader[len("Bearer "):]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
			}
			return []byte("secret"), nil
		})

		if err != nil || !token.Valid {
			log.Println("Token Error:", err)
			return c.Status(401).JSON(fiber.Map{
				"message": "Invalid or expired JWT",
			})
		}

		// Check if token is blacklisted
		var tokens []models.Blacklist
		config.DB.Find(&tokens)

		for _, t := range tokens {
			if t.Token == tokenString {
				return c.Status(401).JSON(fiber.Map{
					"message": "Token is blacklisted",
				})
			}
		}

		c.Locals("user", token.Claims)
		return c.Next()
	}
}

func JWTProtectedAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"message": "Missing or malformed JWT",
			})
		}

		tokenString := authHeader[len("Bearer "):]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
			}
			return []byte("secret"), nil
		})

		if err != nil || !token.Valid {
			log.Println("Token Error:", err)
			return c.Status(401).JSON(fiber.Map{
				"message": "Invalid or expired JWT",
			})
		}

		// Check if token is blacklisted
		var tokens []models.Blacklist
		config.DB.Find(&tokens)

		for _, t := range tokens {
			if t.Token == tokenString {
				return c.Status(401).JSON(fiber.Map{
					"message": "Token is blacklisted",
				})
			}
		}

		// check if user is admin
		var tokenAdmin []models.TokenAdmin
		config.DB.Find(&tokenAdmin)

		for _, t := range tokenAdmin {
			if t.Token != tokenString {
				return c.Status(401).JSON(fiber.Map{
					"message": "Unauthorized",
				})
			}
		}

		c.Locals("user", token.Claims)
		return c.Next()
	}
}
