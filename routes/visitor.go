package routes

import (
	"github.com/agung313/agta-main-system-backend/controllers"
	"github.com/agung313/agta-main-system-backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupVisitorRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Delete("/visitors", middleware.JWTProtected(), controllers.DeleteAllVisitors)
	api.Post("/visitor", controllers.CreateVisitor)
}
