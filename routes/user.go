package routes

import (
	"github.com/agung313/agta-main-system-backend/controllers"
	"github.com/agung313/agta-main-system-backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	api := app.Group("/api")

	// login
	api.Post("/login", controllers.Login)

	// get all users
	api.Get("/users", middleware.JWTProtected(), controllers.GetUsers)

	// get user by id
	api.Get("/user/:id", middleware.JWTProtected(), controllers.GetUser)

	// get all deleted users
	api.Get("/users/deleted", middleware.JWTProtected(), controllers.GetDeletedUsers)

	// restore user by id
	api.Get("/user/:id/restore", middleware.JWTProtected(), controllers.RestoreUser)

	// restore all users
	api.Get("/users/restore", middleware.JWTProtected(), controllers.RestoreAllUsers)

	// create user
	api.Post("/user", middleware.JWTProtected(), controllers.CreateUser)

	// update user by id
	api.Put("/user/:id", middleware.JWTProtected(), controllers.UpdateUser)

	// delete user by id
	api.Delete("/user/:id", middleware.JWTProtected(), controllers.DeleteUser)

	// delete all users
	api.Delete("/users", middleware.JWTProtected(), controllers.DeleteAllUsers)

	// permanent delete user by id
	api.Delete("/user/:id/permanent", middleware.JWTProtected(), controllers.PermanentDeleteUser)

	// permanent delete all users
	api.Delete("/users/permanent", middleware.JWTProtected(), controllers.PermanentDeleteAllUsers)
}
