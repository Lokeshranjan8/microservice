package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Lokeshranjan8/notification-svc/models"
	"github.com/Lokeshranjan8/notification-svc/utils"
)

func SendNotification(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.Notification
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid payload"}`, http.StatusBadRequest)
		return
	}

	err := utils.EmailUser(req.Email, req.OrderID, req.Amount, req.Status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Response{Error: "Failed to send email"})
		return
	}

	json.NewEncoder(w).Encode(models.Response{
		Message: fmt.Sprintf("Email sent to %s regarding order %s", req.Email, req.OrderID),
	})
}
