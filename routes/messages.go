package routes

import (
	"github.com/agung313/agta-main-system-backend/controllers"
	"github.com/agung313/agta-main-system-backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupMessagesRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/messages", middleware.JWTProtected(), controllers.GetMessages)
	api.Post("/messages", controllers.CreateMessage)
	api.Delete("/message/:id", middleware.JWTProtected(), controllers.DeleteMessageById)
	api.Delete("/messages", middleware.JWTProtected(), controllers.PermanentDeleteAllMessages)
}
