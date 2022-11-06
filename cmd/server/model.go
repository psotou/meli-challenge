package main

// FullName       string `json:"full_name"`
// DivCode        int    `json:"div_code"`
// Division       string `json:"division"`
// CountryId      int    `json:"country_id"`
// Country        string `json:"country"`
// Email          string `json:"email"`

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

// business logic
type EmployeeRisk struct {
	Username         string `json:"username"`
	EmployeeRiskCode int    `json:"employee_risk_code"`
	EmployeeRisk     string `json:"employee_risk"`
}

type DepartmentRisk struct {
	DepartmentCode     int    `json:"department_code"`
	Department         string `json:"deparment"`
	DepartmentRiskCode int    `json:"department_risk_code"`
	DepartmentRisk     string `json:"deparment_risk"`
}
