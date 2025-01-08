// filepath: routes/upload.go
package routes

import (
	"github.com/agung313/agta-main-system-backend/controllers"
	"github.com/agung313/agta-main-system-backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func UploadRoute(app *fiber.App) {
	api := app.Group("/admin")
	api.Post("/uploadImage", middleware.JWTProtectedAdmin(), controllers.UploadImage)
	api.Delete("/deleteImage", middleware.JWTProtectedAdmin(), controllers.DeleteImage)
	app.Static("/uploads", "./uploads")

}
