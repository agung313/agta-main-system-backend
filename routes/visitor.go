package routes

import (
	"github.com/agung313/agta-main-system-backend/controllers"
	"github.com/agung313/agta-main-system-backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupVisitorRoutes(app *fiber.App) {
	api := app.Group("/admin")
	api.Get("/visitors", middleware.JWTProtectedAdmin(), controllers.GetVisitors)
	api.Delete("/visitors", middleware.JWTProtectedAdmin(), controllers.DeleteAllVisitors)
}
