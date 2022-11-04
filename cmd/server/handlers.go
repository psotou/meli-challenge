package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func getEmployee(c echo.Context) error {
	var emp Employee
	sql := `SELECT id, status, department, department_code, date_in, coalesce(date_out, ''), username
    FROM pasidb.employee
    WHERE username = ?`
	row := db.QueryRow(sql, c.Param("username"))
	if err := row.Scan(&emp.Id, &emp.Status, &emp.Username, &emp.DepartmentCode, &emp.Department, &emp.DateIn, &emp.DateOut); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, emp)
}

func createEmployee(c echo.Context) error {
	var emp Employee
	if err := c.Bind(&emp); err != nil {
		return err
	}

	sql := `INSERT INTO pasidb.employee (status, department, department_code, date_in, username, inserted_at, updated_at) 
    VALUES ('Active', ?, ?, ?, ?, NOW(), NOW())`
	inserted, _ := db.Exec(sql, emp.Department, emp.DepartmentCode, emp.DateIn, emp.Username)
	rowsAffected, err := inserted.RowsAffected()
	if err != nil {
		return err
	}
	log.Printf("rows afftected: %v", rowsAffected)

	return c.JSON(http.StatusCreated, emp)
}

func updateEmployee(c echo.Context) error {
	var emp Employee
	if err := c.Bind(&emp); err != nil {
		return err
	}

	sql := `UPDATE pasidb.employee
    SET status = ?, date_out = ?, updated_at = NOW()
    WHERE username = ?`
	updated, _ := db.Exec(sql, emp.Status, emp.DateOut, c.Param("username"))
	rowsAffected, err := updated.RowsAffected()
	if err != nil {
		return err
	}
	log.Printf("rows afftected: %v\n", rowsAffected)

	return c.JSON(http.StatusCreated, emp)
}

func createRole(c echo.Context) error {
	var role Role
	if err := c.Bind(&role); err != nil {
		return err
	}
	sql := `INSERT INTO pasidb.role (role_id, role_name, username, inserted_at, updated_at)
    VALUES (?, ?, ?, NOW(), NOW())`
	inserted, _ := db.Exec(sql, role.RoleId, role.RoleName, role.Username)
	rowsAffected, err := inserted.RowsAffected()
	if err != nil {
		return err
	}
	log.Printf("rows afftected: %v\n", rowsAffected)

	return c.JSON(http.StatusCreated, role)
}

func createApplication(c echo.Context) error {
	var app Application
	if err := c.Bind(&app); err != nil {
		return err
	}
	sql := `INSERT INTO pasidb.application (app_id, app_name, role_id, is_critical, inserted_at, updated_at)
    VALUES (?, ?, ?, ?, NOW(), NOW())`
	inserted, _ := db.Exec(sql, app.AppId, app.AppName, app.RoleId, app.IsCritical)
	rowsAffected, err := inserted.RowsAffected()
	if err != nil {
		return err
	}
	log.Printf("rows afftected: %v\n", rowsAffected)

	return c.JSON(http.StatusCreated, app)
}

func createDbAccess(c echo.Context) error {
	var dbaccess DBAccess
	if err := c.Bind(&dbaccess); err != nil {
		return err
	}
	sql := "INSERT INTO pasidb.db_access (username, `table`, is_pii, inserted_at, updated_at) " +
		"VALUES (?, ?, ?, NOW(), NOW())"
	inserted, _ := db.Exec(sql, dbaccess.Username, dbaccess.Table, dbaccess.IsPII)
	rowsAffected, err := inserted.RowsAffected()
	if err != nil {
		return err
	}
	log.Printf("rows afftected: %v\n", rowsAffected)

	return c.JSON(http.StatusCreated, dbaccess)
}
