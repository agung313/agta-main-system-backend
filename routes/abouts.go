package routes

import (
	"github.com/agung313/agta-main-system-backend/controllers"
	"github.com/agung313/agta-main-system-backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupAboutsRoutes(app *fiber.App) {
	api := app.Group("/admin")

	// get abouts
	api.Get("/abouts", middleware.JWTProtected(), controllers.GetAbouts)

	// create or update abouts
	api.Put("/abouts", middleware.JWTProtected(), controllers.CreateOrUpdateAbouts)

	// delete abouts
	api.Delete("/abouts", middleware.JWTProtected(), controllers.DeleteAllAbouts)
}
