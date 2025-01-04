package routes

import (
	"github.com/agung313/agta-main-system-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	api := app.Group("/api")

	// get all users
	api.Get("/users", controllers.GetUsers)

	// get user by id
	api.Get("/user/:id", controllers.GetUser)

	// get all deleted users
	api.Get("/users/deleted", controllers.GetDeletedUsers)

	// restore user by id
	api.Get("/user/:id/restore", controllers.RestoreUser)

	// restore all users
	api.Get("/users/restore", controllers.RestoreAllUsers)

	// create user
	api.Post("/user", controllers.CreateUser)

	// update user by id
	api.Put("/user/:id", controllers.UpdateUser)

	// delete user by id
	api.Delete("/user/:id", controllers.DeleteUser)

	// delete all users
	api.Delete("/users", controllers.DeleteAllUsers)

	// permanent delete user by id
	api.Delete("/user/:id/permanent", controllers.PermanentDeleteUser)

	// permanent delete all users
	api.Delete("/users/permanent", controllers.PermanentDeleteAllUsers)
}
