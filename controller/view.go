package controller

import (
	"fmt"
	"tech_tubbies/global"
	"tech_tubbies/middleware/database"
	utils "tech_tubbies/middleware/util"
	"tech_tubbies/model"

	"github.com/gofiber/fiber/v2"
)

func ViewDocument(c *fiber.Ctx) error {
	fmt.Println("Process: View Filing")
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

	ItemCommittees := &[]model.ViewCommittees{}
	database.DBConn.Debug().Raw("SELECT * FROM view_committees WHERE item_number = ?", itemNumber).Scan(ItemCommittees)

	// Check availability
	var agendaId int
	database.DBConn.Debug().Raw("SELECT agenda_id FROM agendas WHERE item_number = ?", itemNumber).Scan(&agendaId)

	var approveId int
	database.DBConn.Debug().Raw("SELECT approve_id FROM approves WHERE item_number = ?", itemNumber).Scan(&approveId)

	var releasingId int
	database.DBConn.Debug().Raw("SELECT releasing_id FROM releasings WHERE doc_id = ?", docId).Scan(&releasingId)

	var filingId int
	database.DBConn.Debug().Raw("SELECT filing_id FROM filings WHERE doc_id = ?", docId).Scan(&filingId)

	return c.Render("view_document", fiber.Map{
		"pageTitle":      "viewDocument",
		"title":          "View Document",
		"yearNow":        global.YearNow,
		"user":           global.Fullname,
		"userLogged":     global.UserCodeLogged,
		"divisionCode":   global.DivisionCode,
		"viewRoutings":   viewRoutings,
		"viewAgenda":     viewAgenda,
		"viewApproved":   viewApproved,
		"viewReleasings": viewReleasings,
		"viewFilings":    viewFilings,
		"agendaId":       agendaId,
		"approveId":      approveId,
		"releasingId":    releasingId,
		"filingId":       filingId,
		"docId":          docId,
		"itemNumber":     itemNumber,
		"itemCommittees": ItemCommittees,
		"greetings":      utils.GetGreetings(),
		"baseURL":        c.BaseURL(),
	})
}
