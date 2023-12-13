package controller

import (
	"tech_tubbies/middleware/database"
	utils "tech_tubbies/middleware/util"
	"tech_tubbies/model"

	"github.com/gofiber/fiber/v2"
)

func ViewDashboard(c *fiber.Ctx) error {
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	routings := &[]model.ViewRoutings{}
	database.DBConn.Raw("SELECT * FROM view_routings").Scan(routings)

	return c.Render("dashboard", fiber.Map{
		"pageTitle":   "Dashboard",
		"title":       "DASHBOARD",
		"yearNow":     model.YearNow,
		"user":        model.Fullname,
		"userLogged":  model.UserCodeLogged,
		"greetings":   utils.GetGreetings(),
		"notifs":      routings,
		"filings":     CountForFiling(),
		"releases":    CountForReleasing(),
		"ordinances":  CountOrdinances(),
		"resolutions": CountResolutions(),
		"loginStatus": c.Response().StatusCode(),
		"baseURL":     c.BaseURL(),
	})
}

func ViewDashboardSecretariat(c *fiber.Ctx) error {
	if model.Fullname == "" {
		return c.Redirect("/")
	}
	routings := &[]model.ViewRoutings{}

	database.DBConn.Raw("SELECT * FROM view_routings WHERE document_tag = 'For Agenda' OR document_tag = 'Referred to Committee'").Scan(routings)

	return c.Render("dashboard_secretarial", fiber.Map{
		"pageTitle":  "Dashboard - Secretariat",
		"title":      "DASHBOARD - SECRETARIAT",
		"yearNow":    model.YearNow,
		"user":       model.Fullname,
		"userLogged": model.UserCodeLogged,
		"greetings":  utils.GetGreetings(),
		"notifs":     routings,
		"agendas":    CountForAgenda(),
		"baseURL":    c.BaseURL(),
	})
}

func CountForAgenda() int64 {
	total := database.DBConn.Exec("SELECT * FROM routings WHERE document_tag = 'For Agenda' ").RowsAffected
	return total
}

func CountForFiling() int64 {
	total := database.DBConn.Exec("SELECT * FROM routings WHERE document_tag = 'For Filing'").RowsAffected
	return total
}

func CountForReleasing() int64 {
	total := database.DBConn.Exec("SELECT * FROM routings WHERE document_tag = 'For Releasing'").RowsAffected
	return total
}

func CountOrdinances() int64 {
	total := database.DBConn.Exec("SELECT * FROM trackings WHERE law_type = 'Ordinance'").RowsAffected
	return total
}

func CountResolutions() int64 {
	total := database.DBConn.Exec("SELECT * FROM trackings WHERE law_type = 'Resolution'").RowsAffected
	return total
}
