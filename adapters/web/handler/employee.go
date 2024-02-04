package handler

import (
	"net/http"
	"strconv"

	"github.com/claytonssmint/crud-clay/adapters/dto"
	"github.com/claytonssmint/crud-clay/application"
	"github.com/labstack/echo"
)

type Response struct {
	Message string
	Status  bool
}

type EmployeeHandler interface {
	GetAllEmployees(c echo.Context) error
	GetEmployeeByID(c echo.Context) error
	CreateEmployee(c echo.Context) error
	UpdateEmployee(c echo.Context) error
	DeleteEmployee(c echo.Context) error
}

type employeeHandler struct {
	service application.EmployeeService
}

func NewEmployeeHandler(service application.EmployeeService) EmployeeHandler {
	return &employeeHandler{service}
}

func (h *employeeHandler) GetAllEmployees(c echo.Context) error {
	employees, err := h.service.GetAllEmployees()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "Error getting employees", Status: false})
	}
	return c.JSON(http.StatusOK, employees)
}

func (h *employeeHandler) GetEmployeeByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	employee, err := h.service.GetEmplooyeeID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "Error getting employee ID", Status: false})
	}
	return c.JSON(http.StatusOK, employee)
}

func (h *employeeHandler) CreateEmployee(c echo.Context) error {
	var employee dto.Employee
	if err := c.Bind(&employee); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Message: "Invalid request body",
			Status:  false,
		})
	}
	err := h.service.CreateEmployee(employee)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "Error creating employee", Status: false})
	}
	return c.JSON(http.StatusCreated, Response{Message: "Employee created successfully", Status: true})
}

func (h *employeeHandler) UpdateEmployee(c echo.Context) error {
	var employee dto.Employee
	if err := c.Bind(&employee); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Message: "Invalid request body",
			Status:  false,
		})
	}
	id, _ := strconv.Atoi(c.Param("id"))
	employee.Id = id
	err := h.service.UpdateEmployee(employee)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "Error updating employee", Status: false})
	}
	return c.JSON(http.StatusOK, Response{Message: "Employee updated successfully", Status: true})
}

func (h *employeeHandler) DeleteEmployee(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.DeleteEmployee(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "Error deleting employee", Status: false})
	}
	return c.JSON(http.StatusOK, Response{Message: "Employee deleted successfully", Status: true})
}
