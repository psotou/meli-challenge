// employee aggregate is the composition of the entities that
// comprise what an employee actually is within a company
package employee

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
)

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

type employeePersistence struct {
	db *sql.DB
}

func NewEmployeePersistence(db *sql.DB) *employeePersistence {
	return &employeePersistence{db: db}
}

func (ep *employeePersistence) InsertEmployees(c echo.Context) ([]Employee, error) {
	var emps []Employee
	var count int64
	if err := c.Bind(&emps); err != nil {
		return []Employee{}, err
	}

	sql := `INSERT INTO pasidb.employee (status, department, department_code, date_in, username, inserted_at, updated_at)
    VALUES (?, ?, ?, ?, ?, NOW(), NOW())`

	for _, emp := range emps {
		inserted, err := ep.db.Exec(sql, emp.Status, emp.Department, emp.DepartmentCode, emp.DateIn, emp.Username)
		if err != nil {
			return []Employee{}, err
		}

		rowsAffected, err := inserted.RowsAffected()
		if err != nil {
			return []Employee{}, err
		}
		count += rowsAffected
	}
	log.Printf("rows afftected: %v", count)

	return emps, nil
}

func (ep *employeePersistence) InsertRoles(c echo.Context) ([]Role, error) {
	var roles []Role
	var count int64
	if err := c.Bind(&roles); err != nil {
		return []Role{}, err
	}
	sql := `INSERT INTO pasidb.role (role_id, role_name, username, inserted_at, updated_at)
    VALUES (?, ?, ?, NOW(), NOW())`

	for _, role := range roles {
		inserted, err := ep.db.Exec(sql, role.RoleId, role.RoleName, role.Username)
		if err != nil {
			return []Role{}, err
		}
		rowsAffected, err := inserted.RowsAffected()
		if err != nil {
			return []Role{}, err
		}
		count += rowsAffected
	}
	log.Printf("rows afftected: %v\n", count)

	return roles, nil
}

func (ep *employeePersistence) InsertApplications(c echo.Context) ([]Application, error) {
	var apps []Application
	var count int64
	if err := c.Bind(&apps); err != nil {
		return []Application{}, err
	}
	sql := `INSERT INTO pasidb.application (app_id, app_name, role_id, is_critical, inserted_at, updated_at)
    VALUES (?, ?, ?, ?, NOW(), NOW())`

	for _, app := range apps {
		inserted, err := ep.db.Exec(sql, app.AppId, app.AppName, app.RoleId, app.IsCritical)
		if err != nil {
			return []Application{}, err
		}
		rowsAffected, err := inserted.RowsAffected()
		if err != nil {
			return []Application{}, err
		}
		count += rowsAffected
	}
	log.Printf("rows afftected: %v\n", count)

	return apps, nil
}

func (ep *employeePersistence) InsertDbAccess(c echo.Context) ([]DBAccess, error) {
	var dbaccesses []DBAccess
	var count int64
	if err := c.Bind(&dbaccesses); err != nil {
		return []DBAccess{}, err
	}
	sql := "INSERT INTO pasidb.db_access (username, `table`, is_pii, inserted_at, updated_at) " +
		"VALUES (?, ?, ?, NOW(), NOW())"

	for _, dbaccess := range dbaccesses {
		inserted, err := ep.db.Exec(sql, dbaccess.Username, dbaccess.Table, dbaccess.IsPII)
		if err != nil {
			return []DBAccess{}, err
		}
		rowsAffected, err := inserted.RowsAffected()
		if err != nil {
			return []DBAccess{}, err
		}
		count += rowsAffected
	}
	log.Printf("rows afftected: %v\n", count)

	return dbaccesses, nil
}

func (ep *employeePersistence) UpdateEmployees(c echo.Context) ([]Employee, error) {
	var emps []Employee
	var count int64
	if err := c.Bind(&emps); err != nil {
		return []Employee{}, err
	}

	sql := `UPDATE pasidb.employee
    SET status = ?, date_out = ?, updated_at = NOW()
    WHERE username = ?`

	for _, emp := range emps {
		updated, err := ep.db.Exec(sql, emp.Status, emp.DateOut, emp.Username)
		if err != nil {
			return []Employee{}, err
		}
		rowsAffected, err := updated.RowsAffected()
		if err != nil {
			return []Employee{}, err
		}
		count += rowsAffected
	}
	log.Printf("rows afftected: %v\n", count)

	return emps, nil
}
