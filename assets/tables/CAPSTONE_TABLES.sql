-- un:Rec_Admin
-- pw:Letmein.Recadmin
CREATE TABLE committee_lists(
	list_id serial primary key,
	item_number text,
	committee_id bigint,
	user_id bigint,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)

-- ROUTINGS
CREATE TABLE receivings(
	receiving_id serial primary key,
	tracking_number text,
	received_date text,
	received_time text,
	receiver text,
	sender text,
	summary text,
	receiving_tag text,
	receiving_remarks text,
	received_file text,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)

CREATE TABLE agendas(
	agenda_id serial primary key,
	item_number text,
	isUrgent bool,
	date_calendared text,
	date_reported text,
	source text,
	source_result text,
	agenda_tag text,
	agenda_remarks text,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)

CREATE TABLE routings(
	doc_id serial primary key,
	receiving_id int,
	agenda_id int,
	filing_id int,
	releasing_id int,
	item_number text,
	document_tag text,
	remarks text,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)

-- VIEW
CREATE VIEW view_committees AS
SELECT list_id, item_number, name, fullname 
FROM committee_lists cmtl
INNER JOIN committees cmt
ON cmt.id = cmtl.committee_id
INNER JOIN user_credentials usr
ON usr.id = cmtl.user_id

CREATE VIEW view_routings AS
SELECT route.doc_id, route.item_number, route.document_tag, route.remarks,
	rcv.receiving_id, rcv.tracking_number, rcv.received_date, rcv.received_time, rcv.receiver, rcv.sender, rcv.summary, rcv.received_file
FROM routings route
INNER JOIN receivings rcv
ON route.receiving_id = rcv.receiving_id
ORDER BY route.updated_at DESC

CREATE VIEW view_agendas AS
SELECT agd.agenda_id, agd.item_number, agd.isUrgent, agd.date_calendared, agd.date_reported, agd.source, agd.source_result, agd.agenda_remarks,
 vcmt."name"
FROM agendas agd
INNER JOIN view_committees vcmt
ON agd.item_number = vcmt.item_number
ORDER BY vcmt.item_number DESC

SELECT cmtl.list_id, cmtl.item_number FROM committee_lists cmtl
INNER JOIN committees cmt ON cmt.id = cmtl.committee_id

SELECT * FROM divisions
SELECT * FROM receivings
SELECT * FROM user_credentials
SELECT * FROM committees
SELECT * FROM departments
SELECT * FROM committee_lists
SELECT * FROM routings
SELECT * FROM view_committees
SELECT * FROM view_routings
SELECT * FROM view_agendas
SELECT * FROM agendas
SELECT COUNT(*) FROM divisions

SELECT * FROM view_routings WHERE document_tag = 'For Agenda' OR document_tag = 'Kept in Secretariat'
SELECT * FROM view_routings WHERE document_tag = 'Forwarded to Secretariat'
INSERT INTO divisions (name, code) VALUES ('Records','SPCRD'), ('Secretariat','SPCSD'), ('Administration','SPCAD')

UPDATE routings SET document_tag = 'Referred to Committee' WHERE doc_id = 3
INSERT INTO committees (name) VALUES ('Agriculture And Economic Productivity'), ('Anti-Drug Abuse'),('Barangay Affairs'), ('Basic Education'),('Blue Ribbon Committee'), ('Civil Society Organizations and Cooperatives'),('Disaster Risk Reduction and Management'), ('Energy, Transportation & Telecommunication'),('Environmental Protection and Solid Waste Management'), ('Ethics & Discipline'),('Finance and Appropriations'), ('Games, Amusements and Illegal Gambling'),('Government Contracts, Legal Matters, Engineering, Public Works and Zonification'), ('Health and Sanitation'),('Higher Education, Technical and Vocational Courses'), ('History, Arts and Culture'),('Housing, Land Use and Estate Development'), ('Human Resources, Good Governance, Public Ethics and Accountability'),('Human Rights and Public Information'), ('Information Technology, e-Commerce and Mass Media'),('International Affairs'), ('Labor and Employment'),('Markets, Slaughterhouse and Government Economic Enterprise'), ('Rules and Privileges'),('Sports Development'), ('Tourism'),('Trade, Commerce, and Industry'), ('Urban Poor and Livelihood'),('Water Resources Management & Development'), ('Ways and Means'),('Welfare and Protection of Family, Women, Children, Senior Citizens, Persons with Disability and Gender Equality'), ('Youth')

SELECT COUNT(agenda_id) FROM routings WHERE agenda_id IS NULL
