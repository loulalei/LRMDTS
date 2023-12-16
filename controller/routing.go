package controller

import (
	"fmt"
	"mime/multipart"
	"strconv"
	"tech_tubbies/middleware/database"
	utils "tech_tubbies/middleware/util"
	"tech_tubbies/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

// ------------------------
// RECORDS
// ------------------------
func ViewRouting(c *fiber.Ctx) error {
	fmt.Println("Process: View Routing")
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	viewRoutings := &[]model.ViewRoutings{}
	database.DBConn.Debug().Raw("SELECT * FROM view_routings ORDER BY updated_at DESC").Scan(viewRoutings)

	return c.Render("routing", fiber.Map{
		"pageTitle":    "Routing",
		"title":        "ROUTING MAIN",
		"yearNow":      model.YearNow,
		"user":         model.Fullname,
		"userLogged":   model.UserCodeLogged,
		"greetings":    utils.GetGreetings(),
		"viewRoutings": viewRoutings,
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

// ------------------------
// SECRETARIAT
// ------------------------
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
	database.DBConn.Debug().Raw("SELECT * FROM view_routings WHERE doc_id = ?", docId).Scan(viewRoutings)

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

// ------------------------
// RECEIVING
// ------------------------
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
	database.DBConn.Debug().Exec("INSERT INTO receivings (tracking_number, received_date, received_time, receiver, summary, receiving_tag, receiving_remarks, received_file, encoder) VALUES (?,?,?,?,?,?,?,?,?)",
		receivingData.TrackingNumber, receivingData.ReceivedDate, receivingData.ReceivedTime, receivingData.Receiver,
		receivingData.Summary, "For Agenda", "Forwarded to Secretariat", receivingData.ReceivedFile, model.Fullname,
	)

	// GET RECEIVING ID
	var receivingID int
	database.DBConn.Debug().Raw("SELECT receiving_id FROM receivings WHERE tracking_number = ?", receivingData.TrackingNumber).Scan(&receivingID)
	// INSERT NEW ROUTING RECORD
	database.DBConn.Debug().Exec("INSERT INTO routings (receiving_id, document_tag, remarks) VALUES (?,?,?) ", receivingID, "For Agenda", "Forwarded to Secretariat")
	// INSERT NEW TRACKING RECORD
	database.DBConn.Exec("INSERT INTO trackings (tracking_number, summary, received_date) VALUES (?,?,?)", receivingData.TrackingNumber, receivingData.Summary, receivingData.ReceivedDate)
	// VIEW RESULTS
	viewRoutings := &[]model.ViewRoutings{}
	database.DBConn.Raw("SELECT * FROM view_routings").Scan(viewRoutings)

	return c.Redirect("/api/routing")
}

// ------------------------
// FOR AGENDA
// ------------------------
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
	database.DBConn.Debug().Raw("SELECT * FROM view_committees WHERE item_number = ?", itemNumber).Scan(ItemCommittees)

	viewAgenda := &[]model.ViewAgenda{}
	database.DBConn.Debug().Raw("SELECT * FROM agendas WHERE item_number = ?", itemNumber).Scan(viewAgenda)

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
		"itemNumber":     itemNumber,
		"greetings":      utils.GetGreetings(),
		"baseURL":        c.BaseURL(),
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

	agendaFields := &model.Agendas{}
	database.DBConn.Debug().Exec("INSERT INTO agendas (item_number, is_urgent, date_calendared, date_reported, agenda_tag, agenda_remarks, source, source_result, encoder) VALUES (?,?,?,?,?,?,?,?,?)",
		requestAgenda.ItemNumber, requestAgenda.IsUrgent, requestAgenda.DateCalendared, requestAgenda.DateReported,
		requestAgenda.AgendaTag, requestAgenda.AgendaRemarks, requestAgenda.Source, requestAgenda.SourceResult, model.Fullname,
	).Find(agendaFields)

	// UPDATE ROUTINGS
	database.DBConn.Debug().Exec("UPDATE routings SET agenda_id = ?, document_tag = ?, remarks = ?, item_number = ?, updated_at = ? WHERE doc_id = ?", agendaFields.AgendaId, requestAgenda.AgendaTag, requestAgenda.AgendaRemarks, requestAgenda.ItemNumber, dateNow, requestAgenda.DocId)

	// UPDATE TRACKING
	database.DBConn.Debug().Exec("UPDATE trackings SET doc_id = ?, item_number =?, calendared = ?, updated_at = ? WHERE tracking_number = ?", requestAgenda.DocId, requestAgenda.ItemNumber, requestAgenda.DateCalendared, dateNow, requestAgenda.TrackingNumber)

	viewroutings := &[]model.ViewRoutings{}
	database.DBConn.Debug().Raw("SELECT * FROM view_routings").Scan(viewroutings)
	//------------------
	receivings := &[]model.Routings{}
	database.DBConn.Debug().Raw("SELECT * FROM routings").Scan(receivings)

	return c.Redirect("/api/routing/secretariat")
}

func UpdateForAgenda(c *fiber.Ctx) error {
	dateNow := time.Now()
	fmt.Println("Process: Update for Agenda")
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

	database.DBConn.Debug().Exec("UPDATE agendas SET agenda_tag = ?, agenda_remarks = ?, modified_by = ?,updated_at = ? WHERE item_number = ?", requestAgenda.AgendaTag, requestAgenda.AgendaRemarks, model.Fullname, dateNow, requestAgenda.ItemNumber)
	database.DBConn.Debug().Exec("UPDATE routings SET document_tag = ?, remarks = ?, updated_at = ? WHERE item_number = ?", requestAgenda.AgendaTag, requestAgenda.AgendaRemarks, dateNow, requestAgenda.ItemNumber)

	return c.Redirect("/api/routing/secretariat")
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
		database.DBConn.Debug().Raw("SELECT COUNT(*) FROM view_committees WHERE committee_id = ? AND item_number = ?", requestBody.CommitteeId, requestBody.ItemNumber).Scan(&count)
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

// ------------------------
// APPROVED
// ------------------------
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
		"itemNumber":     itemNumber,
		"greetings":      utils.GetGreetings(),
		"baseURL":        c.BaseURL(),
	})
}

