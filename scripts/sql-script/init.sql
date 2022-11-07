USE pasidb;

CREATE TABLE IF NOT EXISTS employee (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY
    , status VARCHAR(255)
    , department_code INT
    , department VARCHAR(255)
    , date_in VARCHAR(255)
    , date_out VARCHAR(255)
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

-- Business logic
CREATE OR REPLACE VIEW risk_view AS
WITH cte AS (
    SELECT
        e.username 
        , e.status
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
        *
        , CASE
            WHEN status = 'Inactive' THEN 0
            WHEN status = 'Active' THEN (select table_risk + app_risk + role_risk)
            END emp_risk
        , (select table_risk + app_risk) dept_risk
    FROM cte
)
, max_risk AS (
    SELECT 
        *
        , MAX(emp_risk) OVER(PARTITION BY username) employee_risk_code
        , MAX(dept_risk) OVER(PARTITION BY department_code) department_risk_code
    FROM agg_fields
)
SELECT
    username
    , department
    , department_code
    , employee_risk_code
    , department_risk_code
    , CASE
        WHEN employee_risk_code = 0 THEN 'no risk'
        WHEN employee_risk_code IN (1, 2) THEN 'low'
        WHEN employee_risk_code IN (3, 4) THEN 'mid'
        WHEN employee_risk_code IN (5, 6) THEN 'high'
        WHEN employee_risk_code = 7  THEN 'very high'
        END employee_risk
    , CASE
        WHEN department_risk_code = 0 THEN 'no risk'
        WHEN department_risk_code = 1 THEN 'low'
        WHEN department_risk_code = 2 THEN 'mid'
        WHEN department_risk_code = 3 THEN 'high'
        END department_risk
FROM max_risk;

CREATE OR REPLACE VIEW employee_risk_view AS 
SELECT 
    username
    , employee_risk_code
    , employee_risk
FROM risk_view
GROUP BY username, employee_risk_code;

CREATE OR REPLACE VIEW department_risk_view AS 
SELECT 
    department
    , department_code
    , department_risk_code
    , department_risk
FROM risk_view
GROUP BY department, department_code, department_risk_code;

