package routes

import (
	"github.com/agung313/agta-main-system-backend/controllers"
	"github.com/agung313/agta-main-system-backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupMessagesRoutes(app *fiber.App) {
	api := app.Group("/admin")
	api.Get("/messages", middleware.JWTProtected(), controllers.GetMessages)
	api.Delete("/message/:id", middleware.JWTProtected(), controllers.DeleteMessageById)
	api.Delete("/messages", middleware.JWTProtected(), controllers.PermanentDeleteAllMessages)
}
