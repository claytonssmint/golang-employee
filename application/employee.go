package application

import (
	database "github.com/claytonssmint/crud-clay/adapters/db"
	"github.com/claytonssmint/crud-clay/adapters/dto"
)

type EmployeeService interface {
	GetAllEmployees() ([]dto.Employee, error)
	GetEmplooyeeID(id int) (dto.Employee, error)
	CreateEmployee(employee dto.Employee) error
	UpdateEmployee(employee dto.Employee) error
	DeleteEmployee(id int) error
}

type employeeService struct {
	repository database.EmployeeRepository
}

func NewEmployeeService(repository database.EmployeeRepository) EmployeeService {
	return &employeeService{repository}
}

// CreateEmployee implements EmployeeService.
func (s *employeeService) CreateEmployee(employee dto.Employee) error {
	return s.repository.CreateEmployee(employee)
}

// GetAllEmployees implements EmployeeService.
func (s *employeeService) GetAllEmployees() ([]dto.Employee, error) {
	return s.repository.GetAllEmployees()

}

// GetEmplooyeeID implements EmployeeService.
func (s *employeeService) GetEmplooyeeID(id int) (dto.Employee, error) {
	return s.repository.GetEmplooyeeID(id)
}

// UpdateEmployee implements EmployeeService.
func (s *employeeService) UpdateEmployee(employee dto.Employee) error {
	return s.repository.UpdateEmployee(employee)
}

// DeleteEmployee implements EmployeeService.
func (s *employeeService) DeleteEmployee(id int) error {
	return s.repository.DeleteEmployee(id)
}
