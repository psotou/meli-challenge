package employee

import "github.com/labstack/echo/v4"

type EmployeeRepository interface {
	InsertEmployees(echo.Context) ([]Employee, error)
	InsertRoles(echo.Context) ([]Role, error)
	InsertApplications(echo.Context) ([]Application, error)
	InsertDbAccess(echo.Context) ([]DBAccess, error)
	UpdateEmployees(echo.Context) ([]Employee, error)
}
