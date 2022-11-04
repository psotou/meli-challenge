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

var db *sql.DB

func main() {
	var err error

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	dsn := fmt.Sprintf("root:%s@tcp(mysqldb:3306)/%s", os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQL_DATABASE"))
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	e.GET("/", getEmployee)
	e.POST("/employee", createEmployee)
	e.PUT("/employee/:username", updateEmployee)
	e.POST("/role", createRole)
	e.POST("/application", createApplication)
	e.POST("/dbaccess", createDbAccess)

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
	}

	e.Logger.Fatal(e.StartServer(server))
}
