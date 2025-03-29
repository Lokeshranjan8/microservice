package utils

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/Lokeshranjan8/payment-svc/models"

)

var DB *gorm.DB

func InitDB(){
	dsn :=  "host=localhost user=postgres password=pass dbname=payment_db port=5432 sslmode=disable"
    var err error
	DB , err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err !=nil{
		log.Fatal("db not connected ",err)
	}
	fmt.Println("connected to payment_db")
	DB.AutoMigrate(&models.Payment{})
}