func RegisterApproved(c *fiber.Ctx) error {
	dateNow := time.Now()
	fmt.Println("Process: Register Approved")
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	requestApproved := &model.RequestApproved{}
	if parsErr := c.BodyParser(requestApproved); parsErr != nil {
		return c.JSON(model.ResponseBody{
			Status:  101,
			Message: "error parsing data",
			Data:    parsErr.Error(),
		})
	}

	if requestApproved.LawType == "1" {
		requestApproved.LawType = "Ordinance"
	} else if requestApproved.LawType == "2" {
		requestApproved.LawType = "Resolution"
	}

	file, err := c.FormFile("attachedFile")
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.SaveFile(file, fmt.Sprintf("./assets/uploads/%s", file.Filename))
	requestApproved.ResOrdFile = file.Filename

	// Insert new data for approved
	approvedFields := &model.Approves{}
	database.DBConn.Debug().Exec("INSERT INTO approves (law_type, law_number, series, enacted_date, motioned_by, author, res_ord_file, title_body, encoder, item_number) VALUES (?,?,?,?,?,?,?,?,?,?)",
		requestApproved.LawType, requestApproved.LawNumber, requestApproved.Series,
		requestApproved.EnactedDate, requestApproved.MotionedBy, requestApproved.Author,
		requestApproved.ResOrdFile, requestApproved.TitleBody, requestApproved.Encoder,
		requestApproved.ItemNumber,
	).Find(approvedFields)

	// GET APPROVED ID
	var approvedId int
	database.DBConn.Debug().Raw("SELECT agenda_id FROM agendas WHERE item_number = ?", requestApproved.ItemNumber).Scan(&approvedId)

	// Update Routing
	database.DBConn.Debug().Exec("UPDATE routings SET approved_id = ?, document_tag = ?, remarks = ?, updated_at = ? WHERE doc_id = ?", approvedId, "For Releasing", "Kept in Records", dateNow, requestApproved.DocId)

	// Update Tracking
	database.DBConn.Debug().Exec("UPDATE trackings SET law_type = ?, law_number = ?, series = ?, enacted_date = ?, author = ?, motioned_by = ? , updated_at = ? WHERE item_number = ?", requestApproved.LawType, requestApproved.LawNumber, requestApproved.Series, requestApproved.EnactedDate, requestApproved.Author, requestApproved.MotionedBy, dateNow, requestApproved.ItemNumber)

	return c.Redirect("/api/routing")
}

