package controller

import (
	"fmt"
	"tech_tubbies/global"
	"tech_tubbies/middleware/database"
	"tech_tubbies/model"

	"github.com/gofiber/fiber/v2"
)

func ShowCalendar(c *fiber.Ctx) error {
	return c.Render("calendar", fiber.Map{})
}

func GetEvent(c *fiber.Ctx) error {
	events := &[]model.EventCalendar{}
	database.DBConn.Debug().Raw("SELECT CONCAT(title, COUNT(*)) title, start FROM event_calendars GROUP BY start, title").Scan(events)

	endpoint := fmt.Sprintf("%v%v", c.BaseURL(), c.Path())
	return c.JSON(model.ResponseBody{
		Status:   100,
		Endpoint: endpoint,
		Message:  global.CurrentDate,
		Data:     events,
	})
}
