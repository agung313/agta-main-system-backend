package middleware

import (
	"log"

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

		c.Locals("user", token.Claims)
		return c.Next()
	}
}
