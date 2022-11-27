package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"personal/risk-calculator/domain/employee"
	"personal/risk-calculator/domain/risk"
	"personal/risk-calculator/service/inventory"
	"personal/risk-calculator/service/riskindex"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	dsn := fmt.Sprintf("root:%s@tcp(mysqldb:3306)/%s", os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQL_DATABASE"))
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	inventoryHandler := inventory.NewHandler(employee.NewEmployeePersistence(db))
	riskIndexHandler := riskindex.NewHandler(risk.NewRiskPersistence(db))

	e.POST("/employees", inventoryHandler.CreateEmployees)
	e.PUT("/employees", inventoryHandler.UpdateEmployees)
	e.POST("/roles", inventoryHandler.CreateRoles)
	e.POST("/applications", inventoryHandler.CreateApplications)
	e.POST("/dbaccesses", inventoryHandler.CreateDbAccesses)

	e.GET("/employeerisk/:username", riskIndexHandler.GetEmployeeRisk)
	e.GET("/departmentrisk/:code", riskIndexHandler.GetDepartmentRisk)

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
	}

	e.Logger.Fatal(e.StartServer(server))
}
