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
	database.DBConn.Debug().Raw("SELECT activity, event, TO_CHAR(created_at, 'MM-DD-YYYY | HH:MI:SS') AS created_at FROM activity_loggers WHERE user_id = ? ORDER BY created_at DESC", model.UserID).Scan(activityLogs)

	return c.Render("profile", fiber.Map{
		"pageTitle":   "Profile",
		"title":       "PROFILE",
		"yearNow":     model.YearNow,
		"user":        model.Fullname,
		"userLogged":  model.UserCodeLogged,
		"greetings":   utils.GetGreetings(),
		"logs":        receiving,
		"activities":  activityLogs,
		"loginStatus": 101,
		"baseURL":     c.BaseURL(),
	})

}
