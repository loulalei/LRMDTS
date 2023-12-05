package controller

import (
	"tech_tubbies/model"

	"github.com/gofiber/fiber/v2"
)

func AddProponents(c *fiber.Ctx) error {
	if model.Fullname == "" {
		return c.Redirect("/")
	}
	return nil
}
