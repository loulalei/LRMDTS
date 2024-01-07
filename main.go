package main

import (
	"fmt"
	"log"
	"tech_tubbies/controller"
	"tech_tubbies/global"
	"tech_tubbies/middleware/database"
	"tech_tubbies/middleware/envRouting"

	"tech_tubbies/routes"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	// "github.com/gofiber/session/v2"
)

func main() {

	// load env
	envRouting.LoadEnv()
	// Connect to DB
	database.ConnectPostgres()
	controller.InitializeTables()

	global.App.Use(recover.New())

	global.App.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}),
		logger.New(logger.Config{}),
	)

	// Dito iloload lahat ng path dun sa assets
	global.App.Static("/", "./assets/css")
	global.App.Static("/", "./assets/js")
	global.App.Static("/", "./assets/images")
	global.App.Static("/", "./assets/files")

	// Initialize default config
	global.App.Use(favicon.New(favicon.Config{
		File: "./assets/images/favicon.ico",
	}))

	// Configure application CORS
	global.App.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	routes.AppRoutes(global.App)

	fmt.Println("SSL:", envRouting.SSL)
	if envRouting.SSL == "disable" {
		fmt.Println("CERT: none")
		fmt.Println("KEY none")
	} else {
		fmt.Println("CERT:", envRouting.SSLCertificate)
		fmt.Println("KEY:", envRouting.SSLKey)

	}

	if envRouting.SSL == "enabled" {
		log.Fatal(global.App.ListenTLS(
			fmt.Sprintf(":%s", envRouting.Port),
			envRouting.SSLCertificate,
			envRouting.SSLKey,
		))
	} else {
		log.Fatal(global.App.Listen(fmt.Sprintf(":%s", envRouting.Port)))
	}
}
