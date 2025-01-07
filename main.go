package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"github.com/agung313/agta-main-system-backend/config"
	"github.com/agung313/agta-main-system-backend/models"
	"github.com/agung313/agta-main-system-backend/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Database
	config.ConnectDatabase()
	config.DB.AutoMigrate(
		&models.User{},
		&models.Slogan{},
		&models.Blacklist{},
		&models.About{},
		&models.ComitmentList{},
		&models.Contacts{},
		&models.Service{},
		&models.TechnologyList{},
		&models.Message{},
		&models.Visitor{},
	)

	// Routes
	routes.SetupUserRoutes(app)
	routes.SetupSloganRoutes(app)
	routes.SetupAboutsRoutes(app)
	routes.SetupContactsRoutes(app)
	routes.SetupServicesRoutes(app)
	routes.SetupMessagesRoutes(app)
	routes.SetupVisitorRoutes(app)
	routes.UploadRoute(app)

	port := os.Getenv("APP_PORT")
	log.Fatal(app.Listen(":" + port))
}
