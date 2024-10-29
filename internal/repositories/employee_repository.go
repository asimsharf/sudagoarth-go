package repositories

import (
	"gorm.io/gorm"
	"sudagoarth.com/internal/models"
)

type EmployeeRepository struct {
	DB *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

func (r *EmployeeRepository) Create(employee *models.Employee) error {
	return r.DB.Create(employee).Error
}

func (r *EmployeeRepository) GetByID(id uint) (*models.Employee, error) {
	var employee models.Employee
	err := r.DB.First(&employee, id).Error
	return &employee, err
}

func (r *EmployeeRepository) GetAll() ([]models.Employee, error) {
	var employees []models.Employee
	err := r.DB.Find(&employees).Error
	return employees, err
}

func (r *EmployeeRepository) Update(employee *models.Employee) error {
	return r.DB.Save(employee).Error
}

func (r *EmployeeRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Employee{}, id).Error
}
