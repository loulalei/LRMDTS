package controller

import (
	"fmt"
	"tech_tubbies/middleware/database"
	utils "tech_tubbies/middleware/util"
	"tech_tubbies/model"

	"github.com/gofiber/fiber/v2"
)

func ViewLogin(c *fiber.Ctx) error {
	divisions := &[]model.Divisions{}
	database.DBConn.Raw("SELECT * FROM divisions").Scan(divisions)

	return c.Render("login", fiber.Map{
		"pageTitle": "Login",
		"title":     "LOGIN",
		"yearNow":   model.YearNow,
		"divisions": divisions,
	})
}

func Logout(c *fiber.Ctx) error {
	event := fmt.Sprintf("%s successfully logged out", model.Fullname)
	database.DBConn.Debug().Exec("INSERT INTO activity_loggers (activity, event,user_id) VALUES(?,?,?)", "logged out", event, model.UserID)
	model.UserID = 0
	model.Fullname = ""

	return c.Redirect("/")
}

func VerifyUser(c *fiber.Ctx) error {
	userCredentials := &model.UserCredentials{}
	if parsErr := c.BodyParser(userCredentials); parsErr != nil {
		return c.JSON(fiber.Map{
			"error": parsErr,
		})
	}

	encryptedPass, _ := utils.Encrypt(userCredentials.Password, utils.GetEnv("SECRET_KEY"))
	isExist := database.DBConn.Raw("SELECT * FROM user_credentials WHERE division_code = ? AND password = ?", userCredentials.DivisionCode, encryptedPass).Scan(userCredentials).RowsAffected
	if isExist > 0 {
		model.Fullname = userCredentials.Fullname
		model.UserCodeLogged = userCredentials.DivisionCode
		model.UserID = userCredentials.Id
		model.DivisionCode = userCredentials.DivisionCode
		event := fmt.Sprintf("%s successfull logged in", model.Fullname)
		if userCredentials.DivisionCode == "SPCRD" { //Records
			database.DBConn.Debug().Exec("UPDATE user_credentials SET is_reset = ? WHERE id = ?", false, userCredentials.Id)
			database.DBConn.Debug().Exec("INSERT INTO activity_loggers (activity, event, user_id) VALUES(?,?,?)", "logged in", event, model.UserID)
			return c.Redirect("/api/dashboard")
		} else if userCredentials.DivisionCode == "SPCSD" { //Secretariat
			database.DBConn.Debug().Exec("UPDATE user_credentials SET is_reset = ? WHERE id = ?", false, userCredentials.Id)
			database.DBConn.Debug().Exec("INSERT INTO activity_loggers (activity, event,user_id) VALUES(?,?,?)", "logged in", event, model.UserID)
			return c.Redirect("/api/dashboard/secretariat")
		} else if userCredentials.DivisionCode == "HOD" { //Secretariat
			database.DBConn.Debug().Exec("UPDATE user_credentials SET is_reset = ? WHERE id = ?", false, userCredentials.Id)
			database.DBConn.Debug().Exec("INSERT INTO activity_loggers (activity, event, user_id) VALUES(?,?,?)", "logged in", event, model.UserID)
			return c.Redirect("/api/dashboard/head_office")
		}
	}

	database.DBConn.Debug().Exec("INSERT INTO activity_loggers (activity, event) VALUES(?,?)", "logged in", "failed to log in")
	divisions := &[]model.Divisions{}
	database.DBConn.Debug().Raw("SELECT * FROM divisions").Scan(divisions)
	return c.Render("login", fiber.Map{
		"pageTitle":   "Login",
		"title":       "LOGIN",
		"yearNow":     model.YearNow,
		"loginStatus": 101,
		"divisions":   divisions,
	})
}

func ViewRegistration(c *fiber.Ctx) error {

	divisions := &[]model.Divisions{}
	database.DBConn.Raw("SELECT * FROM divisions").Scan(divisions)

	return c.Render("userManagement", fiber.Map{
		"pageTitle": "Registration",
		"title":     "User Management",
		"yearNow":   model.YearNow,
		"divisions": divisions,
	})
}

