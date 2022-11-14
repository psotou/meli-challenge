package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
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

	h := NewHandler(NewRiskModel(db))

	e.POST("/employees", h.CreateEmployees)
	e.PUT("/employees", h.UpdateEmployees)
	e.POST("/roles", h.CreateRoles)
	e.POST("/applications", h.CreateApplications)
	e.POST("/dbaccesses", h.CreateDbAccesses)

	e.GET("/employeerisk/:username", h.GetEmployeeRisk)
	e.GET("/departmentrisk/:code", h.GetDepartmentRisk)

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
	}

	e.Logger.Fatal(e.StartServer(server))
}
