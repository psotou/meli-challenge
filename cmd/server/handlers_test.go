package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type RiskModelStub struct{}

func (rm *RiskModelStub) FindByUsername(username string) (EmployeeRisk, error) {
	return EmployeeRisk{
		Username:         "testusername",
		EmployeeRiskCode: 123,
		EmployeeRisk:     "test risk",
	}, nil
}

func TestGetEmployeeRisk(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/employeerisk/:username")
	c.SetParamNames("username")
	c.SetParamValues("testusername")

	rm := &RiskModelStub{}
	h := NewHandler(rm)

	expectedEmployeeRisk := "{\"username\":\"testusername\",\"employee_risk_code\":123,\"employee_risk\":\"test risk\"}\n"

	if assert.NoError(t, h.GetEmployeeRisk(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedEmployeeRisk, rec.Body.String())
	}
}

func (rm *RiskModelStub) FindByDepartmentCode(code string) (DepartmentRisk, error) {
	return DepartmentRisk{
		DepartmentCode:     1313,
		Department:         "test department",
		DepartmentRiskCode: 12,
		DepartmentRisk:     "test risk",
	}, nil
}

func TestGetDepartmentRisk(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/departmentrisk/:code")
	c.SetParamNames("code")
	c.SetParamValues("12")

	rm := &RiskModelStub{}
	h := NewHandler(rm)

	expectedDepartmentRisk := "{\"department_code\":1313,\"department\":\"test department\",\"department_risk_code\":12,\"department_risk\":\"test risk\"}\n"

	if assert.NoError(t, h.GetDepartmentRisk(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedDepartmentRisk, rec.Body.String())
	}
}

func (rm *RiskModelStub) InsertEmployees(c echo.Context) ([]Employee, error) {
	return []Employee{
		{
			Id:             1,
			Status:         "Active",
			DepartmentCode: 123,
			Department:     "test department",
			DateIn:         "10/10/2020",
			DateOut:        "None",
			Username:       "testusername",
		},
	}, nil
}

func TestCreateEmployees(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/employees")

	rm := &RiskModelStub{}
	h := NewHandler(rm)

	expectedEmployees := "[{\"id\":1,\"status\":\"Active\",\"department_code\":123,\"department\":\"test department\",\"date_in\":\"10/10/2020\",\"date_out\":\"None\",\"username\":\"testusername\"}]\n"

	if assert.NoError(t, h.CreateEmployees(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, expectedEmployees, rec.Body.String())
	}
}

func (rm *RiskModelStub) InsertRoles(c echo.Context) ([]Role, error) {
	return []Role{
		{
			RoleId:   12,
			RoleName: "role test",
			Username: "testusername",
		},
	}, nil
}

func TestCreateRoles(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/roles")

	rm := &RiskModelStub{}
	h := NewHandler(rm)

	expectedRoles := "[{\"role_id\":12,\"role_name\":\"role test\",\"username\":\"testusername\"}]\n"

	if assert.NoError(t, h.CreateRoles(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, expectedRoles, rec.Body.String())
	}
}

func (rm *RiskModelStub) InsertApplications(c echo.Context) ([]Application, error) {
	return []Application{
		{
			AppId:      23,
			AppName:    "app test",
			RoleId:     32,
			IsCritical: 0,
		},
	}, nil
}

func TestCreateApplications(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/applications")

	rm := &RiskModelStub{}
	h := NewHandler(rm)

	expectedApps := "[{\"app_id\":23,\"app_name\":\"app test\",\"role_id\":32,\"is_critical\":0}]\n"

	if assert.NoError(t, h.CreateApplications(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, expectedApps, rec.Body.String())
	}
}

func (rm *RiskModelStub) InsertDbAccess(c echo.Context) ([]DBAccess, error) {
	return []DBAccess{
		{
			Username: "testusername",
			Table:    "test table",
			IsPII:    1,
		},
	}, nil
}

func TestCreateDbAccesses(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/dbaccesses")

	rm := &RiskModelStub{}
	h := NewHandler(rm)

	expectedDbAccesses := "[{\"username\":\"testusername\",\"table\":\"test table\",\"is_pii\":1}]\n"

	if assert.NoError(t, h.CreateDbAccesses(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, expectedDbAccesses, rec.Body.String())
	}
}

func (rm *RiskModelStub) UpdateEmployees(c echo.Context) ([]Employee, error) {
	return []Employee{
		{
			Id:             12,
			Status:         "Inactive",
			DepartmentCode: 1234,
			Department:     "test department",
			DateIn:         "10/10/2020",
			DateOut:        "12/12/2022",
			Username:       "testusername",
		},
	}, nil
}

func TestUpdateEmployees(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/employees")

	rm := &RiskModelStub{}
	h := NewHandler(rm)

	expectedEmployees := "[{\"id\":12,\"status\":\"Inactive\",\"department_code\":1234,\"department\":\"test department\",\"date_in\":\"10/10/2020\",\"date_out\":\"12/12/2022\",\"username\":\"testusername\"}]\n"

	if assert.NoError(t, h.UpdateEmployees(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, expectedEmployees, rec.Body.String())
	}
}
