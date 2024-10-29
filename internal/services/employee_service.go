package services

import (
	"sudagoarth.com/internal/interfaces"
	"sudagoarth.com/internal/models"
)

type EmployeeService struct {
	Face interfaces.EmployeeInterface
}

func NewEmployeeService(face interfaces.EmployeeInterface) *EmployeeService {
	return &EmployeeService{Face: face}
}

func (s *EmployeeService) CreateEmployee(employee *models.Employee) error {
	return s.Face.Create(employee)
}

func (s *EmployeeService) GetEmployeeByID(id uint) (*models.Employee, error) {
	return s.Face.GetByID(id)
}

func (s *EmployeeService) GetAllEmployees() ([]models.Employee, error) {
	return s.Face.GetAll()
}

func (s *EmployeeService) UpdateEmployee(employee *models.Employee) error {
	return s.Face.Update(employee)
}

func (s *EmployeeService) DeleteEmployee(id uint) error {
	return s.Face.Delete(id)
}