// ------------------------
// RELEASING
// ------------------------
func ViewReleasing(c *fiber.Ctx) error {
	fmt.Println("Process: View Releasing")
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	docId, _ := strconv.Atoi(c.Params("docId"))
	itemNumber := c.Params("itemNumber")

	viewRoutings := &[]model.ViewRoutings{}
	database.DBConn.Debug().Raw("SELECT * FROM view_routings WHERE document_tag = 'For Releasing' AND doc_id = ?", docId).Scan(viewRoutings)

	ItemCommittees := &[]model.ViewCommittees{}
	database.DBConn.Debug().Raw("SELECT * FROM view_committees WHERE item_number = ?", itemNumber).Scan(ItemCommittees)

	viewAgenda := &[]model.ViewAgenda{}
	database.DBConn.Debug().Raw("SELECT * FROM agendas WHERE item_number = ?", itemNumber).Scan(viewAgenda)

	viewApproved := &[]model.Approves{}
	database.DBConn.Debug().Raw("SELECT * FROM approves WHERE item_number = ?", itemNumber).Scan(viewApproved)

	return c.Render("releasing", fiber.Map{
		"pageTitle":      "Routing - Releasing",
		"title":          "ROUTING RELEASING",
		"yearNow":        model.YearNow,
		"user":           model.Fullname,
		"userLogged":     model.UserCodeLogged,
		"viewRoutings":   viewRoutings,
		"viewAgenda":     viewAgenda,
		"viewApproved":   viewApproved,
		"itemCommittees": ItemCommittees,
		"docId":          docId,
		"itemNumber":     itemNumber,
		"greetings":      utils.GetGreetings(),
		"baseURL":        c.BaseURL(),
	})
}

