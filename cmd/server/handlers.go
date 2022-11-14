package main

import (
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	RiskModel RiskModeler
}

func NewHandler(riskmodel RiskModeler) *Handler {
	return &Handler{riskmodel}
}

func (h *Handler) CreateEmployees(c echo.Context) error {
	employees, err := h.RiskModel.InsertEmployees(c)
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

func (h *Handler) CreateRoles(c echo.Context) error {
	roles, err := h.RiskModel.InsertRoles(c)
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

func (h *Handler) CreateApplications(c echo.Context) error {
	apps, err := h.RiskModel.InsertApplications(c)
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

func (h *Handler) CreateDbAccesses(c echo.Context) error {
	dbaccesses, err := h.RiskModel.InsertDbAccess(c)
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

func (h *Handler) UpdateEmployees(c echo.Context) error {
	employees, err := h.RiskModel.UpdateEmployees(c)
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

func (h *Handler) GetEmployeeRisk(c echo.Context) error {
	empRisk, err := h.RiskModel.FindByUsername(c.Param("username"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, empRisk)
}

func (h *Handler) GetDepartmentRisk(c echo.Context) error {
	dpmtRisk, err := h.RiskModel.FindByDepartmentCode(c.Param("code"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, dpmtRisk)
}
