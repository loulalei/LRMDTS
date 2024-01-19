package controller

import (
	"fmt"
	"tech_tubbies/global"
	"tech_tubbies/middleware/database"
	utils "tech_tubbies/middleware/util"
	"tech_tubbies/model"

	"github.com/gofiber/fiber/v2"
)

func ViewProfile(c *fiber.Ctx) error {
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

	receiving := &[]model.Routings{}
	database.DBConn.Raw("SELECT * FROM routings").Scan(receiving)

	activityLogs := &[]model.ViewActivityLogs{}
	EmployeePerformance := &[]model.EmployeePerformances{}
	if global.DivisionCode == "HOD" {
		database.DBConn.Debug().Raw("SELECT activity_id, activity, event, TO_CHAR(created_at, 'MM-DD-YYYY | HH12:MI AM') AS created_at FROM activity_loggers ORDER BY activity_id DESC").Scan(activityLogs)
		database.DBConn.Debug().Raw("SELECT DISTINCT usr.fullname, TO_CHAR(emp.created_at, 'Mon. DD,YYYY') AS date, COUNT(emp.records_captured) as records_captured, COUNT(emp.records_retrieved) as records_retrieved FROM employee_performaces emp INNER JOIN user_credentials usr ON usr.id = emp.user_id GROUP BY date, usr.fullname").Scan(EmployeePerformance)
	} else {
		database.DBConn.Debug().Raw("SELECT activity_id, activity, event, TO_CHAR(created_at, 'MM-DD-YYYY | HH12:MI AM') AS created_at FROM activity_loggers WHERE user_id = ? ORDER BY activity_id DESC", global.UserID).Scan(activityLogs)
		database.DBConn.Debug().Raw("SELECT DISTINCT TO_CHAR(created_at, 'Mon. DD,YYYY') AS date, COUNT(records_captured) as records_captured, COUNT(records_retrieved) as records_retrieved FROM employee_performaces WHERE user_id = ? GROUP BY date", global.UserID).Scan(EmployeePerformance)
	}

	fmt.Println("Division:", global.DivisionCode)

	return c.Render("profile", fiber.Map{
		"pageTitle":            "Profile",
		"title":                "PROFILE",
		"yearNow":              global.YearNow,
		"user":                 global.Fullname,
		"userLogged":           global.UserCodeLogged,
		"greetings":            utils.GetGreetings(),
		"logs":                 receiving,
		"activities":           activityLogs,
		"loggedDivision":       global.DivisionCode,
		"employeePerformances": EmployeePerformance,
		"baseURL":              c.BaseURL(),
	})
}
