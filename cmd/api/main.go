package main

import (
	"fmt"

	database "github.com/claytonssmint/crud-clay/adapters/db"
	"github.com/claytonssmint/crud-clay/adapters/web/handler"
	"github.com/claytonssmint/crud-clay/adapters/web/routes"
	"github.com/claytonssmint/crud-clay/application"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	sqlx "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sqlx.Connect("postgres", "user=postgres password=devclay18 dbname=db_employee sslmode=disable")

	if err = db.Ping(); err != nil {
		fmt.Println(err)

	}

	employeeRepository := database.NewEmployeeRepository(db)
	employeeService := application.NewEmployeeService(employeeRepository)
	employeeHandler := handler.NewEmployeeHandler(employeeService)

	e := echo.New()
	e.Use(middleware.CORS())

	routes.RegisterRoutes(e, employeeHandler)

	e.Logger.Fatal(e.Start(":8080"))

}
