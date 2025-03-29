package models

import "time"

type Payment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    string    `json:"user_id"`
	Item      string    `json:"item"`
	Amount    int       `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PaymentRequest struct {
	ID     string `json:"order_id"`
	Amount int    `json:"amount"`
}

type PaymentResponse struct {
	OrderID      string `json:"order_id"`
	Status       string `json:"status"`
	PaymentID    string `json:"payment_id,omitempty"`
	ErrorMessage string `json:"error,omitempty"`
}
