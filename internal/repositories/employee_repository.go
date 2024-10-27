package repositories

import (
	"sudagoarth.com/internal/models"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(employee *models.Employee) error
	GetByID(id uint) (*models.Employee, error)
	GetAll() ([]models.Employee, error)
	Update(employee *models.Employee) error
	Delete(id uint) error
}

type GORMEmployeeRepository struct {
	DB *gorm.DB
}

func NewGORMEmployeeRepository(db *gorm.DB) *GORMEmployeeRepository {
	return &GORMEmployeeRepository{DB: db}
}

func (r *GORMEmployeeRepository) Create(employee *models.Employee) error {
	return r.DB.Create(employee).Error
}

func (r *GORMEmployeeRepository) GetByID(id uint) (*models.Employee, error) {
	var employee models.Employee
	err := r.DB.First(&employee, id).Error
	return &employee, err
}

func (r *GORMEmployeeRepository) GetAll() ([]models.Employee, error) {
	var employees []models.Employee
	err := r.DB.Find(&employees).Error
	return employees, err
}

func (r *GORMEmployeeRepository) Update(employee *models.Employee) error {
	return r.DB.Save(employee).Error
}

func (r *GORMEmployeeRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Employee{}, id).Error
}
