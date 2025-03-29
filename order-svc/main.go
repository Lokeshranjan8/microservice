package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/Lokeshranjan8/order-svc/handlers"
	"github.com/Lokeshranjan8/order-svc/utils"
)


func main() {
	// Connect to DB
	utils.InitDB()

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.FetchAllOrders(w, r)
		case http.MethodPost:
			handlers.PlaceOrder(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/orders/user/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 5 {
			http.Error(w, "User ID not provided", http.StatusBadRequest)
			return
		}
		r.URL.Query().Add("user_id", parts[4])
		handlers.FetchUserOrders(w, r)
	})

	log.Println("Order service running on port 8002")
	log.Fatal(http.ListenAndServe(":8002", nil))
}
