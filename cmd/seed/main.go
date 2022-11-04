package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const (
	employeeJSON    = "json-data/employee.json"
	applicationJSON = "json-data/application.json"
	roleJSON        = "json-data/role.json"
	dbaccesJSON     = "json-data/db_access.json"
	employeeCSV     = "csv-data/employee.csv"
	applicationCSV  = "csv-data/application.csv"
	roleCSV         = "csv-data/role.csv"
	dbaccessCSV     = "csv-data/db_access.csv"
)

// we leave out PII fields
type Employee struct {
	Status         string `json:"status"`
	DepartmentCode int    `json:"department_code"`
	DateIn         string `json:"date_in"`
	DateOut        string `json:"date_out"`
	Username       string `json:"username"`
}

type Application struct {
	AppId      int    `json:"app_id"`
	AppName    string `json:"app_name"`
	RoleId     int    `json:"role_id"`
	IsCritical int    `json:"is_critical"`
}

type Role struct {
	RoleId   int    `json:"role_id"`
	RoleName string `json:"role_name"`
	Username string `json:"username"`
}

type DBAccess struct {
	Username string `json:"username"`
	Table    string `json:"table"`
	IsPII    int    `json:"is_pii"`
}

func main() {
	var employees []Employee
	var roles []Role
	var applications []Application
	var dbaccesses []DBAccess

	employeeContent, _ := os.ReadFile(employeeJSON)
	json.Unmarshal(employeeContent, &employees)
	log.Printf("writing content to file %s", employeeCSV)
	err := csvWriter(employees)
	if err != nil {
		log.Fatalf("couldn't writw csv file: %s", err.Error())
	}

	roleContent, _ := os.ReadFile(roleJSON)
	json.Unmarshal(roleContent, &roles)
	log.Printf("writing content to file %s", roleCSV)
	err = csvWriter(roles)
	if err != nil {
		log.Fatalf("couldn't writw csv file: %s", err.Error())
	}

	applicationContent, _ := os.ReadFile(applicationJSON)
	json.Unmarshal(applicationContent, &applications)
	log.Printf("writing content to file %s", applicationCSV)
	err = csvWriter(applications)
	if err != nil {
		log.Fatalf("couldn't writw csv file: %s", err.Error())
	}

	dbaccessContent, _ := os.ReadFile(dbaccesJSON)
	json.Unmarshal(dbaccessContent, &dbaccesses)
	log.Printf("writing content to file %s", dbaccessCSV)
	err = csvWriter(dbaccesses)
	if err != nil {
		log.Fatalf("couldn't writw csv file: %s", err.Error())
	}

}

func csvWriter(records interface{}) error {
	var data [][]string
	var w *csv.Writer

	switch records.(type) {
	case []Employee:
		file, err := os.Create(employeeCSV)
		if err != nil {
			log.Fatalf("couldn't create file: %s\n", err.Error())
		}
		defer file.Close()

		w = csv.NewWriter(file)
		for _, r := range records.([]Employee) {
			// the order in the mysql table is as follows:
			// status, department_code, date_in, date_out, username
			row := []string{r.Status, fmt.Sprint(r.DepartmentCode), r.DateIn, r.DateOut, r.Username}
			data = append(data, row)
		}
	case []Role:
		file, err := os.Create(roleCSV)
		if err != nil {
			log.Fatalf("couldn't create file: %s\n", err.Error())
		}
		defer file.Close()

		w = csv.NewWriter(file)
		for _, r := range records.([]Role) {
			// the order in the mysql table is as follows:
			// role_id, role_name, username
			row := []string{fmt.Sprint(r.RoleId), r.RoleName, r.Username}
			data = append(data, row)
		}
	case []Application:
		file, err := os.Create(applicationCSV)
		if err != nil {
			log.Fatalf("couldn't create file: %s\n", err.Error())
		}
		defer file.Close()

		w = csv.NewWriter(file)
		for _, r := range records.([]Application) {
			// the order in the mysql table is as follows:
			// app_id, app_name, role_id, is_critical
			row := []string{fmt.Sprint(r.AppId), r.AppName, fmt.Sprint(r.RoleId), fmt.Sprint(r.IsCritical)}
			data = append(data, row)
		}
	case []DBAccess:
		file, err := os.Create(dbaccessCSV)
		if err != nil {
			log.Fatalf("couldn't create file: %s\n", err.Error())
		}
		defer file.Close()

		w = csv.NewWriter(file)
		for _, r := range records.([]DBAccess) {
			// the order in the mysql table is as follows:
			//username, table, is_pii
			row := []string{r.Username, r.Table, fmt.Sprint(r.IsPII)}
			data = append(data, row)
		}
	}

	return w.WriteAll(data)
}
