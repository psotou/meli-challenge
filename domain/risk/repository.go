package risk

type RiskRepository interface {
	FindByUsername(string) (EmployeeRisk, error)
	FindByDepartmentCode(string) (DepartmentRisk, error)
}
