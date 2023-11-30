package main

import (
	"fmt"
	"log"
	"tech_tubbies/controller"
	"tech_tubbies/middleware/database"
	"tech_tubbies/middleware/envRouting"
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

	if envRouting.SSL == "enabled" {
		log.Fatal(app.ListenTLS(
			fmt.Sprintf(":%s", envRouting.Port),
			envRouting.SSLCertificate,
			envRouting.SSLKey,
		))
	} else {
		log.Fatal(app.Listen(fmt.Sprintf(":%s", envRouting.Port)))
	}
}
