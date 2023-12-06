package controller

import (
	"fmt"
	"strconv"
	"tech_tubbies/middleware/database"
	utils "tech_tubbies/middleware/util"
	"tech_tubbies/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ViewRouting(c *fiber.Ctx) error {
	fmt.Println("Process: View Routing")
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	viewRoutings := &[]model.ViewRoutings{}
	database.DBConn.Raw("SELECT * FROM view_routings").Scan(viewRoutings)
	tracking := &[]model.ViewRoutings{}
	database.DBConn.Raw("SELECT * FROM view_routings").Scan(tracking)

	return c.Render("routing", fiber.Map{
		"pageTitle":    "Routing",
		"title":        "ROUTING MAIN",
		"yearNow":      model.YearNow,
		"user":         model.Fullname,
		"userLogged":   model.UserCodeLogged,
		"greetings":    utils.GetGreetings(),
		"viewRoutings": viewRoutings,
		"tracking":     tracking,
		"baseURL":      c.BaseURL(),
	})
}

func ViewRoutingRecords(c *fiber.Ctx) error {
	fmt.Println("Process: View Routing Records")
	if model.Fullname == "" {
		return c.Redirect("/")
	}
	return c.Render("routingRecords", fiber.Map{
		"pageTitle":  "Records",
		"title":      "RECORDS DEPARTMENT",
		"yearNow":    model.YearNow,
		"user":       model.Fullname,
		"userLogged": model.UserCodeLogged,
		"greetings":  utils.GetGreetings(),
		"baseURL":    c.BaseURL(),
	})
}

func ViewRoutingSecretariat(c *fiber.Ctx) error {
	fmt.Println("Process: View Routing Secretariat")
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	viewRoutings := &[]model.ViewRoutings{}
	database.DBConn.Debug().Raw("SELECT * FROM view_routings WHERE document_tag = 'For Agenda' OR document_tag = 'Referred to Committee'").Scan(viewRoutings)

	tracking := &[]model.ViewRoutings{}
	database.DBConn.Raw("SELECT * FROM view_routings").Scan(tracking)
	return c.Render("routingSecretariat", fiber.Map{
		"pageTitle":    "Secretariat",
		"title":        "SECRETARIAT DEPARTMENT",
		"yearNow":      model.YearNow,
		"user":         model.Fullname,
		"userLogged":   model.UserCodeLogged,
		"greetings":    utils.GetGreetings(),
		"viewRoutings": viewRoutings,
		"tracking":     tracking,
		"baseURL":      c.BaseURL(),
	})
}

func ViewRoutingForAgenda(c *fiber.Ctx) error {
	fmt.Println("Process: View Routing for Agenda")
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	docId, _ := strconv.Atoi(c.Params("docId"))

	viewRoutings := &[]model.ViewRoutings{}
	database.DBConn.Debug().Raw("SELECT * FROM view_routings WHERE document_tag = 'For Agenda' AND doc_id = ?", docId).Scan(viewRoutings)

	departments := &[]model.Departments{}
	database.DBConn.Raw("SELECT * FROM departments").Scan(departments)

	proponents := &[]model.Proponents{}
	database.DBConn.Raw("SELECT * FROM proponents").Scan(proponents)

	committees := &[]model.Committees{}
	database.DBConn.Raw("SELECT * FROM committees").Scan(committees)

	viewCommittees := &[]model.ViewCommittees{}
	database.DBConn.Raw("SELECT * FROM view_committees").Scan(viewCommittees)
	return c.Render("routingForAgenda", fiber.Map{
		"pageTitle":      "For Agenda",
		"title":          "SECRETARIAT - FOR AGENDA",
		"yearNow":        model.YearNow,
		"user":           model.Fullname,
		"userLogged":     model.UserCodeLogged,
		"userId":         model.UserID,
		"viewRoutings":   viewRoutings,
		"departments":    departments,
		"proponents":     proponents,
		"committees":     committees,
		"viewCommittees": viewCommittees,
		"docId":          docId,
		"greetings":      utils.GetGreetings(),
		"baseURL":        c.BaseURL(),
	})
}

