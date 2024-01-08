package controller

import (
	"fmt"
	"strconv"
	"tech_tubbies/global"
	"tech_tubbies/middleware/database"
	utils "tech_tubbies/middleware/util"
	"tech_tubbies/model"

	"github.com/gofiber/fiber/v2"
)

func ViewLogin(c *fiber.Ctx) error {
	global.Fullname = ""
	global.UserCodeLogged = ""
	global.UserID = 0
	global.DivisionCode = ""

	divisions := &[]model.Divisions{}
	database.DBConn.Raw("SELECT * FROM divisions").Scan(divisions)

	return c.Render("login", fiber.Map{
		"pageTitle": "Login",
		"title":     "LOGIN",
		"yearNow":   global.YearNow,
		"divisions": divisions,
	})
}

func Logout(c *fiber.Ctx) error {
	// Get session from storage
	session, err := global.Store.Get(c)
	if err != nil {
		panic(err)
	}

	event := fmt.Sprintf("%s successfully logged out", global.Fullname)
	database.DBConn.Debug().Exec("INSERT INTO activity_loggers (activity, event,user_id) VALUES(?,?,?)", "logged out", event, global.UserID)
	global.Fullname = ""
	global.UserCodeLogged = ""
	global.UserID = 0
	global.DivisionCode = ""

	session.Destroy()
	return c.Redirect("/")
}

func ValidateUser(c *fiber.Ctx) error {
	// Get session from storage
	session, err := global.Store.Get(c)
	if err != nil {
		panic(err)
	}

	userCredentials := &model.UserCredentials{}
	if parsErr := c.BodyParser(userCredentials); parsErr != nil {
		return c.JSON(fiber.Map{
			"error": parsErr,
		})
	}

	encryptedPass, _ := utils.Encrypt(userCredentials.Password, utils.GetEnv("SECRET_KEY"))
	database.DBConn.Raw("SELECT * FROM user_credentials WHERE division_code = ? AND password = ?", userCredentials.DivisionCode, encryptedPass).Scan(userCredentials)

	fmt.Println("CREDENTIALS:", userCredentials)
	if userCredentials.Id > 0 {
		session.Set("userId", strconv.Itoa(userCredentials.Id))
		session.Save()

		event := fmt.Sprintf("%s successfull logged in", global.Fullname)

		if userCredentials.DivisionCode == "SPCRD" { //Records
			database.DBConn.Debug().Exec("INSERT INTO activity_loggers (activity, event, user_id) VALUES(?,?,?)", "logged in", event, userCredentials.Id)
			return c.Redirect("/api/dashboard")
		} else if userCredentials.DivisionCode == "SPCSD" { //Secretariat
			database.DBConn.Debug().Exec("INSERT INTO activity_loggers (activity, event, user_id) VALUES(?,?,?)", "logged in", event, userCredentials.Id)
			return c.Redirect("/api/dashboard/secretariat")
		} else if userCredentials.DivisionCode == "HOD" { //Secretariat
			database.DBConn.Debug().Exec("INSERT INTO activity_loggers (activity, event, user_id) VALUES(?,?,?)", "logged in", event, userCredentials.Id)
			return c.Redirect("/api/dashboard/head_office")
		} else {
			return c.JSON(model.ResponseBody{
				Status:  101,
				Message: "Division is not valid",
			})
		}
	} else {
		divisions := &[]model.Divisions{}
		database.DBConn.Debug().Raw("SELECT * FROM divisions").Scan(divisions)
		return c.Render("login", fiber.Map{
			"pageTitle": "Login",
			"title":     "LOGIN",
			"yearNow":   global.YearNow,
			"status":    101,
			"divisions": divisions,
		})
	}
}

