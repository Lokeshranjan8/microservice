package utils


import (
	"fmt"
	"log"
	"github.com/Lokeshranjan8/user-svc/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB


func InitDB() {
    dsn := "host=localhost user=postgres password=pass dbname=user_db port=5432 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to DB:", err)
    }
    DB.AutoMigrate(&models.User{})
    fmt.Println("Connected to DB and Migrated")
}