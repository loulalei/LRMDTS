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
	database.CreateTable(model.CommitteeList{})
	fmt.Println("Committee List done")

	// Insert Default Divisions
	var divCount int
	database.DBConn.Raw("SELECT COUNT(*) FROM divisions").Scan(&divCount)

	if divCount <= 0 {
		database.DBConn.Exec("INSERT INTO divisions (name, code) VALUES ('Records','SPCRD'), ('Secretariat','SPCSD')")
	}

	// Insert Default Divisions
	var committeeCount int
	database.DBConn.Raw("SELECT COUNT(*) FROM committees").Scan(&committeeCount)
	if committeeCount <= 0 {
		database.DBConn.Exec("INSERT INTO committees (name) VALUES ('Agriculture And Economic Productivity'), ('Anti-Drug Abuse'),('Barangay Affairs'), ('Basic Education'),('Blue Ribbon Committee'), ('Civil Society Organizations and Cooperatives'),('Disaster Risk Reduction and Management'), ('Energy, Transportation & Telecommunication'),('Environmental Protection and Solid Waste Management'), ('Ethics & Discipline'),('Finance and Appropriations'), ('Games, Amusements and Illegal Gambling'),('Government Contracts, Legal Matters, Engineering, Public Works and Zonification'), ('Health and Sanitation'),('Higher Education, Technical and Vocational Courses'), ('History, Arts and Culture'),('Housing, Land Use and Estate Development'), ('Human Resources, Good Governance, Public Ethics and Accountability'),('Human Rights and Public Information'), ('Information Technology, e-Commerce and Mass Media'),('International Affairs'), ('Labor and Employment'),('Markets, Slaughterhouse and Government Economic Enterprise'), ('Rules and Privileges'),('Sports Development'), ('Tourism'),('Trade, Commerce, and Industry'), ('Urban Poor and Livelihood'),('Water Resources Management & Development'), ('Ways and Means'),('Welfare and Protection of Family, Women, Children, Senior Citizens, Persons with Disability and Gender Equality'), ('Youth')")
		fmt.Println("Success")
	}

}
