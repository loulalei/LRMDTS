package controller

import (
	"tech_tubbies/global"
	"tech_tubbies/middleware/database"
	utils "tech_tubbies/middleware/util"
	"tech_tubbies/model"

	"github.com/gofiber/fiber/v2"
)

func ViewDashboard(c *fiber.Ctx) error {
	// Get session from storage
	session, err := global.Store.Get(c)
	if err != nil {
		panic(err)
	}
	userId, _ := session.Get("userId").(string)

	if userId == "" {
		return c.Redirect("/")
	} else {
		userCredentials := &model.UserCredentials{}
		database.DBConn.Debug().Raw("SELECT * FROM user_credentials WHeRE id = ?", userId).Scan(userCredentials)
		global.Fullname = userCredentials.Fullname
		global.UserCodeLogged = userCredentials.DivisionCode
		global.UserID = userCredentials.Id
		global.DivisionCode = userCredentials.DivisionCode
	}

	routings := &[]model.ViewRoutings{}
	database.DBConn.Debug().Raw("SELECT * FROM view_routings").Scan(routings)

	return c.Render("dashboard", fiber.Map{
		"pageTitle":   "Dashboard",
		"title":       "DASHBOARD",
		"yearNow":     global.YearNow,
		"user":        global.Fullname,
		"userLogged":  global.UserCodeLogged,
		"userId":      global.UserID,
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
	// Get session from storage
	session, err := global.Store.Get(c)
	if err != nil {
		panic(err)
	}
	userId, _ := session.Get("userId").(string)

	if userId == "" {
		return c.Redirect("/")
	} else {
		userCredentials := &model.UserCredentials{}
		database.DBConn.Debug().Raw("SELECT * FROM user_credentials WHeRE id = ?", userId).Scan(userCredentials)
		global.Fullname = userCredentials.Fullname
		global.UserCodeLogged = userCredentials.DivisionCode
		global.UserID = userCredentials.Id
		global.DivisionCode = userCredentials.DivisionCode
	}

	routings := &[]model.ViewRoutings{}

	database.DBConn.Raw("SELECT * FROM view_routings WHERE document_tag = 'For Agenda' OR document_tag = 'Referred to Committee'").Scan(routings)

	return c.Render("dashboard_secretarial", fiber.Map{
		"pageTitle":  "Dashboard - Secretariat",
		"title":      "DASHBOARD - SECRETARIAT",
		"yearNow":    global.YearNow,
		"user":       global.Fullname,
		"userLogged": global.UserCodeLogged,
		"userId":     global.UserID,
		"greetings":  utils.GetGreetings(),
		"notifs":     routings,
		"agendas":    CountForAgenda(),
		"sessions":   CountSessions(),
		"approved":   CountApproved(),
		"baseURL":    c.BaseURL(),
	})
}

func ViewDashboardHeadOffice(c *fiber.Ctx) error {
	// Get session from storage
	session, err := global.Store.Get(c)
	if err != nil {
		panic(err)
	}
	userId, _ := session.Get("userId").(string)

	if userId == "" {
		return c.Redirect("/")
	} else {
		userCredentials := &model.UserCredentials{}
		database.DBConn.Debug().Raw("SELECT * FROM user_credentials WHeRE id = ?", userId).Scan(userCredentials)
		global.Fullname = userCredentials.Fullname
		global.UserCodeLogged = userCredentials.DivisionCode
		global.UserID = userCredentials.Id
		global.DivisionCode = userCredentials.DivisionCode
	}

	proponents := &[]model.Proponents{}
	database.DBConn.Debug().Raw("SELECT DISTINCT * FROM proponents ORDER BY proponent_id ASC").Scan(proponents)

	tracking := &[]model.Trackings{}
	database.DBConn.Debug().Raw("SELECT * FROM trackings").Scan(tracking)

	months := utils.GetMonths()
	years := utils.GetYears()

	return c.Render("dashboard_head", fiber.Map{
		"pageTitle":  "Dashboard - Head Office",
		"title":      "DASHBOARD - HEAD OFFICE",
		"monthLists": months,
		"yearLists":  years,
		"yearNow":    global.YearNow,
		"user":       global.Fullname,
		"userLogged": global.UserCodeLogged,
		"userId":     global.UserID,
		"greetings":  utils.GetGreetings(),
		"tracking":   tracking,
		"proponents": proponents,
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

func CountSessions() int64 {
	total := database.DBConn.Exec("SELECT * FROM trackings WHERE calendared IS NOT NULL").RowsAffected
	return total
}

func CountApproved() int64 {
	total := database.DBConn.Exec("SELECT * FROM trackings WHERE enacted_date IS NOT NULL").RowsAffected
	return total
}
