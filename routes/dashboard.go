package routes

import (
	"github.com/agung313/agta-main-system-backend/controllers"
	"github.com/agung313/agta-main-system-backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupDashboardAdminRoutes(app *fiber.App) {
	api := app.Group("/api/")
	api.Get("/dashboard", middleware.JWTProtected(), controllers.GetDashboardAdmin)
}
