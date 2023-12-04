package controller

import (
	"fmt"
	"tech_tubbies/middleware/database"
	"tech_tubbies/model"
)

func InitializeTables() {
	fmt.Println("Initializing database tables")
	database.CreateTable(model.UserCredentials{})
	fmt.Println("User Credentials ✓")
	database.CreateTable(model.Receivings{})
	fmt.Println("Receivings ✓")
	database.CreateTable(model.Agendas{})
	fmt.Println("Agendas ✓")
	database.CreateTable(model.Divisions{})
	fmt.Println("Divisions ✓")
	database.CreateTable(model.Proponents{})
	fmt.Println("Proponents ✓")
	database.CreateTable(model.Committees{})
	fmt.Println("Committee ✓")
	database.CreateTable(model.Departments{})
	fmt.Println("Departments ✓")

	// Insert Default Divisions
	var count int
	database.DBConn.Raw("SELECT COUNT(*) FROM divisions").Scan(&count)
	if count == 0 {
		database.DBConn.Exec("INSERT INTO divisions (name, code) VALUES ('Records','SPCRD'), ('Secretariat','SPCSD')")
		fmt.Println("Success")
	}

	// Insert Default Divisions
	database.DBConn.Raw("SELECT COUNT(*) FROM committees").Scan(&count)
	if count == 0 {
		database.DBConn.Exec("INSERT INTO committees (name) VALUES ('Agriculture And Economic Productivity'), ('Anti-Drug Abuse'),('Barangay Affairs'), ('Basic Education'),('Blue Ribbon Committee'), ('Civil Society Organizations and Cooperatives'),('Disaster Risk Reduction and Management'), ('Energy, Transportation & Telecommunication'),('Environmental Protection and Solid Waste Management'), ('Ethics & Discipline'),('Finance and Appropriations'), ('Games, Amusements and Illegal Gambling'),('Government Contracts, Legal Matters, Engineering, Public Works and Zonification'), ('Health and Sanitation'),('Higher Education, Technical and Vocational Courses'), ('History, Arts and Culture'),('Housing, Land Use and Estate Development'), ('Human Resources, Good Governance, Public Ethics and Accountability'),('Human Rights and Public Information'), ('Information Technology, e-Commerce and Mass Media'),('International Affairs'), ('Labor and Employment'),('Markets, Slaughterhouse and Government Economic Enterprise'), ('Rules and Privileges'),('Sports Development'), ('Tourism'),('Trade, Commerce, and Industry'), ('Urban Poor and Livelihood'),('Water Resources Management & Development'), ('Ways and Means'),('Welfare and Protection of Family, Women, Children, Senior Citizens, Persons with Disability and Gender Equality'), ('Youth')")
		fmt.Println("Success")
	}

	// Insert Default Divisions
	database.DBConn.Raw("SELECT COUNT(*) FROM departments").Scan(&count)
	if count == 0 {
		database.DBConn.Exec("INSERT INTO departments (name) VALUES ('City Accountant Office'), ('City Budget Office'),('City Administrator Office'), ('City Auditor Office'),('City Treasurer Office'), ('City Planning and Development Coordinator Office'),('City Agriculture Office'), ('City Assesor Office'),('City Business Permit & Licensing Office'), ('City Building Office'), ('City Environment & Natural Resource Office'),('City Health Office'), ('City Human Resource & Management Office'),('City Information Office'), ('City Population Office'),('City Prosecutor Office'), ('City Traffic Management Office'),('Civil Registrar Office'), ('Cooperative Office'),('Disaster Risk Reduction Management Office'), ('Engineering Office'),('General Services Office'), ('Legal Office'),('LEDIPO'), ('Municipal Trial Court'),('Pamantasan ng Lungsod ng San Pablo'), ('Public Employment Office'),('Regional Trial Court'), ('San Pablo City General Hospital'),('San Pablo City General Hospital-Dialysis Unit'), ('Sangguniang Panlungsod'), ('Sangguniang Panlungsod-Library'), ('Senior Citizen Affairs'), ('Social Welfare & Development Office'), ('Solid Waste Management Office'), ('Tourism Office'), ('Veterinarian Office')")
		fmt.Println("Success")
	}

}
