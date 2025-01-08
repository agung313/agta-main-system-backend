package routes

import (
	"github.com/agung313/agta-main-system-backend/controllers"
	"github.com/agung313/agta-main-system-backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupContactsRoutes(app *fiber.App) {
	api := app.Group("/admin")

	// get contacts
	api.Get("/contacts", middleware.JWTProtected(), controllers.GetContacts)

	// create or update contacts
	api.Put("/contacts", middleware.JWTProtectedAdmin(), controllers.CreateOrUpdateContacts)

	// delete contacts
	api.Delete("/contacts", middleware.JWTProtectedAdmin(), controllers.DeleteContacts)
}
