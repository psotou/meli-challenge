package inventory_test

import (
	"net/http"
	"net/http/httptest"
	"personal/risk-calculator/domain/employee"
	"personal/risk-calculator/service/inventory"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// mocks
type EmployeePersistenceStub struct{}

func (ep *EmployeePersistenceStub) InsertEmployees(c echo.Context) ([]employee.Employee, error) {
	return []employee.Employee{
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

func (ep *EmployeePersistenceStub) InsertRoles(c echo.Context) ([]employee.Role, error) {
	return []employee.Role{
		{
			RoleId:   12,
			RoleName: "role test",
			Username: "testusername",
		},
	}, nil
}

func (ep *EmployeePersistenceStub) InsertApplications(c echo.Context) ([]employee.Application, error) {
	return []employee.Application{
		{
			AppId:      23,
			AppName:    "app test",
			RoleId:     32,
			IsCritical: 0,
		},
	}, nil
}

func (ep *EmployeePersistenceStub) InsertDbAccess(c echo.Context) ([]employee.DBAccess, error) {
	return []employee.DBAccess{
		{
			Username: "testusername",
			Table:    "test table",
			IsPII:    1,
		},
	}, nil
}

func (ep *EmployeePersistenceStub) UpdateEmployees(c echo.Context) ([]employee.Employee, error) {
	return []employee.Employee{
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

func TestCreateEmployees(t *testing.T) {
	t.Run("should persist employee related data", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/employees")

		ep := &EmployeePersistenceStub{}
		h := inventory.NewHandler(ep)

		expectedEmployees := "[{\"id\":1,\"status\":\"Active\",\"department_code\":123,\"department\":\"test department\",\"date_in\":\"10/10/2020\",\"date_out\":\"None\",\"username\":\"testusername\"}]\n"

		if assert.NoError(t, h.CreateEmployees(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, expectedEmployees, rec.Body.String())
		}
	})
}

func TestCreteRoles(t *testing.T) {
	t.Run("should persist role related data", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/roles")

		ep := &EmployeePersistenceStub{}
		h := inventory.NewHandler(ep)

		expectedRoles := "[{\"role_id\":12,\"role_name\":\"role test\",\"username\":\"testusername\"}]\n"

		if assert.NoError(t, h.CreateRoles(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, expectedRoles, rec.Body.String())
		}
	})
}

func TestCreateApplications(t *testing.T) {
	t.Run("should persist application related data", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/applications")

		ep := &EmployeePersistenceStub{}
		h := inventory.NewHandler(ep)

		expectedApps := "[{\"app_id\":23,\"app_name\":\"app test\",\"role_id\":32,\"is_critical\":0}]\n"

		if assert.NoError(t, h.CreateApplications(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, expectedApps, rec.Body.String())
		}
	})
}

func TestCreateDbAccesses(t *testing.T) {
	t.Run("should persist db access data", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/dbaccesses")

		ep := &EmployeePersistenceStub{}
		h := inventory.NewHandler(ep)

		expectedDbAccesses := "[{\"username\":\"testusername\",\"table\":\"test table\",\"is_pii\":1}]\n"

		if assert.NoError(t, h.CreateDbAccesses(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, expectedDbAccesses, rec.Body.String())
		}
	})
}

func TestUpdateEmployees(t *testing.T) {
	t.Run("should update employee related data", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/employees")

		ep := &EmployeePersistenceStub{}
		h := inventory.NewHandler(ep)

		expectedEmployees := "[{\"id\":12,\"status\":\"Inactive\",\"department_code\":1234,\"department\":\"test department\",\"date_in\":\"10/10/2020\",\"date_out\":\"12/12/2022\",\"username\":\"testusername\"}]\n"

		if assert.NoError(t, h.UpdateEmployees(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, expectedEmployees, rec.Body.String())
		}
	})
}
