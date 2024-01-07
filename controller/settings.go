package controller

import (
	"fmt"
	"tech_tubbies/global"
	"tech_tubbies/middleware/database"
	utils "tech_tubbies/middleware/util"
	"tech_tubbies/model"

	"github.com/gofiber/fiber/v2"
)

func ViewSettings(c *fiber.Ctx) error {
	fmt.Println("Process: View Settings")
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

	return c.Render("settings", fiber.Map{
		"pageTitle":  "Settings",
		"title":      "SYSTEM SETTINGS",
		"yearNow":    global.YearNow,
		"user":       global.Fullname,
		"userLogged": global.UserCodeLogged,
		"userId":     global.UserID,
		"greetings":  utils.GetGreetings(),
		"baseURL":    c.BaseURL(),
	})
}

func ViewSettingsUsers(c *fiber.Ctx) error {
	fmt.Println("Process: Users Settings")
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

	userList := &[]model.ViewUsers{}
	database.DBConn.Debug().Raw("SELECT * FROM view_users WHERE id NOT IN (?)", global.UserID).Scan(userList)

	divisions := &[]model.Divisions{}
	database.DBConn.Raw("SELECT * FROM divisions").Scan(divisions)

	return c.Render("settingsusers", fiber.Map{
		"pageTitle":  "Settings",
		"title":      "USERS SETTINGS",
		"yearNow":    global.YearNow,
		"user":       global.Fullname,
		"userLogged": global.UserCodeLogged,
		"userId":     global.UserID,
		"userList":   userList,
		"divisions":  divisions,
		"greetings":  utils.GetGreetings(),
		"baseURL":    c.BaseURL(),
	})
}

func ViewSettingsProponents(c *fiber.Ctx) error {
	fmt.Println("Process: Proponents Settings")
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

	proponentList := &[]model.Proponents{}
	database.DBConn.Debug().Raw("SELECT * FROM proponents").Scan(proponentList)

	return c.Render("settingsproponents", fiber.Map{
		"pageTitle":     "Settings",
		"title":         "PROPONENTS SETTINGS",
		"yearNow":       global.YearNow,
		"user":          global.Fullname,
		"userLogged":    global.UserCodeLogged,
		"userId":        global.UserID,
		"proponentList": proponentList,
		"greetings":     utils.GetGreetings(),
		"baseURL":       c.BaseURL(),
	})
}

func ViewSettingsCommittees(c *fiber.Ctx) error {
	fmt.Println("Process: Committee Settings")
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

	committeeList := &[]model.Committees{}
	database.DBConn.Debug().Raw("SELECT * FROM committees").Scan(committeeList)

	return c.Render("settingscommittees", fiber.Map{
		"pageTitle":     "Settings",
		"title":         "COMMITTEE SETTINGS",
		"yearNow":       global.YearNow,
		"user":          global.Fullname,
		"userLogged":    global.UserCodeLogged,
		"userId":        global.UserID,
		"committeeList": committeeList,
		"greetings":     utils.GetGreetings(),
		"baseURL":       c.BaseURL(),
	})
}

func ViewSettingsFolder(c *fiber.Ctx) error {
	fmt.Println("Process: Folder Settings")
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

	return c.Render("settings", fiber.Map{
		"pageTitle":  "Settings",
		"title":      "FOLDER SETTINGS",
		"yearNow":    global.YearNow,
		"user":       global.Fullname,
		"userLogged": global.UserCodeLogged,
		"userId":     global.UserID,
		"greetings":  utils.GetGreetings(),
		"baseURL":    c.BaseURL(),
	})
}

func AddProponents(c *fiber.Ctx) error {
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

	proponentFields := &model.Proponents{}
	if parsErr := c.BodyParser(proponentFields); parsErr != nil {
		return c.JSON(model.ResponseBody{
			Status:  101,
			Message: "Error parsing:" + parsErr.Error(),
		})
	}

	database.DBConn.Debug().Exec("INSERT INTO proponents (name, term) VALUES (?,?)", proponentFields.Name, proponentFields.Term)

	activity := "proponents management"
	event := fmt.Sprintf("%s has been added", proponentFields.Name)
	database.DBConn.Debug().Exec("INSERT INTO activity_loggers (activity, event, user_id) VALUES(?,?,?)", activity, event, global.UserID)

	return c.JSON(model.ResponseBody{
		Status:  100,
		Message: "Added successfuly",
	})
}

func DeleteProponent(c *fiber.Ctx) error {
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

	proponentFields := &model.Proponents{}
	if parsErr := c.BodyParser(proponentFields); parsErr != nil {
		return c.JSON(model.ResponseBody{
			Status:  101,
			Message: "Error parsing:" + parsErr.Error(),
		})
	}

	database.DBConn.Debug().Exec("DELETE FROM proponents WHERE name = ?", proponentFields.Name)

	activity := "proponents management"
	event := fmt.Sprintf("%s has been removed", proponentFields.Name)
	database.DBConn.Debug().Exec("INSERT INTO activity_loggers (activity, event, user_id) VALUES(?,?,?)", activity, event, global.UserID)

	return c.JSON(model.ResponseBody{
		Status:  100,
		Message: "Remove successfuly",
	})
}

func AddCommittee(c *fiber.Ctx) error {
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

	committeeFields := &model.Committees{}
	if parsErr := c.BodyParser(committeeFields); parsErr != nil {
		return c.JSON(model.ResponseBody{
			Status:  101,
			Message: "Error parsing:" + parsErr.Error(),
		})
	}

	var name string
	database.DBConn.Debug().Raw("SELECT name FROM committees WHERE name = ?", committeeFields.Name).Scan(&name)

	if name == "" {
		database.DBConn.Debug().Exec("INSERT INTO committees (name) VALUES (?)", committeeFields.Name)

		activity := "committees management"
		event := fmt.Sprintf("%s has been added", committeeFields.Name)
		database.DBConn.Debug().Exec("INSERT INTO activity_loggers (activity, event, user_id) VALUES(?,?,?)", activity, event, global.UserID)
	} else {
		return c.JSON(model.ResponseBody{
			Status:  101,
			Message: "Committee already exists",
		})
	}

	return c.JSON(model.ResponseBody{
		Status:  100,
		Message: "Added successfuly",
	})
}

func DeleteCommittee(c *fiber.Ctx) error {
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

	committeeFields := &model.Committees{}
	if parsErr := c.BodyParser(committeeFields); parsErr != nil {
		return c.JSON(model.ResponseBody{
			Status:  101,
			Message: "Error parsing:" + parsErr.Error(),
		})
	}

	database.DBConn.Debug().Exec("DELETE FROM committees WHERE name = ?", committeeFields.Name)

	activity := "committees management"
	event := fmt.Sprintf("%s has been removed", committeeFields.Name)
	database.DBConn.Debug().Exec("INSERT INTO activity_loggers (activity, event, user_id) VALUES(?,?,?)", activity, event, global.UserID)

	return c.JSON(model.ResponseBody{
		Status:  100,
		Message: "Added successfuly",
	})
}