func RegisterReleasing(c *fiber.Ctx) error {
	dateNow := time.Now()
	fmt.Println("Process: Register Releasing")
	var spFile, endorsementFile *multipart.FileHeader
	var err error
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	requestReleasing := &model.RequestReleasing{}
	if parsErr := c.BodyParser(requestReleasing); parsErr != nil {
		return c.JSON(model.ResponseBody{
			Status:  101,
			Message: "error parsing data",
			Data:    parsErr.Error(),
		})
	}

	spFile, err = c.FormFile("spResolutionFile")
	if err != nil {
		fmt.Println("Error SP attachment:", err.Error())
	} else {
		c.SaveFile(spFile, fmt.Sprintf("./assets/uploads/%s", spFile.Filename))
		requestReleasing.SPResolutionFile = spFile.Filename
	}

	endorsementFile, err = c.FormFile("endorsementFile")
	if err != nil {
		fmt.Println("Error LOCAL attachment:", err.Error())
	} else {
		c.SaveFile(endorsementFile, fmt.Sprintf("./assets/uploads/%s", endorsementFile.Filename))
		requestReleasing.EndorsementFile = endorsementFile.Filename
	}

	if requestReleasing.IsApprovedLachesMayor {
		requestReleasing.MayorDateApproved = ""
	}

	if requestReleasing.IsApprovedLachesSP {
		requestReleasing.SPDateApproved = ""
	}

	if requestReleasing.MayorDateForwarded != "" && (!requestReleasing.IsApprovedLachesMayor || requestReleasing.MayorDateApproved == "") {
		database.DBConn.Debug().Exec("UPDATE routings SET remarks = 'Forwarded to Mayor', updated_at = ? WHERE doc_id = ?", dateNow, requestReleasing.DocId)
	}

	if requestReleasing.MayorDateForwarded != "" && (requestReleasing.IsApprovedLachesMayor || requestReleasing.MayorDateApproved != "") {
		database.DBConn.Debug().Exec("UPDATE routings SET remarks = 'Approved by Mayor', updated_at = ? WHERE doc_id = ?", dateNow, requestReleasing.DocId)
	}

	if requestReleasing.SPDateForwarded != "" && (!requestReleasing.IsApprovedLachesSP || requestReleasing.SPDateApproved == "") {
		database.DBConn.Debug().Exec("UPDATE routings SET remarks = 'Forwarded to Panlalawigan', updated_at = ?  WHERE doc_id = ?", dateNow, requestReleasing.DocId)
	}

	if requestReleasing.SPDateForwarded != "" && (requestReleasing.IsApprovedLachesSP || requestReleasing.SPDateApproved != "") {
		database.DBConn.Debug().Exec("UPDATE routings SET remarks = 'Approved by Panlalawigan', updated_at = ?  WHERE doc_id = ?", dateNow, requestReleasing.DocId)
	}

	if requestReleasing.LocalDateRelease != "" {
		database.DBConn.Debug().Exec("UPDATE routings SET documrequestReleasingent_tag = 'For Filing', remarks = 'Kept in Records', updated_at = ? WHERE doc_id = ?", dateNow, requestReleasing.DocId)
	}

	database.DBConn.Debug().Exec("INSERT INTO releasings (doc_id,mayor_date_forwarded, mayor_date_approved, is_approved_laches_mayor, sp_date_forwarded, sp_date_approved, is_approved_laches_sp, sp_resolution_number, sp_resolution_file, local_date_release, local_date_published, endorsement_file, encoder) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)",
		requestReleasing.DocId, requestReleasing.MayorDateForwarded, requestReleasing.MayorDateApproved, requestReleasing.IsApprovedLachesMayor,
		requestReleasing.SPDateForwarded, requestReleasing.SPDateApproved, requestReleasing.IsApprovedLachesSP, requestReleasing.SPResolutionNumber,
		requestReleasing.SPResolutionFile, requestReleasing.LocalDateRelease, requestReleasing.LocalDatePublished, requestReleasing.EndorsementFile, model.Fullname)

	var releasingId int
	database.DBConn.Debug().Raw("SELECT releasing_id FROM releasings WHERE doc_id = ?", requestReleasing.DocId).Scan(releasingId)

	database.DBConn.Debug().Exec("UPDATE routings SET releasing_id = ?, updated_at = ? WHRERE doc_id = ?", releasingId, dateNow, requestReleasing.DocId)
	database.DBConn.Debug().Exec("UPDATE trackings SET mayor_date = ?, sta_cruz_date = ?, released_date = ?, published_date = ?, updated_at = ? WHERE item_number = ?", requestReleasing.MayorDateForwarded, requestReleasing.SPDateForwarded, requestReleasing.LocalDateRelease, requestReleasing.LocalDatePublished, dateNow, requestReleasing.ItemNumber)

	return c.Redirect("/api/routing")
}

func UpdateReleasing(c *fiber.Ctx) error {
	fmt.Println("Process: Update Releasing")
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	docId, _ := strconv.Atoi(c.Params("docId"))
	itemNumber := c.Params("itemNumber")

	viewRoutings := &[]model.ViewRoutings{}
	database.DBConn.Debug().Raw("SELECT * FROM view_routings WHERE document_tag = 'For Releasing' AND doc_id = ?", docId).Scan(viewRoutings)

	ItemCommittees := &[]model.ViewCommittees{}
	database.DBConn.Debug().Raw("SELECT * FROM view_committees WHERE item_number = ?", itemNumber).Scan(ItemCommittees)

	viewAgenda := &[]model.ViewAgenda{}
	database.DBConn.Debug().Raw("SELECT * FROM agendas WHERE item_number = ?", itemNumber).Scan(viewAgenda)

	viewApproved := &[]model.Approves{}
	database.DBConn.Debug().Raw("SELECT * FROM approves WHERE item_number = ?", itemNumber).Scan(viewApproved)

	releasing := &[]model.Releasings{}
	database.DBConn.Debug().Raw("SELECT * FROM releasings WHERE doc_id = ?", docId).Scan(releasing)

	return c.Render("updatereleasing", fiber.Map{
		"pageTitle":      "Routing - Releasing",
		"title":          "ROUTING RELEASING",
		"yearNow":        model.YearNow,
		"user":           model.Fullname,
		"userLogged":     model.UserCodeLogged,
		"viewRoutings":   viewRoutings,
		"viewAgenda":     viewAgenda,
		"viewApproved":   viewApproved,
		"releasing":      releasing,
		"itemCommittees": ItemCommittees,
		"docId":          docId,
		"itemNumber":     itemNumber,
		"greetings":      utils.GetGreetings(),
		"baseURL":        c.BaseURL(),
	})
}

