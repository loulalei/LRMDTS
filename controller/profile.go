package controller

import (
	"tech_tubbies/middleware/database"
	utils "tech_tubbies/middleware/util"
	"tech_tubbies/model"

	"github.com/gofiber/fiber/v2"
)

func ViewProfile(c *fiber.Ctx) error {
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	receiving := &[]model.Routings{}
	database.DBConn.Raw("SELECT * FROM routings").Scan(receiving)

	activityLogs := &[]model.ViewActivityLogs{}
	database.DBConn.Debug().Raw("SELECT activity_id, activity, event, TO_CHAR(created_at, 'MM-DD-YYYY | HH12:MI AM') AS created_at FROM activity_loggers WHERE user_id = ? ORDER BY activity_id DESC", model.UserID).Scan(activityLogs)

	EmployeePerformance := &[]model.EmployeePerformances{}
	database.DBConn.Debug().Raw("SELECT DISTINCT TO_CHAR(created_at, 'Mon. DD,YYYY') AS date, COUNT(records_captured) as records_captured, COUNT(records_retrieved) as records_retrieved FROM employee_performaces WHERE user_id = ? GROUP BY date", model.UserID).Scan(EmployeePerformance)
	return c.Render("profile", fiber.Map{
		"pageTitle":            "Profile",
		"title":                "PROFILE",
		"yearNow":              model.YearNow,
		"user":                 model.Fullname,
		"userLogged":           model.UserCodeLogged,
		"greetings":            utils.GetGreetings(),
		"logs":                 receiving,
		"activities":           activityLogs,
		"employeePerformances": EmployeePerformance,
		"loginStatus":          101,
		"baseURL":              c.BaseURL(),
	})
}