func ViewForAgenda(c *fiber.Ctx) error {
	fmt.Println("Process: View Routing for Agenda")
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	docId, _ := strconv.Atoi(c.Params("docId"))
	itemNumber := c.Params("itemNumber")

	viewRoutings := &[]model.ViewRoutings{}
	database.DBConn.Debug().Raw("SELECT * FROM view_routings WHERE document_tag = 'Referred to Committee' AND doc_id = ?", docId).Scan(viewRoutings)

	ItemCommittees := &[]model.ViewCommittees{}
	database.DBConn.Raw("SELECT * FROM view_committees WHERE item_number = ?", itemNumber).Scan(ItemCommittees)

	viewAgenda := &[]model.ViewAgenda{}
	database.DBConn.Raw("SELECT * FROM agendas WHERE item_number = ?", itemNumber).Scan(viewAgenda)

	return c.Render("routingForAgendaUpdate", fiber.Map{
		"pageTitle":      "For Agenda",
		"title":          "SECRETARIAT - FOR AGENDA",
		"yearNow":        model.YearNow,
		"user":           model.Fullname,
		"userLogged":     model.UserCodeLogged,
		"userId":         model.UserID,
		"viewRoutings":   viewRoutings,
		"viewAgenda":     viewAgenda,
		"itemCommittees": ItemCommittees,
		"docId":          docId,
		"greetings":      utils.GetGreetings(),
		"baseURL":        c.BaseURL(),
	})
}

func ViewReceivingRoute(c *fiber.Ctx) error {
	fmt.Println("Process: View Receiving Route")
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	departments := &[]model.Departments{}

	database.DBConn.Raw("SELECT * FROM departments").Scan(departments)

	return c.Render("routingReceivingDocument", fiber.Map{
		"pageTitle":   "Receiving Document",
		"title":       "RECORDS DEPARTMENT",
		"yearNow":     model.YearNow,
		"user":        model.Fullname,
		"userLogged":  model.UserCodeLogged,
		"departments": departments,
		"greetings":   utils.GetGreetings(),
		"baseURL":     c.BaseURL(),
	})
}

func RegisterReceiving(c *fiber.Ctx) error {
	fmt.Println("Process: View Register Receiving")
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	receivingData := &model.RequestReceiving{}
	if parsErr := c.BodyParser(receivingData); parsErr != nil {
		return c.JSON(fiber.Map{
			"error": parsErr,
		})
	}

	if receivingData.ReceivingTag == "1" {
		receivingData.ReceivingTag = "For Agenda"
	} else if receivingData.ReceivingTag == "2" {
		receivingData.ReceivingTag = "For Filing"
	}

	file, err := c.FormFile("receivedFile")
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	fmt.Println("FILENAME:", file.Filename)
	c.SaveFile(file, fmt.Sprintf("./assets/uploads/%s", file.Filename))
	receivingData.ReceivedFile = file.Filename

	// INSERT NEW RECEIVING RECORD
	database.DBConn.Exec("INSERT INTO receivings (tracking_number, received_date, received_time, receiver, summary, receiving_tag, receiving_remarks, received_file, encoder) VALUES (?,?,?,?,?,?,?,?,?)",
		receivingData.TrackingNumber, receivingData.ReceivedDate, receivingData.ReceivedTime, receivingData.Receiver,
		receivingData.Summary, receivingData.ReceivingTag, receivingData.ReceivingRemarks, receivingData.ReceivedFile, model.Fullname)

	// SEARCH RECEIVING ID
	receivingId := &model.RequestRoutingIdentifications{}
	database.DBConn.Debug().Raw("SELECT receiving_id FROM receivings WHERE tracking_number = ?", receivingData.TrackingNumber).Scan(&receivingId)

	// INSERT NEW ROUTING RECORD
	database.DBConn.Exec("INSERT INTO routings (receiving_id, document_tag, remarks) VALUES (?,?,?)", receivingId.ReceivingId, receivingData.ReceivingTag, receivingData.ReceivingRemarks)

	// INSERT NEW TRACKING RECORd
	database.DBConn.Exec("INSERT INTO trackings (tracking_number, summary, received_date) VALUES (?,?,?)", receivingData.TrackingNumber, receivingData.Summary, receivingData.ReceivedDate)
	// VIEW RESULTS
	viewRoutings := &[]model.ViewRoutings{}
	database.DBConn.Raw("SELECT * FROM view_routings").Scan(viewRoutings)

	return c.Redirect("/api/routing")
}

