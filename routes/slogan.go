package routes

import (
	"github.com/agung313/agta-main-system-backend/controllers"
	"github.com/agung313/agta-main-system-backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupSloganRoutes(app *fiber.App) {
	api := app.Group("/api")

	// get slogan
	api.Get("/slogan", middleware.JWTProtected(), controllers.GetSlogan)

	// create slogan
	api.Put("/slogan", middleware.JWTProtected(), controllers.CreateOrUpdateSlogan)
}
