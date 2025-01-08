package routes

import (
	"github.com/agung313/agta-main-system-backend/controllers"
	"github.com/agung313/agta-main-system-backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupMessagesRoutes(app *fiber.App) {
	api := app.Group("/admin")
	api.Get("/messages", middleware.JWTProtectedAdmin(), controllers.GetMessages)
	api.Delete("/message/:id", middleware.JWTProtectedAdmin(), controllers.DeleteMessageById)
	api.Delete("/messages", middleware.JWTProtectedAdmin(), controllers.PermanentDeleteAllMessages)
}
