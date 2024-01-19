-- un:Rec_Admin
-- pw:Letmein.Recadmin
-- VIEW
CREATE VIEW view_committees AS
SELECT list_id, item_number, name, fullname, cmt.committee_id
FROM committee_lists cmtl
INNER JOIN committees cmt
ON cmt.committee_id = cmtl.committee_id
INNER JOIN user_credentials usr
ON usr.id = cmtl.user_id
ORDER BY cmtl.list_id DESC

SELECT to_char(updated_at, 'Mon. DD,YYYY | HH12:MI AM') AS updated_at FROM routings

SELECT * FROM routings

CREATE VIEW view_routings AS
SELECT route.doc_id, route.item_number, route.document_tag, route.remarks, to_char(route.updated_at, 'Mon. DD,YYYY | HH12:MI AM') AS updated_at,
	rcv.receiving_id, route.releasing_id , rcv.tracking_number, rcv.received_date, rcv.received_time, rcv.receiver, rcv.summary, rcv.received_file, rcv.encoder, rcv.modified_by
FROM routings route
INNER JOIN receivings rcv
ON route.receiving_id = rcv.receiving_id
ORDER BY route.updated_at DESC

CREATE VIEW view_agendas AS
SELECT agd.agenda_id, agd.item_number, agd.is_urgent, agd.date_calendared, agd.date_reported, agd.source, agd.source_result, agd.agenda_tag, agd.agenda_remarks, agd.encoder, , rcv.modified_by,
 vcmt.name
FROM agendas agd
INNER JOIN view_committees vcmt
ON agd.item_number = vcmt.item_number
ORDER BY vcmt.item_number DESC

CREATE VIEW view_borrower_history AS
SELECT brw.borrower_id, apr.law_type, apr.res_ord_file, brw.borrower, brw.date_borrowed, brw.date_returned
FROM borrower_histories brw
INNER JOIN filings fil
ON brw.doc_id = fil.doc_id
INNER JOIN view_routings rtg
ON rtg.doc_id = fil.doc_id
INNER JOIN approves apr
ON apr.item_number = rtg.item_number

CREATE VIEW view_users AS
SELECT id, fullname, password, div.name, division_code, is_reset, to_char(created_at, 'Mon. DD,YYYY') AS date_registered 
FROM user_credentials usr
INNER JOIN divisions div
ON div.code = usr.division_code

SELECT * FROM filings
-- TEST QUERY
SELECT cmtl.list_id, cmtl.item_number FROM committee_lists cmtl
INNER JOIN committees cmt ON cmt.id = cmtl.committee_id

SELECT * FROM user_credentials

SELECT * FROM receivings
SELECT * FROM agendas
SELECT * FROM approves
SELECT * FROM releasings
SELECT * FROM filings
SELECT * FROM borrower_histories
SELECT * from employee_performaces
SELECT * FROM user_activities

SELECT * FROM divisions
SELECT * FROM committees
SELECT * FROM departments
SELECT * FROM committee_lists
SELECT * FROM proponents
SELECT * FROM routings
SELECT * FROM folders
SELECT * FROM trackings

SELECT * FROM committee_lists WHERE item_number = '4411-2312'
DELETE FROM committee_lists WHERE item_number = '4411-2312'

nextval('proponents_proponent_id_seq'::regclass)

DELETE FROM proponents WHERE name = '';
UPDATE routings SET remarks = 'Forwarded to Mayor' WHERE doc_id = 4
SELECT * FROM view_committees WHERE item_number = '3512-1345'
SELECT * FROM view_routings WHERE document_tag = 'Referred to Committee' AND doc_id = 6

SELECT * FROM view_routings 
WHERE document_tag = 'For Agenda' 
AND doc_id = 4

SELECT * FROM agendas WHERE item_number = '2023-0869'
SELECT * FROM filings
SELECT * FROM file_paths
SELECT * FROM trackings
SELECT * FROM FilingPaths

SELECT * FROM trackings WHERE calendared IS NOT NULL

SELECT * FROM view_borrower_history
SELECT * FROM user_credentials
SELECT * FROM view_routings
SELECT * FROM view_users
SELECT COUNT(*) FROM divisions
SELECT COUNT(*) FROM view_committees

-- WARNING DO NOT RUN THIS SCRIPT
SELECT truncate_tables()
SELECT reset_tables()
-- ------------------------------

SELECT * FROM view_routings WHERE document_tag = 'For Agenda' OR document_tag = 'For information of the whole body'
SELECT * FROM view_routings WHERE document_tag = 'For Agenda' OR document_tag = 'Kept in Secretariat'
SELECT * FROM view_routings WHERE document_tag = 'Forwarded to Secretariat'
INSERT INTO divisions (name, code) VALUES ('Head Office','HOD'), ('Records','SPCRD'), ('Secretariat','SPCSD')

UPDATE routings SET remarks = 'Forwarded to Mayor' WHERE doc_id = 2

