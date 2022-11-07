package main

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
)

type RiskModeler interface {
	InsertEmployees(echo.Context) ([]Employee, error)
	InsertRoles(echo.Context) ([]Role, error)
	InsertApplications(echo.Context) ([]Application, error)
	InsertDbAccess(echo.Context) ([]DBAccess, error)
	UpdateEmployees(echo.Context) ([]Employee, error)

	FindByUsername(string) (EmployeeRisk, error)
	FindByDepartmentCode(string) (DepartmentRisk, error)
}

type Employee struct {
	Id             int    `json:"id"`
	Status         string `json:"status"`
	DepartmentCode int    `json:"department_code"`
	Department     string `json:"department"`
	DateIn         string `json:"date_in"`
	DateOut        string `json:"date_out"`
	Username       string `json:"username"`
}

type Application struct {
	AppId      int    `json:"app_id"`
	AppName    string `json:"app_name"`
	RoleId     int    `json:"role_id"`
	IsCritical uint8  `json:"is_critical"`
}

type Role struct {
	RoleId   int    `json:"role_id"`
	RoleName string `json:"role_name"`
	Username string `json:"username"`
}

type DBAccess struct {
	Username string `json:"username"`
	Table    string `json:"table"`
	IsPII    uint8  `json:"is_pii"`
}

// business logic
type EmployeeRisk struct {
	Username         string `json:"username"`
	EmployeeRiskCode int    `json:"employee_risk_code"`
	EmployeeRisk     string `json:"employee_risk"`
}

type DepartmentRisk struct {
	DepartmentCode     int    `json:"department_code"`
	Department         string `json:"department"`
	DepartmentRiskCode int    `json:"department_risk_code"`
	DepartmentRisk     string `json:"department_risk"`
}

type RiskModel struct {
	db *sql.DB
}

func NewRiskModel(db *sql.DB) *RiskModel {
	return &RiskModel{db: db}
}

func (rm *RiskModel) InsertEmployees(c echo.Context) ([]Employee, error) {
	var emps []Employee
	var count int64
	if err := c.Bind(&emps); err != nil {
		return []Employee{}, err
	}

	sql := `INSERT INTO pasidb.employee (status, department, department_code, date_in, username, inserted_at, updated_at)
    VALUES (?, ?, ?, ?, ?, NOW(), NOW())`

	for _, emp := range emps {
		inserted, _ := rm.db.Exec(sql, emp.Status, emp.Department, emp.DepartmentCode, emp.DateIn, emp.Username)
		rowsAffected, err := inserted.RowsAffected()
		if err != nil {
			return []Employee{}, err
		}
		count += rowsAffected
	}
	log.Printf("rows afftected: %v", count)

	return emps, nil
}

func (rm *RiskModel) InsertRoles(c echo.Context) ([]Role, error) {
	var roles []Role
	var count int64
	if err := c.Bind(&roles); err != nil {
		return []Role{}, err
	}
	sql := `INSERT INTO pasidb.role (role_id, role_name, username, inserted_at, updated_at)
    VALUES (?, ?, ?, NOW(), NOW())`

	for _, role := range roles {
		inserted, _ := rm.db.Exec(sql, role.RoleId, role.RoleName, role.Username)
		rowsAffected, err := inserted.RowsAffected()
		if err != nil {
			return []Role{}, err
		}
		count += rowsAffected
	}
	log.Printf("rows afftected: %v\n", count)

	return roles, nil
}

func (rm *RiskModel) InsertApplications(c echo.Context) ([]Application, error) {
	var apps []Application
	var count int64
	if err := c.Bind(&apps); err != nil {
		return []Application{}, err
	}
	sql := `INSERT INTO pasidb.application (app_id, app_name, role_id, is_critical, inserted_at, updated_at)
    VALUES (?, ?, ?, ?, NOW(), NOW())`

	for _, app := range apps {
		inserted, _ := rm.db.Exec(sql, app.AppId, app.AppName, app.RoleId, app.IsCritical)
		rowsAffected, err := inserted.RowsAffected()
		if err != nil {
			return []Application{}, err
		}
		count += rowsAffected
	}
	log.Printf("rows afftected: %v\n", count)

	return apps, nil
}

func (rm *RiskModel) InsertDbAccess(c echo.Context) ([]DBAccess, error) {
	var dbaccesses []DBAccess
	var count int64
	if err := c.Bind(&dbaccesses); err != nil {
		return []DBAccess{}, err
	}
	sql := "INSERT INTO pasidb.db_access (username, `table`, is_pii, inserted_at, updated_at) " +
		"VALUES (?, ?, ?, NOW(), NOW())"

	for _, dbaccess := range dbaccesses {
		inserted, _ := rm.db.Exec(sql, dbaccess.Username, dbaccess.Table, dbaccess.IsPII)
		rowsAffected, err := inserted.RowsAffected()
		if err != nil {
			return []DBAccess{}, err
		}
		count += rowsAffected
	}
	log.Printf("rows afftected: %v\n", count)

	return dbaccesses, nil
}
func (rm *RiskModel) UpdateEmployees(c echo.Context) ([]Employee, error) {
	var emps []Employee
	var count int64
	if err := c.Bind(&emps); err != nil {
		return []Employee{}, err
	}

	sql := `UPDATE pasidb.employee
    SET status = ?, date_out = ?, updated_at = NOW()
    WHERE username = ?`

	for _, emp := range emps {
		updated, _ := rm.db.Exec(sql, emp.Status, emp.DateOut, emp.Username)
		rowsAffected, err := updated.RowsAffected()
		if err != nil {
			return []Employee{}, err
		}
		count += rowsAffected
	}
	log.Printf("rows afftected: %v\n", count)

	return emps, nil
}

func (rm *RiskModel) FindByUsername(username string) (EmployeeRisk, error) {
	var empRisk EmployeeRisk
	sql := `SELECT username, employee_risk_code, employee_risk
    FROM pasidb.employee_risk_view
    WHERE username = ?`
	row := rm.db.QueryRow(sql, username)
	if err := row.Scan(&empRisk.Username, &empRisk.EmployeeRiskCode, &empRisk.EmployeeRisk); err != nil {
		return EmployeeRisk{}, err
	}

	return empRisk, nil
}

func (rm *RiskModel) FindByDepartmentCode(code string) (DepartmentRisk, error) {
	var dpmtRisk DepartmentRisk
	sql := `SELECT department_code, department, department_risk_code, department_risk
    FROM pasidb.department_risk_view
    WHERE department_code = ?`
	row := rm.db.QueryRow(sql, code)
	if err := row.Scan(&dpmtRisk.DepartmentCode, &dpmtRisk.Department, &dpmtRisk.DepartmentRiskCode, &dpmtRisk.DepartmentRisk); err != nil {
		return DepartmentRisk{}, err
	}

	return dpmtRisk, nil
}
