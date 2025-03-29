package utils

import (
	"fmt"
	"log"

	"github.com/Lokeshranjan8/order-svc/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	dbURL := "host=localhost user=postgres password=pass dbname=order_db port=5432 sslmode=disable"

	conn, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	conn.AutoMigrate(&models.Order{})
    fmt.Println("Connected to DB and Migrated")
	
}
