package routes

import (
	"github.com/claytonssmint/crud-clay/adapters/web/handler"
	"github.com/labstack/echo"
)

func RegisterRoutes(e *echo.Echo, handler handler.EmployeeHandler) {
	e.GET("/users", handler.GetAllEmployees)
	e.GET("users/:id", handler.GetEmployeeByID)
	e.POST("users", handler.CreateEmployee)
	e.PUT("/users/update/:id", handler.UpdateEmployee)
	e.DELETE("/users/delete/:id", handler.DeleteEmployee)
}
