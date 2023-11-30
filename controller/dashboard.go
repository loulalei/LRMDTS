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
	receiving := &[]model.Routings{}

	database.DBConn.Raw("SELECT * FROM routings").Scan(receiving)

	return c.Render("dashboard", fiber.Map{
		"pageTitle":   "Dashboard",
		"title":       "DASHBOARD",
		"yearNow":     model.YearNow,
		"user":        model.Fullname,
		"userLogged":  model.UserCodeLogged,
		"greetings":   utils.GetGreetings(),
		"notifs":      receiving,
		"filings":     CountForFiling(),
		"loginStatus": 100,
	})
}

func ViewDashboardSecretariat(c *fiber.Ctx) error {
	if model.Fullname == "" {
		return c.Redirect("/")
	}
	receiving := &[]model.Routings{}

	database.DBConn.Raw("SELECT * FROM routings WHERE document_tag = 'For Agenda'").Scan(receiving)

	return c.Render("dashboard_secretarial", fiber.Map{
		"pageTitle":   "Dashboard - Secretariat",
		"title":       "DASHBOARD - SECRETARIAT",
		"yearNow":     model.YearNow,
		"user":        model.Fullname,
		"userLogged":  model.UserCodeLogged,
		"greetings":   utils.GetGreetings(),
		"notifs":      receiving,
		"agendas":     CountForAgenda(),
		"loginStatus": 100,
	})
}

func CountForAgenda() int64 {
	total := database.DBConn.Exec("SELECT * FROM routings WHERE document_tag = 'For Agenda'").RowsAffected
	return total
}

func CountForFiling() int64 {
	total := database.DBConn.Exec("SELECT * FROM routings WHERE document_tag = 'For Filing'").RowsAffected
	return total
}
