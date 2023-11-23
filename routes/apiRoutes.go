package routes

import (
	"tech_tubbies/controller"

	"github.com/gofiber/fiber/v2"
)

func AppRoutes(app *fiber.App) {
	// ROOT ENDPOINT
	app.Get("/", controller.ViewLogin)

	apiEndpoint := app.Group("/api")
	userEndpoint := apiEndpoint.Group("/user")
	userEndpoint.Get("/register", controller.ViewRegistration)
	userEndpoint.Post("/save", controller.RegisterUser)
	userEndpoint.Get("/login", controller.ViewLogin).Name("login_page")
	userEndpoint.Post("/verify", controller.VerifyUser)
	userEndpoint.Get("/profile", controller.ViewProfile)

	dashboardEndpoint := apiEndpoint.Group("/dashboard")
	dashboardEndpoint.Get("/", controller.ViewDashboard)
	dashboardEndpoint.Get("/secretariat", controller.ViewDashboardSecretariat)

	routingEndpoint := apiEndpoint.Group("/routing")
	routingEndpoint.Get("/", controller.ViewRouting)
	routingEndpoint.Get("/records", controller.ViewRoutingRecords)
	routingEndpoint.Get("/secretariat", controller.ViewRoutingSecretariat)
	routingEndpoint.Get("/foragenda/:id", controller.ViewRoutingForAgenda)
	routingEndpoint.Get("/receiving", controller.ViewReceivingRoute)
	routingEndpoint.Get("/download/:filename", controller.DownloadAttachment)
	routingEndpoint.Post("/register_receiving", controller.RegisterReceiving)
	routingEndpoint.Get("/for_filing/:id", controller.GetForFiling)
	routingEndpoint.Post("/update_filing", controller.UpdateForFiling)
	routingEndpoint.Get("/routingApproved", controller.ViewApproved)
	routingEndpoint.Get("/add_committee/:itemNo/:committeeId/:userId", controller.InsertCommitteeForAgenda)

	trackingEndpoint := apiEndpoint.Group("/tracking")
	trackingEndpoint.Get("/", controller.ViewTracking)

	testEndpoint := apiEndpoint.Group("/test")
	testEndpoint.Get("/set_cookie", controller.SetCookie)
	testEndpoint.Get("/get_cookie", controller.GetCookie)
}
