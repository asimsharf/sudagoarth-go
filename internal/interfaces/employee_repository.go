package interfaces

import "sudagoarth.com/internal/models"

type EmployeeInterface interface {
	Create(employee *models.Employee) error
	GetByID(id int) (*models.Employee, error)
	GetAll() ([]models.Employee, error)
	Update(employee *models.Employee) error
	Delete(id int) error
}
