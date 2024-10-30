package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"sudagoarth.com/api/controllers"
	"sudagoarth.com/internal/repositories"
	"sudagoarth.com/internal/services"
	"sudagoarth.com/pkg/db"
)

func main() {
	db, port := db.ConnectMySQL()

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

	if port == "" {
		port = "3000"
	}

	log.Printf("Server started at :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
