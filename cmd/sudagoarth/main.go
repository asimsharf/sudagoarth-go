package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sudagoarth.com/api/controllers"
	"sudagoarth.com/config"
	"sudagoarth.com/internal/models"
	"sudagoarth.com/internal/repositories"
	"sudagoarth.com/internal/services"
)

func main() {
	// Load the configuration
	cfg := config.LoadConfig()

	// Establish a database connection
	db, err := gorm.Open(mysql.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Perform auto-migration
	err = db.AutoMigrate(&models.Employee{})
	if err != nil {
		log.Fatalf("Could not migrate the database: %v", err)
	}

	// Set up the repository, service, and controller
	repo := repositories.NewEmployeeRepository(db)
	service := services.NewEmployeeService(repo)
	controller := &controllers.EmployeeController{Service: service}

	// Set up the router and endpoints
	router := mux.NewRouter()
	router.HandleFunc("/api/employees", controller.CreateEmployee).Methods("POST")
	router.HandleFunc("/api/employees/{id}", controller.GetEmployee).Methods("GET")
	router.HandleFunc("/api/employees", controller.GetAllEmployees).Methods("GET")
	router.HandleFunc("/api/employees/{id}", controller.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/api/employees/{id}", controller.DeleteEmployee).Methods("DELETE")

	// Start the server
	port := cfg.AppPort
	if port == "" {
		port = "3000"
	}

	log.Printf("Server started at :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
