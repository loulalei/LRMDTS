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

	return c.Render("profile", fiber.Map{
		"pageTitle":   "Profile",
		"title":       "PROFILE",
		"yearNow":     model.YearNow,
		"user":        model.Fullname,
		"userLogged":  model.UserCodeLogged,
		"greetings":   utils.GetGreetings(),
		"logs":        receiving,
		"loginStatus": 101,
	})

}
