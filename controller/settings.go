package controller

import (
	"fmt"
	"tech_tubbies/middleware/database"
	utils "tech_tubbies/middleware/util"
	"tech_tubbies/model"

	"github.com/gofiber/fiber/v2"
)

func ViewSettings(c *fiber.Ctx) error {
	fmt.Println("Process: View Settings")
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	return c.Render("settings", fiber.Map{
		"pageTitle":  "Settings",
		"title":      "SYSTEM SETTINGS",
		"yearNow":    model.YearNow,
		"user":       model.Fullname,
		"userLogged": model.UserCodeLogged,
		"userId":     model.UserID,
		"greetings":  utils.GetGreetings(),
		"baseURL":    c.BaseURL(),
	})
}

func ViewSettingsUsers(c *fiber.Ctx) error {
	fmt.Println("Process: Users Settings")
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	userList := &[]model.ViewUsers{}
	database.DBConn.Debug().Raw("SELECT * FROM view_users").Scan(userList)

	divisions := &[]model.Divisions{}
	database.DBConn.Raw("SELECT * FROM divisions").Scan(divisions)

	return c.Render("settingsusers", fiber.Map{
		"pageTitle":  "Settings",
		"title":      "USERS SETTINGS",
		"yearNow":    model.YearNow,
		"user":       model.Fullname,
		"userLogged": model.UserCodeLogged,
		"userId":     model.UserID,
		"userList":   userList,
		"divisions":  divisions,
		"greetings":  utils.GetGreetings(),
		"baseURL":    c.BaseURL(),
	})
}

func ViewSettingsProponents(c *fiber.Ctx) error {
	fmt.Println("Process: Proponents Settings")
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	proponentList := &[]model.Proponents{}
	database.DBConn.Debug().Raw("SELECT * FROM proponents").Scan(proponentList)

	return c.Render("settingsproponents", fiber.Map{
		"pageTitle":     "Settings",
		"title":         "PROPONENTS SETTINGS",
		"yearNow":       model.YearNow,
		"user":          model.Fullname,
		"userLogged":    model.UserCodeLogged,
		"userId":        model.UserID,
		"proponentList": proponentList,
		"greetings":     utils.GetGreetings(),
		"baseURL":       c.BaseURL(),
	})
}

func ViewSettingsCommittees(c *fiber.Ctx) error {
	fmt.Println("Process: Committee Settings")
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	committeeList := &[]model.Committees{}
	database.DBConn.Debug().Raw("SELECT * FROM committees").Scan(committeeList)

	return c.Render("settingscommittees", fiber.Map{
		"pageTitle":     "Settings",
		"title":         "COMMITTEE SETTINGS",
		"yearNow":       model.YearNow,
		"user":          model.Fullname,
		"userLogged":    model.UserCodeLogged,
		"userId":        model.UserID,
		"committeeList": committeeList,
		"greetings":     utils.GetGreetings(),
		"baseURL":       c.BaseURL(),
	})
}

func ViewSettingsFolder(c *fiber.Ctx) error {
	fmt.Println("Process: Folder Settings")
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	return c.Render("settings", fiber.Map{
		"pageTitle":  "Settings",
		"title":      "FOLDER SETTINGS",
		"yearNow":    model.YearNow,
		"user":       model.Fullname,
		"userLogged": model.UserCodeLogged,
		"userId":     model.UserID,
		"greetings":  utils.GetGreetings(),
		"baseURL":    c.BaseURL(),
	})
}

func AddProponents(c *fiber.Ctx) error {
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	proponentFields := &model.Proponents{}
	if parsErr := c.BodyParser(proponentFields); parsErr != nil {
		return c.JSON(model.ResponseBody{
			Status:  101,
			Message: "Error parsing:" + parsErr.Error(),
		})
	}

	database.DBConn.Debug().Exec("INSERT INTO proponents (name, term) VALUES (?,?)", proponentFields.Name, proponentFields.Term)

	return c.JSON(model.ResponseBody{
		Status:  100,
		Message: "Added successfuly",
	})
}

func DeleteProponent(c *fiber.Ctx) error {
	if model.Fullname == "" {
		return c.Redirect("/")
	}
	proponentFields := &model.Proponents{}
	if parsErr := c.BodyParser(proponentFields); parsErr != nil {
		return c.JSON(model.ResponseBody{
			Status:  101,
			Message: "Error parsing:" + parsErr.Error(),
		})
	}

	database.DBConn.Debug().Exec("DELETE FROM proponents WHERE name = ?", proponentFields.Name)

	return c.JSON(model.ResponseBody{
		Status:  100,
		Message: "Remove successfuly",
	})
}
func AddCommittee(c *fiber.Ctx) error {
	if model.Fullname == "" {
		return c.Redirect("/")
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
	if model.Fullname == "" {
		return c.Redirect("/")
	}

	committeeFields := &model.Committees{}
	if parsErr := c.BodyParser(committeeFields); parsErr != nil {
		return c.JSON(model.ResponseBody{
			Status:  101,
			Message: "Error parsing:" + parsErr.Error(),
		})
	}

	database.DBConn.Debug().Exec("DELETE FROM committees WHERE name = ?", committeeFields.Name)

	return c.JSON(model.ResponseBody{
		Status:  100,
		Message: "Added successfuly",
	})
}