func SaveReleasing(c *fiber.Ctx) error {
	dateNow := time.Now()
	fmt.Println("Process: Save Releasing")
	var spFile, endorsementFile *multipart.FileHeader
	var err error
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	requestReleasing := &model.RequestReleasing{}
	if parsErr := c.BodyParser(requestReleasing); parsErr != nil {
		return c.JSON(model.ResponseBody{
			Status:  101,
			Message: "error parsing data",
			Data:    parsErr.Error(),
		})
	}

	spFile, err = c.FormFile("spResolutionFile")
	if err != nil {
		fmt.Println("Error SP attachment:", err.Error())
	} else {
		c.SaveFile(spFile, fmt.Sprintf("./assets/uploads/%s", spFile.Filename))
		requestReleasing.SPResolutionFile = spFile.Filename
	}

	endorsementFile, err = c.FormFile("endorsementFile")
	if err != nil {
		fmt.Println("Error LOCAL attachment:", err.Error())
	} else {
		c.SaveFile(endorsementFile, fmt.Sprintf("./assets/uploads/%s", endorsementFile.Filename))
		requestReleasing.EndorsementFile = endorsementFile.Filename
	}

	if requestReleasing.IsApprovedLachesMayor {
		requestReleasing.MayorDateApproved = ""
	}

	if requestReleasing.IsApprovedLachesSP {
		requestReleasing.SPDateApproved = ""
	}

	if requestReleasing.MayorDateForwarded != "" && (!requestReleasing.IsApprovedLachesMayor || requestReleasing.MayorDateApproved == "") {
		fmt.Println("mayor forwarded")
		database.DBConn.Debug().Exec("UPDATE routings SET remarks = 'Forwarded to Mayor', updated_at = ? WHERE doc_id = ?", dateNow, requestReleasing.DocId)
	}

	if requestReleasing.MayorDateForwarded != "" && (requestReleasing.IsApprovedLachesMayor || requestReleasing.MayorDateApproved != "") {
		fmt.Println("mayor approved")
		database.DBConn.Debug().Exec("UPDATE routings SET remarks = 'Approved by Mayor', updated_at = ? WHERE doc_id = ?", dateNow, requestReleasing.DocId)
	}

	if requestReleasing.SPDateForwarded != "" && (!requestReleasing.IsApprovedLachesSP || requestReleasing.SPDateApproved == "") {
		fmt.Println("sta cruz forwarded")
		database.DBConn.Debug().Exec("UPDATE routings SET remarks = 'Forwarded to Panlalawigan', updated_at = ?  WHERE doc_id = ?", dateNow, requestReleasing.DocId)
	}

	if requestReleasing.SPDateForwarded != "" && (requestReleasing.IsApprovedLachesSP || requestReleasing.SPDateApproved != "") {
		fmt.Println("sta cruz approved")
		database.DBConn.Debug().Exec("UPDATE routings SET remarks = 'Approved by Panlalawigan', updated_at = ?  WHERE doc_id = ?", dateNow, requestReleasing.DocId)
	}

	if requestReleasing.LocalDateRelease != "" {
		fmt.Println("local released")
		database.DBConn.Debug().Exec("UPDATE routings SET document_tag = 'For Filing', remarks = 'Kept in Records', updated_at = ?  WHERE doc_id = ?", dateNow, requestReleasing.DocId)
	}

	database.DBConn.Debug().Exec("UPDATE releasings SET mayor_date_forwarded = ?, mayor_date_approved = ?, is_approved_laches_mayor = ?, sp_date_forwarded = ?, sp_date_approved = ?, is_approved_laches_sp = ?, sp_resolution_number = ?, sp_resolution_file = ?, local_date_release = ?, local_date_published = ?, endorsement_file = ?, modified_by = ? WHERE doc_id = ?",
		requestReleasing.MayorDateForwarded, requestReleasing.MayorDateApproved, requestReleasing.IsApprovedLachesMayor,
		requestReleasing.SPDateForwarded, requestReleasing.SPDateApproved, requestReleasing.IsApprovedLachesSP, requestReleasing.SPResolutionNumber,
		requestReleasing.SPResolutionFile, requestReleasing.LocalDateRelease, requestReleasing.LocalDatePublished, requestReleasing.EndorsementFile, model.Fullname, requestReleasing.DocId)

	var releasingId int
	database.DBConn.Debug().Raw("SELECT releasing_id FROM releasings WHERE doc_id = ?", releasingId)

	database.DBConn.Debug().Exec("UPDATE trackings SET mayor_date = ?, sta_cruz_date = ?, released_date = ?, published_date = ?, updated_at = ? WHERE item_number = ?", requestReleasing.MayorDateForwarded, requestReleasing.SPDateForwarded, requestReleasing.LocalDateRelease, requestReleasing.LocalDatePublished, dateNow, requestReleasing.ItemNumber)

	return c.Redirect("/api/routing")
}