SELECT id, fullname, password, division_code, is_reset, to_char(created_at, 'Mon. DD,YYYY') AS date_registered FROM user_credentials
SELECT * FROM view_users
SELECT * FROM trackings
SELECT * FROM view_routings
INSERT INTO committees (name) VALUES ('Agriculture And Economic Productivity'), ('Anti-Drug Abuse'),('Barangay Affairs'), ('Basic Education'),('Blue Ribbon Committee'), ('Civil Society Organizations and Cooperatives'),('Disaster Risk Reduction and Management'), ('Energy, Transportation & Telecommunication'),('Environmental Protection and Solid Waste Management'), ('Ethics & Discipline'),('Finance and Appropriations'), ('Games, Amusements and Illegal Gambling'),('Government Contracts, Legal Matters, Engineering, Public Works and Zonification'), ('Health and Sanitation'),('Higher Education, Technical and Vocational Courses'), ('History, Arts and Culture'),('Housing, Land Use and Estate Development'), ('Human Resources, Good Governance, Public Ethics and Accountability'),('Human Rights and Public Information'), ('Information Technology, e-Commerce and Mass Media'),('International Affairs'), ('Labor and Employment'),('Markets, Slaughterhouse and Government Economic Enterprise'), ('Rules and Privileges'),('Sports Development'), ('Tourism'),('Trade, Commerce, and Industry'), ('Urban Poor and Livelihood'),('Water Resources Management & Development'), ('Ways and Means'),('Welfare and Protection of Family, Women, Children, Senior Citizens, Persons with Disability and Gender Equality'), ('Youth')

SELECT COUNT(agenda_id) FROM routings WHERE agenda_id IS NULL
SELECT COUNT(*) FROM view_committees WHERE committee_id = 

SELECT * FROM view_routings WHERE document_tag = 'For Releasing' AND doc_id = 1
SELECT * FROM view_committees WHERE item_number = '2023-0002'
SELECT * FROM approves WHERE item_number = '2023-0002'

SELECT * FROM activity_loggers
SELECT * FROM view_routings WHERE document_tag = 'For Releasing' AND doc_id = 1

-- FUNCTIONS
CREATE OR REPLACE FUNCTION update_agenda(_agenda_id int, _tag text, _remarks text, _encoder)
RETURNS int AS
$$
DECLARE
BEGIN
	UPDATE agendas SET agenda_tag = _tag, agenda_remarks = remarks, modified_by = _encoder
	WHERE agenda_id = _agenda_id;
	RETURN 100
END
$$
LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION truncate_tables()
RETURNS text AS
$$
BEGIN
	TRUNCATE TABLE activity_loggers;
	TRUNCATE TABLE employee_performaces;
	TRUNCATE TABLE borrower_histories;
	TRUNCATE TABLE activity_loggers;
	TRUNCATE TABLE receivings;
	TRUNCATE TABLE agendas;
	TRUNCATE TABLE approves;
	TRUNCATE TABLE releasings;
	TRUNCATE TABLE filings;
	TRUNCATE TABLE committee_lists;
	TRUNCATE TABLE routings;
	TRUNCATE TABLE trackings;
	TRUNCATE TABLE event_calendars;
	RETURN 'All tables are truncated';
END;
$$
LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION reset_tables()
RETURNS text AS
$$
BEGIN
	DROP VIEW view_agendas;
	DROP VIEW view_committees;
	DROP VIEW view_borrower_history;
	DROP VIEW view_routings;
	DROP VIEW view_users;
	DROP TABLE receivings;
	DROP TABLE activity_loggers;
	DROP TABLE employee_performaces;
	DROP TABLE borrower_histories;
	DROP TABLE agendas;
	DROP TABLE approves;
	DROP TABLE releasings;
	DROP TABLE filings;
	DROP TABLE committee_lists;
	DROP TABLE routings;
	DROP TABLE trackings;
	DROP TABLE event_calendars;
	RETURN 'All tables are DROPPED!';
END;
$$
LANGUAGE plpgsql;


SELECT * FROM event_calendars
INSERT INTO event_calendars (title,start) 
VALUES ('Session Day!','2024-01-09'),
('Session Day!','2024-01-20'),
('Session Day!','2024-01-27')

SELECT CONCAT(title, COUNT(*)) title, start
FROM event_calendars
GROUP BY start, title

SELECT * FROM event_calendars
SELECT * FROM folders
SELECT * FROM activity_loggers WHERE user_id = 5
SELECT COUNT(records_captured) AS records_captured, TO_CHAR(created_at, 'Mon. DD,YYYY') AS date FROM employee_performaces WHERE user_id = 6 GROUP BY date

SELECT DISTINCT TO_CHAR(created_at, 'Mon. DD,YYYY') AS date, COUNT(records_captured) as records_captured, COUNT(records_retrieved) as records_retrieved FROM employee_performaces WHERE user_id = 6 GROUP BY date