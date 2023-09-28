package controller

import (
	"github.com/gofiber/fiber/v2"
)

func SetCookie(c *fiber.Ctx) error {
	page := fiber.Cookie{
		Name:  "page",
		Value: "Dashboard",
	}
	uID := fiber.Cookie{
		Name:  "userID",
		Value: "234dfds234",
	}
	c.Cookie(&page)
	c.Cookie(&uID)
	return c.SendString("Cookie is set")
}

func GetCookie(c *fiber.Ctx) error {
	page := c.Cookies("page")
	uID := c.Cookies("userID")

	return c.SendString("This are the cookies:" + page + "|" + uID)
}
