USE pasidb;

CREATE TABLE IF NOT EXISTS employee (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY
    , status VARCHAR(255)
    , department_code INT
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


LOAD DATA INFILE '/var/lib/mysql-files/seed/employee.csv'
INTO TABLE pasidb.employee
FIELDS TERMINATED BY ','
LINES TERMINATED BY '\n'
(status, department_code, @datein, @dateout, username, @insertedat, @updatedat)
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
