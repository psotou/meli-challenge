USE pasidb;

CREATE TABLE IF NOT EXISTS employee (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY
    , status VARCHAR(255)
    , department_code INT
    , department VARCHAR(255)
    , date_in DATE
    , date_out DATE
    , username VARCHAR(255)
    , inserted_at DATETIME
    , updated_at DATETIME
);

CREATE TABLE IF NOT EXISTS role (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY
    , role_id INT 
    , role_name VARCHAR(255) 
    , username VARCHAR(255)
    , inserted_at DATETIME
    , updated_at DATETIME
);


CREATE TABLE IF NOT EXISTS application (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY
    , app_id INT 
    , app_name VARCHAR(255) 
    , role_id INT
    , is_critical TINYINT
    , inserted_at DATETIME
    , updated_at DATETIME
);

CREATE TABLE IF NOT EXISTS db_access (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY
    , username VARCHAR(255)
    , `table` VARCHAR(255)
    , is_pii TINYINT
    , inserted_at DATETIME
    , updated_at DATETIME
);

-- Load data from csv

LOAD DATA INFILE '/var/lib/mysql-files/seed/employee.csv'
INTO TABLE pasidb.employee
FIELDS TERMINATED BY ','
LINES TERMINATED BY '\n'
(status, department_code, department, @datein, @dateout, username, @insertedat, @updatedat)
SET date_in = STR_TO_DATE(@datein, '%d/%m/%Y')
    , date_out = STR_TO_DATE(NULLIF(@dateout, 'None'), '%d/%m/%Y')
    , inserted_at = NOW()
    , updated_at = NOW();

LOAD DATA INFILE '/var/lib/mysql-files/seed/role.csv'
INTO TABLE pasidb.role
FIELDS TERMINATED BY ','
LINES TERMINATED BY '\n'
(role_id, role_name, username, @insertedat, @updatedat)
SET inserted_at = NOW()
    , updated_at = NOW();

LOAD DATA INFILE '/var/lib/mysql-files/seed/application.csv'
INTO TABLE pasidb.application
FIELDS TERMINATED BY ','
LINES TERMINATED BY '\n'
(app_id, app_name, role_id, is_critical, @insertedat, @updatedat)
SET inserted_at = NOW()
    , updated_at = NOW();

LOAD DATA INFILE '/var/lib/mysql-files/seed/db_access.csv'
INTO TABLE pasidb.db_access
FIELDS TERMINATED BY ','
LINES TERMINATED BY '\n'
(username, `table`, is_pii, @insertedat, @updatedat)
SET inserted_at = NOW()
    , updated_at = NOW();

-- Create business logic
CREATE OR REPLACE VIEW risk_view AS
WITH cte AS (
    SELECT
        e.username 
        , e.department 
        , e.department_code
        , da.`table`
        , da.is_pii
        , r.role_name 
        , a.app_name 
        , a.is_critical 
        , CASE 
            WHEN da.is_pii = 1 THEN 2 
            WHEN da.is_pii = 0 THEN 1
            ELSE 0 END table_risk
        , CASE WHEN a.is_critical THEN 1 ELSE 0 END app_risk
        , CASE
            WHEN r.role_name IN ('ADMIN', 'ADMINS_SU') THEN 4
            WHEN r.role_name IN ('WRITER_CS') THEN 3
            WHEN r.role_name IN ('READER', 'READER_CS', 'USER_READ') THEN 2
            WHEN r.role_name IN ('OPERATOR', 'CONSULTANT') THEN 1
            ELSE 0 END role_risk
    FROM employee e
    LEFT JOIN role r ON r.username = e.username
    LEFT JOIN db_access da ON e.username = da.username
    LEFT JOIN application a ON r.role_id = a.role_id
)
, agg_fields AS (
    SELECT
        cte.*
        , (select table_risk + app_risk) dept_risk
        , (select table_risk + app_risk + role_risk) emp_risk
    FROM cte
)
, max_risk AS (
    SELECT 
        af.*
        , MAX(af.emp_risk) OVER(PARTITION BY af.username) employee_risk
        , MAX(af.dept_risk) OVER(PARTITION BY af.department_code) department_risk
    FROM agg_fields af 
)
SELECT
    mr.username
    , mr.department
    , mr.department_code
    , mr.employee_risk
    , mr.department_risk
    , CASE
        WHEN employee_risk = 0 THEN 'no risk'
        WHEN employee_risk IN (1, 2) THEN 'low'
        WHEN employee_risk IN (3, 4) THEN 'mid'
        WHEN employee_risk IN (5, 6) THEN 'high'
        WHEN employee_risk = 7  THEN 'very high'
        END employee_risk_desc
    , CASE
        WHEN department_risk = 0 THEN 'no risk'
        WHEN department_risk = 1 THEN 'low'
        WHEN department_risk = 2 THEN 'mid'
        WHEN department_risk = 3 THEN 'high'
        END department_risk_desc
FROM max_risk mr;

CREATE OR REPLACE VIEW employee_risk_view AS 
SELECT 
    username
    , employee_risk
    , employee_risk_desc
FROM risk_view
GROUP BY username, employee_risk;

CREATE OR REPLACE VIEW department_risk_view AS 
SELECT 
    department
    , department_code
    , department_risk
    , department_risk_desc
FROM risk_view
GROUP BY department, department_code, department_risk;

