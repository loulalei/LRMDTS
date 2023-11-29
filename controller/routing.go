package controller

import (
	"fmt"
	"tech_tubbies/middleware/database"
	utils "tech_tubbies/middleware/util"
	"tech_tubbies/model"

	"github.com/gofiber/fiber/v2"
)

func ViewRouting(c *fiber.Ctx) error {
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	receiving := &[]model.Routings{}
	database.DBConn.Debug().Raw("SELECT * FROM routings").Scan(receiving)
	tracking := &[]model.Routings{}
	database.DBConn.Debug().Raw("SELECT * FROM routings").Scan(tracking)
	return c.Render("routing", fiber.Map{
		"pageTitle":  "Routing",
		"title":      "ROUTING MAIN",
		"yearNow":    model.YearNow,
		"user":       model.Fullname,
		"userLogged": model.UserCodeLogged,
		"greetings":  utils.GetGreetings(),
		"receivings": receiving,
		"tracking":   tracking,
	})
}

func ViewRoutingRecords(c *fiber.Ctx) error {
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
	})
}

func ViewRoutingSecretariat(c *fiber.Ctx) error {
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	receiving := &[]model.Routings{}
	database.DBConn.Debug().Raw("SELECT * FROM routings WHERE document_tag = 'For Agenda'").Scan(receiving)
	tracking := &[]model.Routings{}
	database.DBConn.Debug().Raw("SELECT * FROM routings").Scan(tracking)
	return c.Render("routingSecretariat", fiber.Map{
		"pageTitle":  "Secretariat",
		"title":      "SECRETARIAT DEPARTMENT",
		"yearNow":    model.YearNow,
		"user":       model.Fullname,
		"userLogged": model.UserCodeLogged,
		"greetings":  utils.GetGreetings(),
		"receivings": receiving,
		"tracking":   tracking,
	})
}

func ViewRoutingForAgenda(c *fiber.Ctx) error {
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	id := c.Params("id")

	receiving := &[]model.Routings{}
	database.DBConn.Debug().Raw("SELECT * FROM routings WHERE document_tag = 'For Agenda' AND doc_id = ?", id).Scan(receiving)

	departments := &[]model.Departments{}
	database.DBConn.Debug().Raw("SELECT * FROM departments").Scan(departments)

	proponents := &[]model.Proponents{}
	database.DBConn.Debug().Raw("SELECT * FROM proponents").Scan(proponents)

	committees := &[]model.Committees{}
	database.DBConn.Debug().Raw("SELECT * FROM committees").Scan(committees)

	viewCommittees := &[]model.ViewCommittees{}
	database.DBConn.Debug().Raw("SELECT * FROM view_committees").Scan(viewCommittees)
	return c.Render("routingForAgenda", fiber.Map{
		"pageTitle":      "For Agenda",
		"title":          "SECRETARIAT - FOR AGENDA",
		"yearNow":        model.YearNow,
		"user":           model.Fullname,
		"userLogged":     model.UserCodeLogged,
		"userId":         model.UserID,
		"receivings":     receiving,
		"departments":    departments,
		"proponents":     proponents,
		"committees":     committees,
		"viewCommittees": viewCommittees,
		"greetings":      utils.GetGreetings(),
	})
}

func ViewReceivingRoute(c *fiber.Ctx) error {
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	departments := &[]model.Departments{}

	database.DBConn.Debug().Raw("SELECT * FROM departments").Scan(departments)

	return c.Render("routingReceivingDocument", fiber.Map{
		"pageTitle":   "Receiving Document",
		"title":       "RECORDS DEPARTMENT",
		"yearNow":     model.YearNow,
		"user":        model.Fullname,
		"userLogged":  model.UserCodeLogged,
		"departments": departments,
		"greetings":   utils.GetGreetings(),
	})
}

