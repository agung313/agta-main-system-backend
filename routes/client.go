package routes

import (
	"github.com/agung313/agta-main-system-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupClientRoutes(app *fiber.App) {
	app.Post("/signup", controllers.SignUp)
	app.Post("/login", controllers.Login)
	app.Get("/dashboard", controllers.GetDashboard)
	app.Get("/abouts", controllers.GetAbouts)
	app.Get("/services", controllers.GetServices)
	app.Get("/contacts", controllers.GetContacts)
	app.Post("/messages", controllers.CreateMessage)
	app.Post("/visitor", controllers.CreateVisitor)
	app.Post("/resetPassword", controllers.ResetPassword)
}