func GetForFiling(c *fiber.Ctx) error {
	fmt.Println("Process: Get for Filing")
	if model.Fullname == "" {
		return c.Redirect("/")
	}
	id := c.Params("id")
	receiving := &[]model.Routings{}

	database.DBConn.Raw("SELECT * FROM routings WHERE doc_id=?", id).Scan(receiving)

	folders := &[]model.Folders{}
	database.DBConn.Raw("SELECT * FROM folders").Scan(folders)
	return c.Render("routingFiling", fiber.Map{
		"pageTitle":  "forFiling",
		"title":      "FOR FILING",
		"yearNow":    model.YearNow,
		"user":       model.Fullname,
		"userLogged": model.UserCodeLogged,
		"greetings":  utils.GetGreetings(),
		"receivings": receiving,
		"folders":    folders,
		"baseURL":    c.BaseURL(),
	})
}

func RegisterForAgenda(c *fiber.Ctx) error {
	dateNow := time.Now()
	fmt.Println("Process: Register for Agenda")
	if model.Fullname == "" {
		return c.Redirect("/")
	}
	requestAgenda := &model.RequestForAgenda{}
	if parsErr := c.BodyParser(requestAgenda); parsErr != nil {
		return c.JSON(model.ResponseBody{
			Status:  101,
			Message: "parsing request error",
			Request: requestAgenda,
			Data:    parsErr.Error(),
		})
	}

	if requestAgenda.AgendaTag == "1" {
		requestAgenda.AgendaTag = "For information of the whole body"
	} else if requestAgenda.AgendaTag == "2" {
		requestAgenda.AgendaTag = "Referred to Committee"
	} else if requestAgenda.AgendaTag == "3" {
		requestAgenda.AgendaTag = "Approved"
	}

	if requestAgenda.Source == "1" {
		requestAgenda.Source = "Other Department"
		requestAgenda.SourceResult = requestAgenda.Department
	} else if requestAgenda.Source == "2" {
		requestAgenda.Source = "Priviledged Speech"
		requestAgenda.SourceResult = requestAgenda.Proponent
	} else if requestAgenda.Source == "3" {
		requestAgenda.Source = "Proposed"
		requestAgenda.SourceResult = requestAgenda.Proponent
	} else if requestAgenda.Source == "4" {
		requestAgenda.Source = "Other"
		requestAgenda.SourceResult = requestAgenda.Other
	}

	database.DBConn.Debug().Exec("INSERT INTO agendas (item_number, is_urgent, date_calendared, date_reported, agenda_tag, agenda_remarks, source, source_result, encoder) VALUES (?,?,?,?,?,?,?,?,?)",
		requestAgenda.ItemNumber, requestAgenda.IsUrgent, requestAgenda.DateCalendared, requestAgenda.DateReported, requestAgenda.AgendaTag, requestAgenda.AgendaRemarks, requestAgenda.Source, requestAgenda.SourceResult, model.Fullname)

	// SEARCH AGENDA ID
	agendaId := &model.RequestRoutingIdentifications{}
	database.DBConn.Debug().Raw("SELECT agenda_id FROM agendas WHERE item_number = ?", requestAgenda.ItemNumber).Scan(agendaId)

	fmt.Println("DOC ID:", requestAgenda.DocId)
	// UPDATE ROUTINGS
	database.DBConn.Debug().Exec("UPDATE routings SET agenda_id = ?, document_tag = ?, remarks = ?, item_number = ?, updated_at = ? WHERE doc_id = ?", agendaId.AgendaId, requestAgenda.AgendaTag, requestAgenda.AgendaRemarks, requestAgenda.ItemNumber, dateNow, requestAgenda.DocId)

	// UPDATE TRACKING
	database.DBConn.Debug().Exec("UPDATE trackings SET item_number =?, calendared = ? WHERE tracking_number = ?", requestAgenda.ItemNumber, requestAgenda.DateCalendared, requestAgenda.TrackingNumber)

	viewroutings := &[]model.ViewRoutings{}
	database.DBConn.Debug().Raw("SELECT * FROM view_routings").Scan(viewroutings)
	//------------------
	receivings := &[]model.Routings{}
	database.DBConn.Debug().Raw("SELECT * FROM routings").Scan(receivings)

	return c.Redirect("/api/routing/secretariat")
}

