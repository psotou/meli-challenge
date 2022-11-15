// risk aggregate
package risk

import "database/sql"

type EmployeeRisk struct {
	Username         string `json:"username"`
	EmployeeRiskCode int    `json:"employee_risk_code"`
	EmployeeRisk     string `json:"employee_risk"`
}

type DepartmentRisk struct {
	DepartmentCode     int    `json:"department_code"`
	Department         string `json:"department"`
	DepartmentRiskCode int    `json:"department_risk_code"`
	DepartmentRisk     string `json:"department_risk"`
}

type riskPersistence struct {
	db *sql.DB
}

func NewRiskPersistence(db *sql.DB) *riskPersistence {
	return &riskPersistence{db: db}
}

func (rp *riskPersistence) FindByUsername(username string) (EmployeeRisk, error) {
	var empRisk EmployeeRisk
	sql := `SELECT username, employee_risk_code, employee_risk
    FROM pasidb.employee_risk_view
    WHERE username = ?`
	row := rp.db.QueryRow(sql, username)
	if err := row.Scan(&empRisk.Username, &empRisk.EmployeeRiskCode, &empRisk.EmployeeRisk); err != nil {
		return EmployeeRisk{}, err
	}

	return empRisk, nil
}

func (rp *riskPersistence) FindByDepartmentCode(code string) (DepartmentRisk, error) {
	var dpmtRisk DepartmentRisk
	sql := `SELECT department_code, department, department_risk_code, department_risk
    FROM pasidb.department_risk_view
    WHERE department_code = ?`
	row := rp.db.QueryRow(sql, code)
	if err := row.Scan(&dpmtRisk.DepartmentCode, &dpmtRisk.Department, &dpmtRisk.DepartmentRiskCode, &dpmtRisk.DepartmentRisk); err != nil {
		return DepartmentRisk{}, err
	}

	return dpmtRisk, nil
}
