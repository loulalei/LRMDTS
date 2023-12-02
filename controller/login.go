package controller

import (
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
	model.Fullname = ""
	return c.Render("login", fiber.Map{
		"pageTitle": "Login",
		"title":     "LOGIN",
	})
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
		if userCredentials.DivisionCode == "SPCRD" { //Records
			return c.Redirect("/api/dashboard")
		} else if userCredentials.DivisionCode == "SPCSD" { //Secretariat
			return c.Redirect("/api/dashboard/secretariat")
		}
	}

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
