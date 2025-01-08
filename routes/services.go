package routes

import (
	"github.com/agung313/agta-main-system-backend/controllers"
	"github.com/agung313/agta-main-system-backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupServicesRoutes(app *fiber.App) {
	api := app.Group("/admin")

	// get services
	api.Get("/services", middleware.JWTProtected(), controllers.GetServices)

	// create or update services
	api.Put("/services", middleware.JWTProtectedAdmin(), controllers.CreateOrUpdateServices)

	// delete services
	api.Delete("/services", middleware.JWTProtectedAdmin(), controllers.DeleteAllServices)
}
