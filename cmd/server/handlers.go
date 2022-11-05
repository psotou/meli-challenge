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
	if err := row.Scan(&emp.Id, &emp.Status, &emp.Department, &emp.DepartmentCode, &emp.DateIn, &emp.DateOut, &emp.Username); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, emp)
}

func createEmployees(c echo.Context) error {
	var emps []Employee
	var count int64
	if err := c.Bind(&emps); err != nil {
		return err
	}

	sql := `INSERT INTO pasidb.employee (status, department, department_code, date_in, username, inserted_at, updated_at) 
    VALUES (?, ?, ?, ?, ?, NOW(), NOW())`

	for _, emp := range emps {
		inserted, _ := db.Exec(sql, emp.Status, emp.Department, emp.DepartmentCode, emp.DateIn, emp.Username)
		rowsAffected, err := inserted.RowsAffected()
		if err != nil {
			return err
		}
		count += rowsAffected
	}
	log.Printf("rows afftected: %v", count)

	return c.JSON(http.StatusCreated, emps)
}

func updateEmployees(c echo.Context) error {
	var emps []Employee
	var count int64
	if err := c.Bind(&emps); err != nil {
		return err
	}

	sql := `UPDATE pasidb.employee
    SET status = ?, date_out = ?, updated_at = NOW()
    WHERE username = ?`

	for _, emp := range emps {
		updated, _ := db.Exec(sql, emp.Status, emp.DateOut, emp.Username)
		rowsAffected, err := updated.RowsAffected()
		if err != nil {
			return err
		}
		count += rowsAffected
	}
	log.Printf("rows afftected: %v\n", count)

	return c.JSON(http.StatusCreated, emps)
}

func createRoles(c echo.Context) error {
	var roles []Role
	var count int64
	if err := c.Bind(&roles); err != nil {
		return err
	}
	sql := `INSERT INTO pasidb.role (role_id, role_name, username, inserted_at, updated_at)
    VALUES (?, ?, ?, NOW(), NOW())`

	for _, role := range roles {
		inserted, _ := db.Exec(sql, role.RoleId, role.RoleName, role.Username)
		rowsAffected, err := inserted.RowsAffected()
		if err != nil {
			return err
		}
		count += rowsAffected
	}
	log.Printf("rows afftected: %v\n", count)

	return c.JSON(http.StatusCreated, roles)
}

func createApplications(c echo.Context) error {
	var apps []Application
	var count int64
	if err := c.Bind(&apps); err != nil {
		return err
	}
	sql := `INSERT INTO pasidb.application (app_id, app_name, role_id, is_critical, inserted_at, updated_at)
    VALUES (?, ?, ?, ?, NOW(), NOW())`

	for _, app := range apps {
		inserted, _ := db.Exec(sql, app.AppId, app.AppName, app.RoleId, app.IsCritical)
		rowsAffected, err := inserted.RowsAffected()
		if err != nil {
			return err
		}
		count += rowsAffected
	}
	log.Printf("rows afftected: %v\n", count)

	return c.JSON(http.StatusCreated, apps)
}

func createDbAccesses(c echo.Context) error {
	var dbaccesses []DBAccess
	var count int64
	if err := c.Bind(&dbaccesses); err != nil {
		return err
	}
	sql := "INSERT INTO pasidb.db_access (username, `table`, is_pii, inserted_at, updated_at) " +
		"VALUES (?, ?, ?, NOW(), NOW())"

	for _, dbaccess := range dbaccesses {
		inserted, _ := db.Exec(sql, dbaccess.Username, dbaccess.Table, dbaccess.IsPII)
		rowsAffected, err := inserted.RowsAffected()
		if err != nil {
			return err
		}
		count += rowsAffected
	}
	log.Printf("rows afftected: %v\n", count)

	return c.JSON(http.StatusCreated, dbaccesses)
}
