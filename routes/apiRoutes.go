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

	routingEndpoint.Get("/download/:filename", controller.DownloadAttachment)
	// Receiving
	routingEndpoint.Get("/receiving", controller.ViewReceivingRoute)
	routingEndpoint.Post("/register_receiving", controller.RegisterReceiving)
	// Agenda
	routingEndpoint.Get("/secretariat", controller.ViewRoutingSecretariat)
	routingEndpoint.Get("/foragenda/:docId", controller.ViewRoutingForAgenda)
	routingEndpoint.Post("/register_agenda", controller.RegisterForAgenda)
	routingEndpoint.Post("/update_agenda", controller.UpdateForAgenda)
	routingEndpoint.Get("/viewagenda/:docId/:itemNumber", controller.ViewForAgenda)
	// Approved
	routingEndpoint.Get("/approved/:docId/:itemNumber", controller.ViewApproved)
	routingEndpoint.Post("/register_approved", controller.RegisterApproved)

	// Releasing
	routingEndpoint.Get("/releasing/:docId/:itemNumber", controller.ViewReleasing)
	routingEndpoint.Post("/register_releasing", controller.RegisterReleasing)
	routingEndpoint.Get("/update_releasing/:docId/:itemNumber", controller.UpdateReleasing)
	routingEndpoint.Post("/save_releasing", controller.SaveReleasing)

	// Filing
	routingEndpoint.Get("/filing/:docId/:itemNumber", controller.ViewForFiling)
	routingEndpoint.Post("/register_filing", controller.RegisterFiling)
	routingEndpoint.Get("/update_filing/:docId/:itemNumber", controller.UpdateFiling)
	routingEndpoint.Post("/save_filing", controller.SaveFiling)

	// Tracking
	trackingEndpoint := apiEndpoint.Group("/tracking")
	trackingEndpoint.Get("/", controller.ViewTracking)
	trackingEndpoint.Post("/search", controller.SearchTracking)

	// OTHER
	routingEndpoint.Get("/add_committee/:itemNo/:committeeId/:userId", controller.InsertCommitteeForAgenda)
	routingEndpoint.Post("/add_committee", controller.PostInsertCommitteeForAgenda)

	testEndpoint := apiEndpoint.Group("/test")
	testEndpoint.Get("/set_cookie", controller.SetCookie)
	testEndpoint.Get("/get_cookie", controller.GetCookie)
}
