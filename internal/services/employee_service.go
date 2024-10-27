package services

import (
	"sudagoarth.com/internal/models"
	"sudagoarth.com/internal/repositories"
)

type EmployeeService struct {
	Repo repositories.EmployeeRepository
}

func NewEmployeeService(repo repositories.EmployeeRepository) *EmployeeService {
	return &EmployeeService{Repo: repo}
}

func (s *EmployeeService) CreateEmployee(employee *models.Employee) error {
	return s.Repo.Create(employee)
}

func (s *EmployeeService) GetEmployeeByID(id uint) (*models.Employee, error) {
	return s.Repo.GetByID(id)
}

func (s *EmployeeService) GetAllEmployees() ([]models.Employee, error) {
	return s.Repo.GetAll()
}

func (s *EmployeeService) UpdateEmployee(employee *models.Employee) error {
	return s.Repo.Update(employee)
}

func (s *EmployeeService) DeleteEmployee(id uint) error {
	return s.Repo.Delete(id)
}
