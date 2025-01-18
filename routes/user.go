package routes

import (
	"github.com/agung313/agta-main-system-backend/controllers"
	"github.com/agung313/agta-main-system-backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	api := app.Group("/admin")

	// aut routes >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

	// logout
	api.Post("/logout", controllers.Logout)

	// get blacklist token
	api.Get("/blacklist", middleware.JWTProtectedAdmin(), controllers.GetBlacklistTokens)

	// DELTE BLACKLIST TOKEN
	api.Delete("/blacklist", middleware.JWTProtectedAdmin(), controllers.DeleteAllBlacklistTokens)

	// get list token admin
	api.Get("/listTokenAdmin", middleware.JWTProtectedAdmin(), controllers.GetListTokenAdmin)

	// delete all list token admin
	api.Delete("/listTokenAdmin", middleware.JWTProtectedAdmin(), controllers.DeleteAllListTokenAdmin)

	// user routes >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

	// get all users
	api.Get("/users", middleware.JWTProtectedAdmin(), controllers.GetUsers)

	// get all deleted users
	api.Get("/users/deleted", middleware.JWTProtectedAdmin(), controllers.GetDeletedUsers)

	// restore user by id
	api.Get("/user/:id/restore", middleware.JWTProtectedAdmin(), controllers.RestoreUser)

	// restore all users
	api.Get("/users/restore", middleware.JWTProtectedAdmin(), controllers.RestoreAllUsers)

	// create user
	api.Post("/user", middleware.JWTProtectedAdmin(), controllers.CreateUser)

	// delete user by id
	api.Delete("/user/:id", middleware.JWTProtectedAdmin(), controllers.DeleteUser)

	// delete all users
	api.Delete("/users", middleware.JWTProtectedAdmin(), controllers.DeleteAllUsers)

	// permanent delete user by id
	api.Delete("/user/:id/permanent", middleware.JWTProtectedAdmin(), controllers.PermanentDeleteUser)

	// permanent delete all users
	api.Delete("/users/permanent", middleware.JWTProtectedAdmin(), controllers.PermanentDeleteAllUsers)

	// get user by id
	api.Get("/user/:id", middleware.JWTProtected(), controllers.GetUser)

	// update user by id
	api.Put("/user/:email", middleware.JWTProtected(), controllers.UpdateUser)
}