func RegisterReceiving(c *fiber.Ctx) error {
	if model.Fullname == "" {
		return c.Redirect("/")
	}
	receivingData := &model.Routings{}
	if parsErr := c.BodyParser(receivingData); parsErr != nil {
		return c.JSON(fiber.Map{
			"error": parsErr,
		})
	}

	if receivingData.DocumentTag == "1" {
		receivingData.DocumentTag = "For Agenda"
	} else if receivingData.DocumentTag == "2" {
		receivingData.DocumentTag = "For Filing"
	}

	file, err := c.FormFile("filename")
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	fmt.Println("FILENAME:", file.Filename)
	c.SaveFile(file, fmt.Sprintf("./assets/uploads/%s", file.Filename))

	database.DBConn.Debug().Exec("INSERT INTO routings (tracking_number,item_number,receive_date,receive_time,receiver,sender,summary,document_tag,remarks,received_file) VALUES (?,?,?,?,?,?,?,?,?,?)", receivingData.TrackingNumber, receivingData.ItemNumber, receivingData.ReceiveDate, receivingData.ReceiveTime, receivingData.Receiver, receivingData.Sender, receivingData.Summary, receivingData.DocumentTag, receivingData.Remarks, file.Filename)

	receiving := &[]model.Routings{}
	database.DBConn.Debug().Raw("SELECT * FROM routings").Scan(receiving)

	return c.Render("routing", fiber.Map{
		"pageTitle":          "Routing",
		"title":              "ROUTING MAIN",
		"yearNow":            model.YearNow,
		"user":               model.Fullname,
		"userLogged":         model.UserCodeLogged,
		"greetings":          utils.GetGreetings(),
		"receivings":         receiving,
		"registrationStatus": 100,
	})
}

func GetForFiling(c *fiber.Ctx) error {
	id := c.Params("id")
	receiving := &[]model.Routings{}

	database.DBConn.Debug().Raw("SELECT * FROM routings WHERE doc_id=?", id).Scan(receiving)

	folders := &[]model.Folders{}
	database.DBConn.Debug().Raw("SELECT * FROM folders").Scan(folders)
	return c.Render("routingFiling", fiber.Map{
		"pageTitle":  "forFiling",
		"title":      "FOR FILING",
		"yearNow":    model.YearNow,
		"user":       model.Fullname,
		"userLogged": model.UserCodeLogged,
		"greetings":  utils.GetGreetings(),
		"receivings": receiving,
		"folders":    folders,
	})
}

func UpdateForFiling(c *fiber.Ctx) error {
	receiving := &model.Routings{}
	c.BodyParser(receiving)

	fmt.Println("Receiving:", receiving.Remarks)
	database.DBConn.Debug().Exec("UPDATE routings SET cabinet=?, folder=?, is_borrowed=?, date_borrowed=?, borrower=?, remarks=? WHERE doc_id=?", receiving.Cabinet, receiving.Folder, receiving.IsBorrowed, receiving.DateBorrowed, receiving.Borrower, receiving.Remarks, receiving.DocId)
	//------------------
	receivings := &[]model.Routings{}
	database.DBConn.Debug().Raw("SELECT * FROM routings").Scan(receivings)
	return c.Render("routing", fiber.Map{
		"pageTitle":  "Routing",
		"title":      "ROUTING MAIN",
		"yearNow":    model.YearNow,
		"user":       model.Fullname,
		"userLogged": model.UserCodeLogged,
		"greetings":  utils.GetGreetings(),
		"receivings": receivings,
	})
}

func ViewApproved(c *fiber.Ctx) error {
	return c.Render("routingApproved", fiber.Map{
		"pageTitle":  "Routing - Approved",
		"title":      "ROUTING APPROVED",
		"yearNow":    model.YearNow,
		"user":       model.Fullname,
		"userLogged": model.UserCodeLogged,
		"greetings":  utils.GetGreetings(),
	})
}

func InsertCommitteeForAgenda(c *fiber.Ctx) error {
	itemNo := c.Params("itemNo")
	committeeId := c.Params("committeeId")
	userId := c.Params("userId")
	database.DBConn.Debug().Exec("INSERT INTO committee_lists (item_number, committee_id, user_id) VALUES (?,?,?)", itemNo, committeeId, userId)
	return c.JSON(fiber.Map{
		"ItemNo":      itemNo,
		"CommitteeId": committeeId,
		"UserId":      userId,
	})
}

func RemoveCommitteeForAgenda(c *fiber.Ctx) error {
	itemNo := c.Params("itemNo")
	committeeId := c.Params("committeeId")
	database.DBConn.Debug().Exec("DELETE FROM committee_lists WHERE committee_id = ? AND item_number = ?", committeeId, itemNo)
	return c.JSON(fiber.Map{
		"message": "successs",
		"ItemNo":  itemNo,
	})
}
