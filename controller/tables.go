package controller

import (
	"fmt"
	"tech_tubbies/middleware/database"
	"tech_tubbies/model"
)

func InitializeTables() {
	fmt.Println("Initializing database tables")
	database.CreateTable(model.Document{})
	database.CreateTable(model.UserCredentials{})
	database.CreateTable(model.Divisions{})
	database.CreateTable(model.Proponents{})
	database.CreateTable(model.Committees{})
	database.CreateTable(model.Departments{})

	// Insert Default Divisions
	var divCount int
	database.DBConn.Raw("SELECT COUNT(*) FROM divisions").Scan(&divCount)
	if divCount <= 0 {
		database.DBConn.Debug().Exec("INSERT INTO divisions (name, code) VALUES ('Records','SPCRD'), ('Secretariat','SPCSD')")
	}

}
