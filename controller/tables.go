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
	database.CreateTable(model.Approves{})
	fmt.Println("Approved ✓")
	database.CreateTable(model.Releasings{})
	fmt.Println("Releasing Fields ✓")
	database.CreateTable(model.Filings{})
	fmt.Println("Filing ✓")
	database.CreateTable(model.Routings{})
	fmt.Println("Routing ✓")
	database.CreateTable(model.Trackings{})
	fmt.Println("Tracking ✓")
	database.CreateTable(model.CommitteeLists{})
	fmt.Println("Committee Lists ✓")
	database.CreateTable(model.FilePaths{})
	fmt.Println("Borrower Fields ✓")
	database.CreateTable(model.BorrowerHistories{})
	fmt.Println("File Path ✓")
	database.CreateTable(model.Divisions{})
	fmt.Println("Divisions ✓")
	database.CreateTable(model.Proponents{})
	fmt.Println("Proponents ✓")
	database.CreateTable(model.Committees{})
	fmt.Println("Committee ✓")
	database.CreateTable(model.Departments{})
	fmt.Println("Departments ✓")
	database.CreateTable(model.Folders{})
	fmt.Println("Folder Fields ✓")
	database.CreateTable(model.DashboardTotal{})
	fmt.Println("Dashboard Fields ✓")
	database.CreateTable(model.ActivityLogger{})
	fmt.Println("Activity Logger ✓")
	database.CreateTable(model.EmployeePerformace{})
	fmt.Println("Employee Performance Logger ✓")
	database.CreateTable(model.UserActivities{})
	fmt.Println("User Activities ✓")
	database.CreateTable(model.EventCalendar{})
	fmt.Println("Event Calendar ✓")

	// Insert Default Divisions
	var count int
	database.DBConn.Raw("SELECT COUNT(*) FROM divisions").Scan(&count)
	if count == 0 {
		database.DBConn.Exec("INSERT INTO divisions (name, code) VALUES ('Head Office','HOD'), ('Records','SPCRD'), ('Secretariat','SPCSD')")
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

	// Insert Default Folders
	database.DBConn.Raw("SELECT COUNT(*) FROM folders").Scan(&count)
	if count == 0 {
		database.DBConn.Exec("INSERT INTO folders (name) VALUES ('20 PERCENT DEV. FUND'),('2011 REVISED TRAFFIC CODE'),('2012 REVISED REVENUE CODE'),('AMUSEMENTS'),('ANIMALS'),('ANNUAL BUDGET'),('ANNUAL DEVELOPMENT PLAN'),('ANNUAL INVESTMENT PLAN (AIP)'),('ANTI-DRUG'),('ANTI-LITTERING'),('ANTI-RABIES'),('ANTI-TRAFFICKING ACT OF 2003'),('ARTS AND CULTURE'),('AWARDS'),('BANK LOANS'),('BARANGAY EMERGENCY RESPONSE TEAM (BERT)'),('BARANGAY MICRO BUSINESS'),('BARANGAY UNIFORM SCHEDULE OF FEES'),('BUDGET'),('BUSINESS IDENTIFICATION NUMBERS (BIN)'),('CALL OF NATURE'),('CASH INCENTIVES'),('CCTV (CLOSED CIRCUIT TV)'),('CEDERA'),('CEMETERY'),('CIGARETTES, TOBACCO PRODUCTS'),('CITY AGRICULTURAL AND FISHERY COUNCIL (CAFC)'),('CITY URBAN DEVELOPMENT AND HOUSING BOARD'),('CIVIL REGISTRAR'),('COCONUT'),('COMMISSION ON DECORUM AND INVESTIGATION'),('COMMON PARKING'),('COMPREHENSIVE TRAFFIC CODE'),('CONDONATION'),('CONSTRUCTIONS'),('CREATIONS'),('DEATH OR BURIAL'),('DEVELOPMENT PLAN'),('DISCOHOUSES'),('DLSP'),('DOMESTIC VIOLENCE'),('DRAG RACING'),('ELECTRICAL ENGINEERING LAW'),('ENVIRONMENT'),('FARMS'),('FUNERAL PARLOR'),('GENDER AND DEVELOPMENT PLAN (GAD)'),('GENERAL FUND'),('GENERIC CODE (SOLID WASTE MANAGEMENT)'),('HEALTH'),('HOLIDAY'),('HONORARIUM'),('HOUSING (HLURB)'),('HYMN'),('ILLEGAL LOGGING'),('INTERNET CAFÉ or COMPUTER GAME CENTERS'),('INVESTMENT INCENTIVES'),('LABOR'),('LAKES'),('LAND CLASSIFICATION'),('LAND CONVERSION'),('LOANS'),('LOCAL CODE ON CHILDREN'),('LOCAL TAX CODE'),('MARKET CODE'),('MORATORIUM'),('MOVIEHOUSES'),('NIGHT MARKET'),('OFFICES - SPC DISASTER RISK REDUCTION AND MANAGEMENT COUNCIL'),('ORGANIC AGRICULTURE ACT OF 2010'),('OVERSEAS FILIPINO AFFAIRS COUNCIL'),('PARENTAL RESPONSIBILITY CODE'),('PARKS'),('PEOPLES LAW ENFORCEMENT BOARD (PLEB)'),('PERSONS WITH DISABILITY (PWD)'),('PHILHEALTH TRUST FUND'),('PUBLIC EMPLOYMENT SERVICE OFFICE (PESO)'),('PUBLIC MARKET'),('PUBLIC TRANSPORTATION'),('QUASI JUDICIAL POWERS OF SP'),('REALIGNMENT'),('RECREATION'),('REVENUE CODE'),('REVERSION'),('SCHOOLS'),('SENIOR CITIZENS'),('SEPTAGE MANAGEMENT SYSTEM'),('SPC CDRRMC'),('SPC CITIZENS CHARTER'),('SPC CONSUMER WELFARE AND PROTECTION CENTER'),('SPC TRIPARTITE INDUSTRIAL PEACE COUNCIL (SPCTIPC)'),('SPECIAL DAYS'),('STREETS'),('SWAT'),('TAX EXEMPTION'),('THE 2012 FAMILY INTEGRITY WELFARE AND PROTECTION ACT OF THE CITY OF SAN PABLO'),('TIPPING FEE'),('TOURISM CODE'),('TOYS'),('TRAFFIC'),('TRANSFER OF PROJECT'),('TRICYCLES'),('VETOED CM'),('VIDEO'),('YOUTH'),('ZONING')")
		fmt.Println("Success")
	}

	// Create View Tables
	if database.DBErr = database.DBConn.Exec("CREATE VIEW view_committees AS SELECT list_id, item_number, name, fullname, cmt.committee_id FROM committee_lists cmtl INNER JOIN committees cmt ON cmt.committee_id = cmtl.committee_id INNER JOIN user_credentials usr ON usr.id = cmtl.user_id ORDER BY cmtl.list_id DESC").Error; database.DBErr != nil {
		fmt.Println("DB ERROR:", database.DBErr.Error())
	}

	if database.DBErr = database.DBConn.Exec("CREATE VIEW view_routings AS SELECT route.doc_id, route.item_number, route.document_tag, route.remarks, to_char(route.updated_at, 'Mon. DD,YYYY | HH12:MI AM') AS updated_at, rcv.receiving_id, route.releasing_id, rcv.tracking_number, rcv.received_date, rcv.received_time, rcv.receiver, rcv.summary, rcv.received_file, rcv.encoder, rcv.modified_by FROM routings route INNER JOIN receivings rcv ON route.receiving_id = rcv.receiving_id ORDER BY route.updated_at DESC").Error; database.DBErr != nil {
		fmt.Println("DB ERROR:", database.DBErr.Error())
	}

	if database.DBErr = database.DBConn.Exec("CREATE VIEW view_agendas AS SELECT agd.agenda_id, agd.item_number, agd.is_urgent, agd.date_calendared, agd.date_reported, agd.source, agd.source_result, agd.agenda_tag, agd.agenda_remarks, agd.encoder, agd.modified_by, vcmt.name FROM agendas agd INNER JOIN view_committees vcmt ON agd.item_number = vcmt.item_number ORDER BY vcmt.item_number DESC").Error; database.DBErr != nil {
		fmt.Println("DB ERROR:", database.DBErr.Error())
	}

	if database.DBErr = database.DBConn.Exec("CREATE VIEW view_borrower_history AS SELECT brw.borrower_id, apr.law_type, apr.res_ord_file, brw.borrower, brw.date_borrowed, brw.date_returned FROM borrower_histories brw INNER JOIN filings fil ON brw.doc_id = fil.doc_id INNER JOIN view_routings rtg ON rtg.doc_id = fil.doc_id INNER JOIN approves apr ON apr.item_number = rtg.item_number").Error; database.DBErr != nil {
		fmt.Println("DB ERROR:", database.DBErr.Error())
	}

	if database.DBErr = database.DBConn.Exec("CREATE VIEW view_users AS SELECT id, fullname, password, div.name, division_code, is_reset, to_char(created_at, 'Mon. DD,YYYY') AS date_registered  FROM user_credentials usr INNER JOIN divisions div ON div.code = usr.division_code").Error; database.DBErr != nil {
		fmt.Println("DB ERROR:", database.DBErr.Error())
	}

}