// ------------------------
// FILING
// ------------------------
func ViewForFiling(c *fiber.Ctx) error {
	fmt.Println("Process: View Filing")
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	docId := c.Params("docId")
	itemNumber := c.Params("itemNumber")

	viewRoutings := &[]model.ViewRoutings{}
	database.DBConn.Debug().Raw("SELECT * FROM view_routings WHERE doc_id = ?", docId).Scan(viewRoutings)

	ItemCommittees := &[]model.ViewCommittees{}
	database.DBConn.Debug().Raw("SELECT * FROM view_committees WHERE item_number = ?", itemNumber).Scan(ItemCommittees)

	viewAgenda := &[]model.ViewAgenda{}
	database.DBConn.Debug().Raw("SELECT * FROM agendas WHERE item_number = ?", itemNumber).Scan(viewAgenda)

	viewApproved := &[]model.Approves{}
	database.DBConn.Debug().Raw("SELECT * FROM approves WHERE item_number = ?", itemNumber).Scan(viewApproved)

	viewReleasings := &[]model.Releasings{}
	database.DBConn.Debug().Raw("SELECT * FROM releasings WHERE doc_id = ?", docId).Scan(viewReleasings)

	viewFilings := &[]model.Filings{}
	database.DBConn.Debug().Raw("SELECT * FROM filings WHERE doc_id = ?", docId).Scan(viewFilings)

	folders := &[]model.Folders{}
	database.DBConn.Debug().Raw("SELECT * FROM folders").Scan(folders)

	return c.Render("filing", fiber.Map{
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
		"itemCommittees": ItemCommittees,
		"docId":          docId,
		"itemNumber":     itemNumber,
		"folder":         folders,
		"greetings":      utils.GetGreetings(),
		"baseURL":        c.BaseURL(),
	})
}

