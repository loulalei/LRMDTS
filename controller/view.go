package controller

import (
	"fmt"
	"tech_tubbies/middleware/database"
	utils "tech_tubbies/middleware/util"
	"tech_tubbies/model"

	"github.com/gofiber/fiber/v2"
)

func ViewDocument(c *fiber.Ctx) error {
	fmt.Println("Process: View Filing")
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	docId := c.Params("docId")
	itemNumber := c.Params("itemNumber")

	viewRoutings := &[]model.ViewRoutings{}
	database.DBConn.Debug().Raw("SELECT * FROM view_routings WHERE doc_id = ?", docId).Scan(viewRoutings)

	viewAgenda := &[]model.ViewAgenda{}
	database.DBConn.Debug().Raw("SELECT * FROM agendas WHERE item_number = ?", itemNumber).Scan(viewAgenda)

	viewApproved := &[]model.Approves{}
	database.DBConn.Debug().Raw("SELECT * FROM approves WHERE item_number = ?", itemNumber).Scan(viewApproved)

	viewReleasings := &[]model.Releasings{}
	database.DBConn.Debug().Raw("SELECT * FROM releasings WHERE doc_id = ?", docId).Scan(viewReleasings)

	viewFilings := &[]model.Filings{}
	database.DBConn.Debug().Raw("SELECT * FROM filings WHERE doc_id = ?", docId).Scan(viewFilings)

	return c.Render("view_document", fiber.Map{
		"pageTitle":      "forFiling",
		"title":          "FOR FILING",
		"yearNow":        model.YearNow,
		"user":           model.Fullname,
		"userLogged":     model.UserCodeLogged,
		"viewRoutings":   viewRoutings,
		"viewAgenda":     viewAgenda,
		"viewApproved":   viewApproved,
		"viewReleasings": viewReleasings,
		"viewFilings":    viewFilings,
		"docId":          docId,
		"itemNumber":     itemNumber,
		"greetings":      utils.GetGreetings(),
		"baseURL":        c.BaseURL(),
	})
}