func VerifyUser(c *fiber.Ctx) error {
	userCredentials := &model.UserCredentials{}
	if parsErr := c.BodyParser(userCredentials); parsErr != nil {
		return c.JSON(fiber.Map{
			"error": parsErr,
		})
	}

	encryptedPass, _ := utils.Encrypt(userCredentials.Password, utils.GetEnv("SECRET_KEY"))
	database.DBConn.Raw("SELECT * FROM user_credentials WHERE division_code = ? AND password = ?", userCredentials.DivisionCode, encryptedPass).Scan(userCredentials)
	if userCredentials.Id == 0 {
		global.Fullname = userCredentials.Fullname
		global.UserCodeLogged = userCredentials.DivisionCode
		global.UserID = userCredentials.Id
		global.DivisionCode = userCredentials.DivisionCode
		event := fmt.Sprintf("%s successfull logged in", global.Fullname)

		database.DBConn.Debug().Raw("SELECT * FROM user_activities WHERE user_id = ?", global.UserID).Scan(model.UserActivity)

		if userCredentials.DivisionCode == "SPCRD" { //Records
			if model.UserActivity.UserId == 0 {
				database.DBConn.Debug().Exec("INSERT INTO user_activities (user_id, activity, is_logged) VALUES(?,?,?)", global.UserID, "LOGGED IN", true)
			} else {
				database.DBConn.Debug().Exec("UPDATE user_activities SET activity = 'LOGGED IN', is_logged = true WHERE user_id = ?", global.UserID)
			}
			database.DBConn.Debug().Exec("UPDATE user_credentials SET is_reset = ? WHERE id = ?", false, userCredentials.Id)
			database.DBConn.Debug().Exec("INSERT INTO activity_loggers (activity, event, user_id) VALUES(?,?,?)", "logged in", event, global.UserID)
			return c.Redirect("/api/dashboard")
		} else if userCredentials.DivisionCode == "SPCSD" { //Secretariat
			if model.UserActivity.UserId == 0 {
				database.DBConn.Debug().Exec("INSERT INTO user_activities (user_id, activity, is_logged) VALUES(?,?,?)", global.UserID, "LOGGED IN", true)
			} else {
				database.DBConn.Debug().Exec("UPDATE user_activities SET activity = 'LOGGED IN', is_logged = true WHERE user_id = ?", global.UserID)
			}
			database.DBConn.Debug().Exec("UPDATE user_credentials SET is_reset = ? WHERE id = ?", false, userCredentials.Id)
			database.DBConn.Debug().Exec("INSERT INTO activity_loggers (activity, event,user_id) VALUES(?,?,?)", "logged in", event, global.UserID)
			return c.Redirect("/api/dashboard/secretariat")
		} else if userCredentials.DivisionCode == "HOD" { //Secretariat
			if model.UserActivity.UserId == 0 {
				database.DBConn.Debug().Exec("INSERT INTO user_activities (user_id, activity, is_logged) VALUES(?,?,?)", global.UserID, "LOGGED IN", true)
			} else {
				database.DBConn.Debug().Exec("UPDATE user_activities SET activity = 'LOGGED IN', is_logged = true WHERE user_id = ?", global.UserID)
			}
			database.DBConn.Debug().Exec("UPDATE user_credentials SET is_reset = ? WHERE id = ?", false, userCredentials.Id)
			database.DBConn.Debug().Exec("INSERT INTO activity_loggers (activity, event, user_id) VALUES(?,?,?)", "logged in", event, global.UserID)
			return c.Redirect("/api/dashboard/head_office")
		}
	}
	divisions := &[]model.Divisions{}
	database.DBConn.Debug().Raw("SELECT * FROM divisions").Scan(divisions)
	return c.Render("login", fiber.Map{
		"pageTitle":   "Login",
		"title":       "LOGIN",
		"yearNow":     global.YearNow,
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
		"yearNow":   global.YearNow,
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
			"yearNow":            global.YearNow,
			"registrationStatus": 102,
		})
	}

	if userCredentials.Password != c.FormValue("cpassword") {
		return c.Render("userManagement", fiber.Map{
			"pageTitle":          "Registration",
			"title":              "User Management",
			"yearNow":            global.YearNow,
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
	database.DBConn.Debug().Exec("INSERT INTO activity_loggers (activity, event, user_id) VALUES(?,?,?)", activity, event, global.UserID)

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
	database.DBConn.Debug().Exec("INSERT INTO activity_loggers (activity, event, user_id) VALUES(?,?,?)", activity, event, global.UserID)

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

func GetUserDetails(userId int) *model.UserCredentials {
	userDetails := &model.UserCredentials{}
	database.DBConn.Debug().Raw("SELECT * FROM view_users WHERE id = ?", userId).Scan(userDetails)
	return userDetails
}
