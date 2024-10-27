// internal/controllers/employee_controller.go
package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"sudagoarth.com/internal/exceptions"
	"sudagoarth.com/internal/models"
	"sudagoarth.com/internal/services"
)

type EmployeeController struct {
	Service *services.EmployeeService
}

// CreateEmployee handles the creation of a new employee
func (c *EmployeeController) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		exceptions.SendResponse(w, http.StatusBadRequest, "Invalid request payload", nil, err.Error())
		return
	}

	if err := c.Service.CreateEmployee(&employee); err != nil {
		exceptions.SendResponse(w, http.StatusInternalServerError, "Could not create employee", nil, err.Error())
		return
	}

	exceptions.SendResponse(w, http.StatusCreated, "Employee created successfully", employee, "")
}

// GetEmployee handles the retrieval of a single employee by ID
func (c *EmployeeController) GetEmployee(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		exceptions.SendResponse(w, http.StatusBadRequest, "Invalid employee ID", nil, err.Error())
		return
	}

	employee, err := c.Service.GetEmployeeByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			exceptions.SendResponse(w, http.StatusNotFound, "Employee not found", nil, "")
		} else {
			exceptions.SendResponse(w, http.StatusInternalServerError, "Could not retrieve employee", nil, err.Error())
		}
		return
	}

	exceptions.SendResponse(w, http.StatusOK, "Employee retrieved successfully", employee, "")
}

// GetAllEmployees handles the retrieval of all employees
func (c *EmployeeController) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := c.Service.GetAllEmployees()
	if err != nil {
		exceptions.SendResponse(w, http.StatusInternalServerError, "Could not retrieve employees", nil, err.Error())
		return
	}

	exceptions.SendResponse(w, http.StatusOK, "Employees retrieved successfully", employees, "")
}

// UpdateEmployee handles the updating of an existing employee
func (c *EmployeeController) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		exceptions.SendResponse(w, http.StatusBadRequest, "Invalid employee ID", nil, err.Error())
		return
	}

	var employee models.Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		exceptions.SendResponse(w, http.StatusBadRequest, "Invalid request payload", nil, err.Error())
		return
	}
	employee.ID = uint(id)

	if err := c.Service.UpdateEmployee(&employee); err != nil {
		if err == gorm.ErrRecordNotFound {
			exceptions.SendResponse(w, http.StatusNotFound, "Employee not found", nil, "")
		} else {
			exceptions.SendResponse(w, http.StatusInternalServerError, "Could not update employee", nil, err.Error())
		}
		return
	}

	exceptions.SendResponse(w, http.StatusOK, "Employee updated successfully", employee, "")
}

// DeleteEmployee handles the deletion of an employee by ID
func (c *EmployeeController) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		exceptions.SendResponse(w, http.StatusBadRequest, "Invalid employee ID", nil, err.Error())
		return
	}

	if err := c.Service.DeleteEmployee(uint(id)); err != nil {
		if err == gorm.ErrRecordNotFound {
			exceptions.SendResponse(w, http.StatusNotFound, "Employee not found", nil, "")
		} else {
			exceptions.SendResponse(w, http.StatusInternalServerError, "Could not delete employee", nil, err.Error())
		}
		return
	}

	exceptions.SendResponse(w, http.StatusNoContent, "Employee deleted successfully", nil, "")
}
