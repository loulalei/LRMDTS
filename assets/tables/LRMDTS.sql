CREATE VIEW view_committees AS
SELECT list_id, item_number, name, fullname 
FROM committee_lists cmtl
INNER JOIN committees cmt
ON cmt.id = cmtl.committee_id
INNER JOIN user_credentials usr
ON usr.id = cmtl.user_id

SELECT * FROM view_committees
SELECT * FROM committees
SELECT * FROM committee_lists
SELECT * FROM user_credentials
INSERT INTO committee_lists (item_number, committee_id, user_id)
VALUES ('2023-587',10,5)

DELETE FROM committee_lists WHERE list_id = 1