func RegisterFiling(c *fiber.Ctx) error {
	dateNow := time.Now()
	fmt.Println("Process: Register Filing")
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	requestFiling := &model.RequestFiling{}
	if parsErr := c.BodyParser(requestFiling); parsErr != nil {
		c.JSON(model.ResponseBody{
			Status:  101,
			Message: "error parsing request",
			Data:    parsErr.Error(),
		})
	}

	if requestFiling.FolderName == "new" {
		requestFiling.FolderName = requestFiling.NewFolderName
		database.DBConn.Debug().Exec("INSERT INTO folders (name) VALUES (?)", requestFiling.NewFolderName)
	}

	if !requestFiling.IsBorrowed {
		requestFiling.Borrower = ""
		requestFiling.DateBorrowed = ""
	}
	filingFields := &model.Filings{
		DocId:         requestFiling.DocId,
		CabinetNumber: requestFiling.CabinetNumber,
		FolderName:    requestFiling.FolderName,
		IsBorrowed:    requestFiling.IsBorrowed,
		DateBorrowed:  requestFiling.DateBorrowed,
		Borrower:      requestFiling.Borrower,
		DateFiled:     requestFiling.DateFiled,
		DatePublished: requestFiling.DatePublished,
		Publisher:     requestFiling.Publisher,
		Encoder:       model.Fullname,
		CreatedAt:     dateNow,
		UpdatedAt:     dateNow,
	}

	database.DBConn.Debug().Exec("INSERT INTO filings (doc_id, cabinet_number, folder_name, date_filed, is_borrowed, date_borrowed, borrower, date_published, publisher, encoder) VALUES (?,?,?,?,?,?,?,?,?,?)",
		filingFields.DocId, filingFields.CabinetNumber, filingFields.FolderName, filingFields.DateFiled,
		filingFields.IsBorrowed, filingFields.DateBorrowed, filingFields.Borrower,
		filingFields.DatePublished, filingFields.Publisher, filingFields.Encoder)

	var filingId int
	database.DBConn.Debug().Raw("SELECT filing_id FROM filings WHERE doc_id = ?", filingFields.DocId).Scan(&filingId)

	database.DBConn.Debug().Exec("UPDATE routings SET filing_id = ?, document_tag = ?, remarks = ?, updated_at = ? WHERE doc_id = ?", filingId, "Filed", "Kept in Records", dateNow, filingFields.DocId)
	database.DBConn.Debug().Exec("UPDATE trackings SET filed_date = ?, published_date = ?, updated_at = ? WHERE doc_id = ?", filingFields.DateFiled, filingFields.DatePublished, dateNow, filingFields.DocId)

	return c.Redirect("/api/routing")
}

func UpdateFiling(c *fiber.Ctx) error {
	fmt.Println("Process: Update Filing")
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	docId := c.Params("docId")
	itemNumber := c.Params("itemNumber")

	viewRoutings := &[]model.ViewRoutings{}
	database.DBConn.Debug().Raw("SELECT * FROM view_routings WHERE doc_id = ?", docId).Scan(viewRoutings)

	ItemCommittees := &[]model.ViewCommittees{}
	database.DBConn.Debug().Raw("SELECT * FROM view_committees WHERE item_number = ?", itemNumber).Scan(ItemCommittees)

	viewAgenda := &[]model.ViewAgenda{}
	database.DBConn.Debug().Raw("SELECT * FROM agendas WHERE item_number = ?", itemNumber).Scan(viewAgenda)

	viewApproved := &[]model.Approves{}
	database.DBConn.Debug().Raw("SELECT * FROM approves WHERE item_number = ?", itemNumber).Scan(viewApproved)

	viewReleasings := &[]model.Releasings{}
	database.DBConn.Debug().Raw("SELECT * FROM releasings WHERE doc_id = ?", docId).Scan(viewReleasings)

	viewFilings := &[]model.Filings{}
	database.DBConn.Debug().Raw("SELECT * FROM filings WHERE doc_id = ?", docId).Scan(viewFilings)

	viewBorrowerHistory := &[]model.ViewBrrowerHistory{}
	database.DBConn.Debug().Raw("SELECT * FROM view_borrower_history").Scan(viewBorrowerHistory)

	folders := &[]model.Folders{}
	database.DBConn.Debug().Raw("SELECT * FROM folders").Scan(folders)

	return c.Render("filingupdate", fiber.Map{
		"pageTitle":           "for Update Filing",
		"title":               "FOR UPDATE FILING",
		"yearNow":             model.YearNow,
		"user":                model.Fullname,
		"userLogged":          model.UserCodeLogged,
		"viewRoutings":        viewRoutings,
		"viewAgenda":          viewAgenda,
		"viewApproved":        viewApproved,
		"viewReleasings":      viewReleasings,
		"viewFilings":         viewFilings,
		"viewBorrowerHistory": viewBorrowerHistory,
		"itemCommittees":      ItemCommittees,
		"docId":               docId,
		"itemNumber":          itemNumber,
		"folder":              folders,
		"greetings":           utils.GetGreetings(),
		"baseURL":             c.BaseURL(),
	})
}

