package main

import (
	"net/http"

	"github.com/Lokeshranjan8/payment-svc/handlers"
	"github.com/Lokeshranjan8/payment-svc/utils"
)

func main() {
	utils.InitDB()
	http.HandleFunc("/pay", handlers.ProcessPayment)
	http.ListenAndServe(":8082", nil)
}