func UpdateForAgenda(c *fiber.Ctx) error {
	fmt.Println("Process: Register for Agenda")
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	requestAgenda := &model.RequestForAgenda{}
	if parsErr := c.BodyParser(requestAgenda); parsErr != nil {
		return c.JSON(model.ResponseBody{
			Status:  101,
			Message: "parsing request error",
			Request: requestAgenda,
			Data:    parsErr.Error(),
		})
	}

	return c.JSON(model.ResponseBody{
		Status:  100,
		Message: "success",
		Request: requestAgenda,
	})
}

func ViewApproved(c *fiber.Ctx) error {
	fmt.Println("Process: View Approved")
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	docId, _ := strconv.Atoi(c.Params("docId"))
	itemNumber := c.Params("itemNumber")

	viewRoutings := &[]model.ViewRoutings{}
	database.DBConn.Raw("SELECT * FROM view_routings WHERE document_tag = 'Approved' AND doc_id = ?", docId).Scan(viewRoutings)

	ItemCommittees := &[]model.ViewCommittees{}
	database.DBConn.Raw("SELECT * FROM view_committees WHERE item_number = ?", itemNumber).Scan(ItemCommittees)

	viewAgenda := &[]model.ViewAgenda{}
	database.DBConn.Raw("SELECT * FROM agendas WHERE item_number = ?", itemNumber).Scan(viewAgenda)

	proponents := &[]model.Proponents{}
	database.DBConn.Raw("SELECT * FROM proponents").Scan(proponents)

	return c.Render("routingApproved", fiber.Map{
		"pageTitle":      "Document Approved",
		"title":          "APPROVED",
		"yearNow":        model.YearNow,
		"user":           model.Fullname,
		"userLogged":     model.UserCodeLogged,
		"userId":         model.UserID,
		"viewRoutings":   viewRoutings,
		"viewAgenda":     viewAgenda,
		"itemCommittees": ItemCommittees,
		"proponents":     proponents,
		"docId":          docId,
		"greetings":      utils.GetGreetings(),
		"baseURL":        c.BaseURL(),
	})
}

func ViewReleasing(c *fiber.Ctx) error {
	fmt.Println("Process: View Releasing")
	if model.Fullname == "" {
		return c.Redirect("/")
	}
	return c.Render("routingReleasing", fiber.Map{
		"pageTitle":  "Routing - Releasing",
		"title":      "ROUTING RELEASING",
		"yearNow":    model.YearNow,
		"user":       model.Fullname,
		"userLogged": model.UserCodeLogged,
		"greetings":  utils.GetGreetings(),
		"baseURL":    c.BaseURL(),
	})
}

