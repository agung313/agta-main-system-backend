package routes

import (
	"github.com/agung313/agta-main-system-backend/controllers"
	"github.com/agung313/agta-main-system-backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupServicesRoutes(app *fiber.App) {
	api := app.Group("/api")

	// get services
	api.Get("/services", controllers.GetServices)

	// create or update services
	api.Put("/services", middleware.JWTProtected(), controllers.CreateOrUpdateServices)

	// delete services
	api.Delete("/services", middleware.JWTProtected(), controllers.DeleteAllServies)
}
