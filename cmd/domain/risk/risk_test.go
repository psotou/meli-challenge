package risk_test

import (
	"net/http"
	"net/http/httptest"
	"personal/risk-calculator/cmd/domain/risk"
	"personal/risk-calculator/cmd/service/riskindex"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// mocks
type RiskPersistenceStub struct{}

func (rp *RiskPersistenceStub) FindByUsername(username string) (risk.EmployeeRisk, error) {
	return risk.EmployeeRisk{
		Username:         "testusername",
		EmployeeRiskCode: 123,
		EmployeeRisk:     "test risk",
	}, nil
}

func (rp *RiskPersistenceStub) FindByDepartmentCode(code string) (risk.DepartmentRisk, error) {
	return risk.DepartmentRisk{
		DepartmentCode:     1313,
		Department:         "test department",
		DepartmentRiskCode: 12,
		DepartmentRisk:     "test risk",
	}, nil
}

func TestGetRiskIndex(t *testing.T) {
	t.Run("should return the risk index of an employee", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/employeerisk/:username")
		c.SetParamNames("username")
		c.SetParamValues("testusername")

		rp := &RiskPersistenceStub{}
		h := riskindex.NewHandler(rp)

		expectedEmployeeRisk := "{\"username\":\"testusername\",\"employee_risk_code\":123,\"employee_risk\":\"test risk\"}\n"

		if assert.NoError(t, h.GetEmployeeRisk(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, expectedEmployeeRisk, rec.Body.String())
		}
	})

	t.Run("should return the risk index of a department", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/departmentrisk/:code")
		c.SetParamNames("code")
		c.SetParamValues("12")

		rp := &RiskPersistenceStub{}
		h := riskindex.NewHandler(rp)

		expectedDepartmentRisk := "{\"department_code\":1313,\"department\":\"test department\",\"department_risk_code\":12,\"department_risk\":\"test risk\"}\n"

		if assert.NoError(t, h.GetDepartmentRisk(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, expectedDepartmentRisk, rec.Body.String())
		}
	})
}
