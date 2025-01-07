package routes

import (
	"github.com/agung313/agta-main-system-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupClientRoutes(app *fiber.App) {
	api := app.Group("/client")

	api.Post("/signup", controllers.SignUp)
	api.Post("/login", controllers.Login)
	api.Get("/dashboard", controllers.GetDashboard)
	api.Get("/abouts", controllers.GetAbouts)
	api.Get("/services", controllers.GetServices)
	api.Get("/contacts", controllers.GetContacts)
	api.Post("/messages", controllers.CreateMessage)
	api.Post("/visitor", controllers.CreateVisitor)
}
