package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Lokeshranjan8/payment-svc/models"
	"github.com/Lokeshranjan8/payment-svc/utils"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/paymentintent"
)
 
func init() {
	_ = godotenv.Load() 
}

func ProcessPayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	pi, err := paymentintent.New(&stripe.PaymentIntentParams{
		Amount:             stripe.Int64(int64(req.Amount * 100)),
		Currency:           stripe.String("usd"),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		PaymentMethod:      stripe.String("pm_card_visa"),
		Confirm:            stripe.Bool(true),
	})

	status := "pending"
	if err != nil {
		status = "failed"
	} else if pi.Status == stripe.PaymentIntentStatusSucceeded {
		status = "completed"
	}

	payment := models.Payment{
		UserID: req.ID,
		Item:   "Test Item", 
		Amount: req.Amount,
		Status: status,
	}
	utils.DB.Create(&payment)

	resp := models.PaymentResponse{
		OrderID:   req.ID,
		Status:    status,
		PaymentID: pi.ID,
	}
	if err != nil {
		resp.ErrorMessage = err.Error()
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
