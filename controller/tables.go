package controller

import (
	"fmt"
	"tech_tubbies/middleware/database"
	"tech_tubbies/model"
)

func InitializeTables() {
	fmt.Println("Initializing database tables")
	database.CreateTable(model.Routings{})
	fmt.Println("Routings done")
	database.CreateTable(model.Document{})
	fmt.Println("Documents done")
	database.CreateTable(model.Receiving{})
	fmt.Println("Receiving done")
	database.CreateTable(model.ForAgenda{})
	fmt.Println("For Agenda done")
	database.CreateTable(model.ResOrd{})
	fmt.Println("ResOrd done")
	database.CreateTable(model.Forward{})
	fmt.Println("Forward done")
	database.CreateTable(model.Releasing{})
	fmt.Println("Releasing done")
	database.CreateTable(model.Filing{})
	fmt.Println("Filing done")
	database.CreateTable(model.DocumentFilepath{})
	fmt.Println("Document Filepath done")
	database.CreateTable(model.UserCredentials{})
	fmt.Println("User Credentials done")
	database.CreateTable(model.Divisions{})
	fmt.Println("Divisions done")
	database.CreateTable(model.Proponents{})
	fmt.Println("Proponents done")
	database.CreateTable(model.Committees{})
	fmt.Println("Committee done")
	database.CreateTable(model.Departments{})
	fmt.Println("Departments done")

	// Insert Default Divisions
	var divCount int
	database.DBConn.Raw("SELECT COUNT(*) FROM divisions").Scan(&divCount)

	if divCount <= 0 {
		database.DBConn.Debug().Exec("INSERT INTO divisions (name, code) VALUES ('Records','SPCRD'), ('Secretariat','SPCSD')")
	}

}
