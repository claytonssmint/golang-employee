package db

import (
	"github.com/claytonssmint/crud-clay/adapters/dto"
	sqlx "github.com/jmoiron/sqlx"
)

type EmployeeRepository interface {
	GetAllEmployees() ([]dto.Employee, error)
	GetEmplooyeeID(id int) (dto.Employee, error)
	CreateEmployee(employee dto.Employee) error
	UpdateEmployee(employee dto.Employee) error
	DeleteEmployee(id int) error
}

type employeeRepository struct {
	db *sqlx.DB
}

func NewEmployeeRepository(db *sqlx.DB) EmployeeRepository {
	return &employeeRepository{db}
}

// GetAllEmployees implements EmployeeRepository.
func (r *employeeRepository) GetAllEmployees() ([]dto.Employee, error) {
	var employees []dto.Employee
	err := r.db.Select(&employees, "SELECT * FROM users")
	return employees, err

}

// GetEmplooyeeID implements EmployeeRepository.
func (r *employeeRepository) GetEmplooyeeID(id int) (dto.Employee, error) {
	var employee dto.Employee
	err := r.db.Get(&employee, "SELECT * FROM users WHERE id = $1", id)
	return employee, err

}

// CreateEmployee implements EmployeeRepository.
func (r *employeeRepository) CreateEmployee(employee dto.Employee) error {
	_, err := r.db.NamedExec("INSERT INTO users(name, phone, address) values (:name, :phone, :address)", employee)
	return err
}

// UpdateEmployee implements EmployeeRepository.
func (r *employeeRepository) UpdateEmployee(employee dto.Employee) error {
	_, err := r.db.NamedExec("UPDATE users SET name= :name, phone= :phone, address= :address WHERE id= :id", employee)
	return err
}

// DeleteEmployee implements EmployeeRepository.
func (r *employeeRepository) DeleteEmployee(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