func InsertCommitteeForAgenda(c *fiber.Ctx) error {
	if model.Fullname == "" {
		return c.Redirect("/")
	}
	itemNo := c.Params("itemNo")
	committeeId := c.Params("committeeId")
	userId := c.Params("userId")

	// count if over 5 committees
	var count int64
	if dbErr := database.DBConn.Raw("SELECT COUNT(*) FROM view_committees").Scan(&count); dbErr.Error != nil {
		fmt.Println("error:", dbErr.Error)
	}

	fmt.Println("Count:", count)
	if count > 5 {
		return c.JSON(model.ResponseBody{
			Status:  101,
			Message: "error adding committee",
		})
	} else {

		// Insert new committees
		database.DBConn.Exec("INSERT INTO committee_lists (item_number, committee_id, user_id) VALUES (?,?,?)", itemNo, committeeId, userId)

		// Get all committees
		committeesLists := &model.ViewCommittees{}
		database.DBConn.Debug().Raw("SELECT * FROM view_committees WHERE item_number = ? ORDER BY list_id DESC", itemNo).Scan(committeesLists)
		return c.JSON(model.ResponseBody{
			Status:  100,
			Message: "show added committees",
			Request: fiber.Map{
				"itemNo":      itemNo,
				"committeeId": committeeId,
				"userId":      userId,
			},
			Data: committeesLists,
		})
	}
}

func PostInsertCommitteeForAgenda(c *fiber.Ctx) error {
	requestBody := &model.RequestCommittee{}
	if parsErr := c.BodyParser(requestBody); parsErr != nil {
		return c.JSON(fiber.Map{
			"message": "error parsing",
			"data":    parsErr.Error(),
		})
	}

	// count if over 5 committees
	var count int64
	if dbErr := database.DBConn.Raw("SELECT COUNT(*) FROM view_committees WHERE item_number = ?", requestBody.ItemNumber).Scan(&count); dbErr.Error != nil {
		message := fmt.Sprintf("Error database: %v", dbErr.Error)
		return c.JSON(model.ResponseBody{
			Status:  101,
			Message: message,
		})
	}
	if count > 2 {
		return c.JSON(model.ResponseBody{
			Status:  101,
			Message: "Maximum of 3 committees only",
		})
	} else {
		committeesLists := &model.ViewCommittees{}
		// check if existing in the view committee
		database.DBConn.Debug().Raw("SELECT COUNT(*) FROM view_committees WHERE committee_id = ?", requestBody.CommitteeId).Scan(&count)
		if count > 0 {
			return c.JSON(model.ResponseBody{
				Status:  101,
				Message: "This committee is already in the list",
			})
		} else {
			// Insert new committees
			database.DBConn.Exec("INSERT INTO committee_lists (item_number, committee_id, user_id) VALUES (?,?,?)", requestBody.ItemNumber, requestBody.CommitteeId, requestBody.UserId)

			// Get all committees
			database.DBConn.Debug().Raw("SELECT * FROM view_committees WHERE item_number = ? ORDER BY list_id DESC", requestBody.ItemNumber).Scan(committeesLists)
			message := fmt.Sprintf("%v has been added", committeesLists.Name)
			return c.JSON(model.ResponseBody{
				Status:  100,
				Message: message,
				Request: fiber.Map{
					"itemNo":      requestBody.ItemNumber,
					"committeeId": requestBody.CommitteeId,
					"userId":      requestBody.UserId,
					"listId":      committeesLists.ListId,
				},
				Data: committeesLists,
			})
		}
	}
}

func RemoveCommitteeForAgenda(c *fiber.Ctx) error {
	itemNo := c.Params("itemNo")
	committeeId := c.Params("committeeId")
	database.DBConn.Exec("DELETE FROM committee_lists WHERE committee_id = ? AND item_number = ?", committeeId, itemNo)
	return c.JSON(fiber.Map{
		"message": "successs",
		"ItemNo":  itemNo,
	})
}
