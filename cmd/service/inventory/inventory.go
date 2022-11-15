package inventory

import (
	"net/http"
	"personal/risk-calculator/cmd/domain/employee"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type handler struct {
	inventory employee.EmployeeRepository
}

func NewHandler(inventory employee.EmployeeRepository) *handler {
	return &handler{inventory: inventory}
}

func (h *handler) CreateEmployees(c echo.Context) error {
	employees, err := h.inventory.InsertEmployees(c)
	if err != nil {
		switch err.(type) {
		case *echo.HTTPError:
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		default:
			// the default case handles all db related errors
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusCreated, employees)
}

func (h *handler) CreateRoles(c echo.Context) error {
	roles, err := h.inventory.InsertRoles(c)
	if err != nil {
		switch err.(type) {
		case *echo.HTTPError:
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		default:
			// the default case handles all db related errors
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusCreated, roles)
}

func (h *handler) CreateApplications(c echo.Context) error {
	apps, err := h.inventory.InsertApplications(c)
	if err != nil {
		switch err.(type) {
		case *echo.HTTPError:
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		default:
			// the default case handles all db related errors
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusCreated, apps)
}

func (h *handler) CreateDbAccesses(c echo.Context) error {
	dbaccesses, err := h.inventory.InsertDbAccess(c)
	if err != nil {
		switch err.(type) {
		case *echo.HTTPError:
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		default:
			// the default case handles all db related errors
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusCreated, dbaccesses)
}

func (h *handler) UpdateEmployees(c echo.Context) error {
	employees, err := h.inventory.UpdateEmployees(c)
	if err != nil {
		switch err.(type) {
		case *echo.HTTPError:
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		case *mysql.MySQLError:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		default:
			return echo.NewHTTPError(http.StatusConflict, err.Error())
		}
	}

	return c.JSON(http.StatusCreated, employees)
}
