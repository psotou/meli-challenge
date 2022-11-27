package riskindex

import (
	"net/http"
	"personal/risk-calculator/domain/risk"

	"github.com/labstack/echo/v4"
)

type handler struct {
	risk risk.RiskRepository
}

func NewHandler(risk risk.RiskRepository) *handler {
	return &handler{risk: risk}
}

func (h *handler) GetEmployeeRisk(c echo.Context) error {
	empRisk, err := h.risk.FindByUsername(c.Param("username"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, empRisk)
}

func (h *handler) GetDepartmentRisk(c echo.Context) error {
	dpmtRisk, err := h.risk.FindByDepartmentCode(c.Param("code"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, dpmtRisk)
}
