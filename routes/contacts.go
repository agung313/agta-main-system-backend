package routes

import (
	"github.com/agung313/agta-main-system-backend/controllers"
	"github.com/agung313/agta-main-system-backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupContactsRoutes(app *fiber.App) {
	api := app.Group("/api")

	// get contacts
	api.Get("/contacts", controllers.GetContacts)

	// create or update contacts
	api.Put("/contacts", middleware.JWTProtected(), controllers.CreateOrUpdateContacts)

	// delete contacts
	api.Delete("/contacts", middleware.JWTProtected(), controllers.DeleteContacts)
}
