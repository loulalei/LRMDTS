package main

import (
	"tech_tubbies/controller"
	"tech_tubbies/middleware/database"
	"tech_tubbies/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Connect to DB
	database.ConnectPostgres()
	controller.InitializeTables()
	app := fiber.New(fiber.Config{
		Views: html.New("./template", ".html"),
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}),
		logger.New(logger.Config{}),
	)

	// Register the session manager middleware
	// app.Use(model.SessionManager)

	// Dito iloload lahat ng path dun sa assets
	app.Static("/", "./assets/css")
	app.Static("/", "./assets/js")
	app.Static("/", "./assets/images")
	app.Static("/", "./assets/files")

	// Configure application CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	routes.AppRoutes(app)
	app.Listen(":5000")
}