func RegisterUser(c *fiber.Ctx) error {
	userCredentials := &model.UserCredentials{}
	if parsErr := c.BodyParser(userCredentials); parsErr != nil {
		return c.JSON(parsErr.Error())
	}

	if userCredentials.Fullname == "" {
		return c.Render("userManagement", fiber.Map{
			"pageTitle":          "Registration",
			"title":              "User Management",
			"yearNow":            model.YearNow,
			"registrationStatus": 102,
		})
	}

	if userCredentials.Password != c.FormValue("cpassword") {
		return c.Render("userManagement", fiber.Map{
			"pageTitle":          "Registration",
			"title":              "User Management",
			"yearNow":            model.YearNow,
			"registrationStatus": 101,
		})
	}

	encryptPassword, encErr := utils.Encrypt(userCredentials.Password, utils.GetEnv("SECRET_KEY"))
	if encErr != nil {
		return c.JSON(fiber.Map{
			"error": encErr.Error(),
		})
	}

	userCredentials.Password = encryptPassword

	database.DBConn.Debug().Raw("INSERT INTO user_credentials (fullname,password,division_code) VALUES (?,?,?)", userCredentials.Fullname, userCredentials.Password, userCredentials.DivisionCode).Find(userCredentials)

	return ViewLogin(c)
}

func AddUser(c *fiber.Ctx) error {
	userCredentials := &model.UserCredentials{}
	if parsErr := c.BodyParser(userCredentials); parsErr != nil {
		return c.JSON(parsErr.Error())
	}

	encryptPassword, encErr := utils.Encrypt(userCredentials.Password, utils.GetEnv("SECRET_KEY"))
	if encErr != nil {
		return c.JSON(fiber.Map{
			"error": encErr.Error(),
		})
	}

	userCredentials.Password = encryptPassword

	database.DBConn.Debug().Raw("INSERT INTO user_credentials (fullname,password,division_code) VALUES (?,?,?)", userCredentials.Fullname, userCredentials.Password, userCredentials.DivisionCode).Find(userCredentials)

	activity := "user management"
	event := fmt.Sprintf("added new user %s from %s", userCredentials.Fullname, userCredentials.DivisionCode)
	database.DBConn.Debug().Exec("INSERT INTO activity_loggers (activity, event, user_id) VALUES(?,?,?)", activity, event, model.UserID)

	return c.JSON(model.ResponseBody{
		Status:  100,
		Message: "Successful Registration",
	})
}

func ResetPasssword(c *fiber.Ctx) error {
	userCredentials := &model.UserCredentials{}
	if parsErr := c.BodyParser(userCredentials); parsErr != nil {
		return c.JSON(parsErr.Error())
	}

	passCode := utils.GeneratePasscode()

	encryptPassword, encErr := utils.Encrypt(passCode, utils.GetEnv("SECRET_KEY"))
	if encErr != nil {
		return c.JSON(fiber.Map{
			"error": encErr.Error(),
		})
	}

	database.DBConn.Debug().Exec("UPDATE user_credentials SET password = ?, is_reset = ? WHERE id = ?", encryptPassword, true, userCredentials.Id)

	activity := "user management"
	event := fmt.Sprintf("password of user %s from %s has been reset", userCredentials.Fullname, userCredentials.DivisionCode)
	database.DBConn.Debug().Exec("INSERT INTO activity_loggers (activity, event, user_id) VALUES(?,?,?)", activity, event, model.UserID)

	return c.JSON(model.ResponseBody{
		Status:  100,
		Message: "Successful password reset, please take a screenshot of this passcode to login.",
		Data:    passCode,
	})
}

func ChangePasssword(c *fiber.Ctx) error {
	userCredentials := &model.UserCredentials{}
	if parsErr := c.BodyParser(userCredentials); parsErr != nil {
		return c.JSON(parsErr.Error())
	}

	encryptPassword, encErr := utils.Encrypt(userCredentials.Password, utils.GetEnv("SECRET_KEY"))
	if encErr != nil {
		return c.JSON(fiber.Map{
			"error": encErr.Error(),
		})
	}

	database.DBConn.Debug().Exec("UPDATE user_credentials SET password = ?, is_reset = ? WHERE id = ?", encryptPassword, false, userCredentials.Id)

	return c.JSON(model.ResponseBody{
		Status:  100,
		Message: "Successful password update.",
	})
}
