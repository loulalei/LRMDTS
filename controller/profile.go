package controller

import (
	"github.com/gofiber/fiber/v2"
)

func ViewProfile(c *fiber.Ctx) error {
	return c.Render("profile", fiber.Map{
		"pageTitle": "Profile",
		"title":     "PROFILE",
	})

}