func SaveFiling(c *fiber.Ctx) error {
	dateNow := time.Now()
	fmt.Println("Process: Save Filing")
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	requestFiling := &model.RequestFiling{}
	if parsErr := c.BodyParser(requestFiling); parsErr != nil {
		c.JSON(model.ResponseBody{
			Status:  101,
			Message: "error parsing request",
			Data:    parsErr.Error(),
		})
	}

	if requestFiling.FolderName == "new" {
		requestFiling.FolderName = requestFiling.NewFolderName
		database.DBConn.Debug().Exec("INSERT INTO folders (name) VALUES (?)", requestFiling.NewFolderName)
	}

	if !requestFiling.IsBorrowed {
		requestFiling.Borrower = ""
		requestFiling.DateBorrowed = ""
	}
	filingFields := &model.Filings{
		DocId:         requestFiling.DocId,
		CabinetNumber: requestFiling.CabinetNumber,
		FolderName:    requestFiling.FolderName,
		IsBorrowed:    requestFiling.IsBorrowed,
		DateBorrowed:  requestFiling.DateBorrowed,
		Borrower:      requestFiling.Borrower,
		DateReturned:  requestFiling.DateReturned,
		DateFiled:     requestFiling.DateFiled,
		DatePublished: requestFiling.DatePublished,
		Publisher:     requestFiling.Publisher,
		ModifiedBy:    model.Fullname,
		CreatedAt:     dateNow,
		UpdatedAt:     dateNow,
	}

	database.DBConn.Debug().Exec("UPDATE filings SET cabinet_number = ?, folder_name = ?, date_filed = ?, is_borrowed = ?, date_borrowed = ?, borrower = ?, date_returned = ?,date_published = ?, publisher = ?, modified_by = ?, updated_at = ? WHERE doc_id = ?",
		filingFields.CabinetNumber, filingFields.FolderName, filingFields.DateFiled,
		filingFields.IsBorrowed, filingFields.DateBorrowed, filingFields.Borrower, filingFields.DateReturned,
		filingFields.DatePublished, filingFields.Publisher, filingFields.ModifiedBy, dateNow, filingFields.DocId)

	database.DBConn.Debug().Exec("UPDATE trackings SET filed_date = ?, published_date = ?, updated_at = ? WHERE doc_id = ?", filingFields.DateFiled, filingFields.DatePublished, dateNow, filingFields.DocId)

	var borrowerId int
	database.DBConn.Debug().Raw("SELECT borrower_id FROM borrower_histories WHERE borrower = ?", filingFields.Borrower).Scan(&borrowerId)
	fmt.Println("Borrowed ID:", borrowerId)

	if filingFields.IsBorrowed {
		if borrowerId == 0 {
			borrowed := fmt.Sprintf("Borrowed - %v", filingFields.Borrower)
			database.DBConn.Debug().Exec("UPDATE routings SET remarks = ?, updated_at = ? WHERE doc_id = ?", borrowed, dateNow, filingFields.DocId)
			database.DBConn.Debug().Exec("INSERT INTO borrower_histories (doc_id, borrower, date_borrowed) VALUES(?,?,?)", filingFields.DocId, filingFields.Borrower, filingFields.DateBorrowed)
		}
	} else {
		database.DBConn.Debug().Exec("UPDATE routings SET remarks = 'Kept in Records', updated_at = ? WHERE doc_id = ?", dateNow, filingFields.DocId)
		database.DBConn.Debug().Exec("UPDATE borrower_histories SET date_returned = ?, updated_at = ? WHERE doc_id = ? AND borrower = ?", filingFields.DateReturned, dateNow, filingFields.DocId, filingFields.Borrower)
	}

	return c.Redirect("/api/routing")
}
