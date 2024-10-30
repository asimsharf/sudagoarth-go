package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sudagoarth.com/config"
	"sudagoarth.com/internal/models"
)

func ConnectMySQL() (*gorm.DB, string) {
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

	return db, cfg.AppPort
}
