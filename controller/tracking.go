package controller

import (
	"fmt"
	"tech_tubbies/global"
	"tech_tubbies/middleware/database"
	utils "tech_tubbies/middleware/util"
	"tech_tubbies/model"

	"github.com/gofiber/fiber/v2"
)

func ViewTracking(c *fiber.Ctx) error {
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
		"user":       global.Fullname,
		"userLogged": global.UserCodeLogged,
		"greetings":  utils.GetGreetings(),
		"proponents": proponents,
		"tracking":   tracking,
		"status":     101,
	})
}

// TRACKING
func SearchTracking(c *fiber.Ctx) error {
	fmt.Println("Process: Search Tracking")
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

	requestTracking := &model.RequestTracking{}
	if parsErr := c.BodyParser(requestTracking); parsErr != nil {
		return c.JSON(model.ResponseBody{
			Status:  101,
			Message: "error parsing fields",
			Data:    parsErr.Error(),
		})
	}

	proponents := &[]model.Proponents{}
	database.DBConn.Debug().Raw("SELECT DISTINCT * FROM proponents ORDER BY proponent_id ASC").Scan(proponents)

	responseTrackings := &[]model.Trackings{}
	if requestTracking.MotionedBy != "" {
		database.DBConn.Debug().Raw("SELECT * FROM trackings WHERE motioned_by = ?", requestTracking.MotionedBy).Scan(responseTrackings)
	} else if requestTracking.Author != "" {
		database.DBConn.Debug().Raw("SELECT * FROM trackings WHERE author = ?", requestTracking.Author).Scan(responseTrackings)
	} else if requestTracking.MotionedBy != "" && requestTracking.Author != "" {
		database.DBConn.Debug().Raw("SELECT * FROM trackings WHERE motioned_by AND author = ?", requestTracking.MotionedBy, requestTracking.Author).Scan(responseTrackings)
	} else {
		database.DBConn.Debug().Raw("SELECT * FROM trackings").Scan(responseTrackings)
	}

	return c.Render("tracking", fiber.Map{
		"pageTitle":  "Tracking",
		"title":      "TRACKING",
		"user":       global.Fullname,
		"userLogged": global.UserCodeLogged,
		"greetings":  utils.GetGreetings(),
		"proponents": proponents,
		"tracking":   responseTrackings,
		"status":     100,
	})
}

func DownloadAttachment(c *fiber.Ctx) error {
	filename := c.Params("filename")

	recordsCaptured := filename
	database.DBConn.Debug().Exec("INSERT INTO employee_performaces (records_retrieved,user_id) VALUES (?,?)", recordsCaptured, global.UserID)

	event := fmt.Sprintf("downloaded the file %s", filename)
	activity := "file downloaded"
	database.DBConn.Debug().Exec("INSERT INTO activity_loggers (activity, event, user_id) VALUES(?,?,?)", activity, event, global.UserID)

	return c.Download("./assets/uploads/" + filename)
}
