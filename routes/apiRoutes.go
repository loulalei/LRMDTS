package routes

import (
	"tech_tubbies/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func AppRoutes(app *fiber.App, sess *session.Store) {
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
	routingEndpoint.Get("/foragenda/:docId", controller.ViewRoutingForAgenda)

	routingEndpoint.Get("/download/:filename", controller.DownloadAttachment)
	// Receiving
	routingEndpoint.Get("/receiving", controller.ViewReceivingRoute)
	routingEndpoint.Post("/register_receiving", controller.RegisterReceiving)
	// Agenda
	routingEndpoint.Get("/foragenda/:docId", controller.ViewRoutingForAgenda)
	routingEndpoint.Post("/register_agenda", controller.RegisterForAgenda)
	routingEndpoint.Get("/viewagenda/:docId/:itemNumber", controller.ViewForAgenda)
	routingEndpoint.Post("/update_agenda", controller.RegisterForAgenda)

	routingEndpoint.Get("/for_filing/:id", controller.GetForFiling)
	routingEndpoint.Get("/routingApproved", controller.ViewApproved)
	routingEndpoint.Get("/routingReleasing", controller.ViewReleasing)
	routingEndpoint.Get("/add_committee/:itemNo/:committeeId/:userId", controller.InsertCommitteeForAgenda)
	routingEndpoint.Post("/add_committee", controller.PostInsertCommitteeForAgenda)

	trackingEndpoint := apiEndpoint.Group("/tracking")
	trackingEndpoint.Get("/", controller.ViewTracking)

	testEndpoint := apiEndpoint.Group("/test")
	testEndpoint.Get("/set_cookie", controller.SetCookie)
	testEndpoint.Get("/get_cookie", controller.GetCookie)
}
