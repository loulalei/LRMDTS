package controller

import (
	"tech_tubbies/middleware/database"
	utils "tech_tubbies/middleware/util"
	"tech_tubbies/model"

	"github.com/gofiber/fiber/v2"
)

func ViewTracking(c *fiber.Ctx) error {
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	proponents := &[]model.Proponents{}
	database.DBConn.Debug().Raw("SELECT DISTINCT * FROM proponents ORDER BY proponent_id ASC").Scan(proponents)
	
	tracking := &[]model.Trackings{}
	database.DBConn.Debug().Raw("SELECT * FROM trackings").Scan(tracking)

	months := utils.GetMonths()
	years := utils.GetYears()
	return c.Render("tracking", fiber.Map{
		"pageTitle":  "Tracking",
		"title":      "TRACKING",
		"monthLists": months,
		"yearLists":  years,
		"user":       model.Fullname,
		"userLogged": model.UserCodeLogged,
		"greetings":  utils.GetGreetings(),
		"proponents": proponents,
		"tracking":   tracking,
		"status":     101,
	})
}

func DownloadAttachment(c *fiber.Ctx) error {
	filename := c.Params("filename")
	return c.Download("./assets/uploads/" + filename)
}
