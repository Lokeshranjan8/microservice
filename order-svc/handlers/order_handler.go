package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/Lokeshranjan8/order-svc/utils"
	"github.com/Lokeshranjan8/order-svc/models"
)

func FetchAllOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var orderList []models.Order
	if err := utils.DB.Find(&orderList).Error; err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	resp := models.Response{
		Message: "Orders Retrieved",
		Info:    "Fetched all available orders from DB",
		Result:  orderList,
	}
	json.NewEncoder(w).Encode(resp)
}

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newOrder models.Order
	if err := json.NewDecoder(r.Body).Decode(&newOrder); err != nil {
		http.Error(w, `{"error": "Invalid JSON data"}`, http.StatusBadRequest)
		return
	}
	newOrder.Status = "pending"

	paymentBody, err := json.Marshal(newOrder)
	if err != nil {
		http.Error(w, `{"error": "Failed to encode payment data"}`, http.StatusInternalServerError)
		return
	}

	resp, err := http.Post("http://localhost:8003/api/processpayment", "application/json", bytes.NewBuffer(paymentBody))
	if err != nil || resp.StatusCode != http.StatusOK {
		http.Error(w, `{"error": "Payment service failure"}`, http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	var payResp models.PaymentStatus
	if err := json.NewDecoder(resp.Body).Decode(&payResp); err != nil {
		http.Error(w, `{"error": "Unable to decode payment response"}`, http.StatusInternalServerError)
		return
	}

	switch payResp.Status {
	case "completed":
		newOrder.Status = "completed"
	case "pending":
		newOrder.Status = "pending"
	default:
		http.Error(w, `{"message": "Payment unsuccessful"}`, http.StatusBadRequest)
		return
	}

	if err := utils.DB.Create(&newOrder).Error; err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	// Fetch user info
	userResp, err := http.Get(fmt.Sprintf("http://localhost:8001/api/getuser/%s", newOrder.OrderID))
	if err != nil || userResp.StatusCode != http.StatusOK {
		http.Error(w, `{"error": "User service error"}`, http.StatusBadGateway)
		return
	}
	defer userResp.Body.Close()

	var userData models.UserInfo
	if err := json.NewDecoder(userResp.Body).Decode(&userData); err != nil {
		http.Error(w, `{"error": "User data decode failed"}`, http.StatusInternalServerError)
		return
	}

	notif := models.MailPayload{
		Email:   userData.Email,
		OrderID: newOrder.OrderID,
		Amount:  float64(newOrder.Price),
		Status:  newOrder.Status,
	}

	notifJSON, err := json.Marshal(notif)
	if err != nil {
		http.Error(w, `{"error": "Failed to encode notification payload"}`, http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:8004/api/notify", bytes.NewBuffer(notifJSON))
	if err != nil {
		http.Error(w, `{"error": "Notification request error"}`, http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	notifResp, err := client.Do(req)
	if err != nil || notifResp.StatusCode != http.StatusOK {
		http.Error(w, `{"error": "Notification failed"}`, http.StatusBadGateway)
		return
	}

	json.NewEncoder(w).Encode(models.Response{
		Message: "Order Created",
		Info:    "Order placed and user notified successfully",
		Result:  newOrder,
	})
}


func FetchUserOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID := mux.Vars(r)["user_id"]
	var userOrders []models.Order

	result := utils.DB.Where("user_id = ?", userID).Find(&userOrders)
	if result.Error != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, result.Error.Error()), http.StatusInternalServerError)
		return
	}
	if len(userOrders) == 0 {
		http.Error(w, `{"message": "No orders found for this user"}`, http.StatusNotFound)
		return
	}

	resp := models.Response{
		Message: "User Orders",
		Info:    fmt.Sprintf("Found %d order(s) for user %s", len(userOrders), userID),
		Result:  userOrders,
	}
	json.NewEncoder(w).Encode(resp)
}
