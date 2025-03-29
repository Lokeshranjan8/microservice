package models

import "time"



type Order struct {
	OrderID   string    `json:"order_id,omitempty" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Customer  string    `json:"customer_id" gorm:"not null;index"` // Foreign key to user service
	Product   string    `json:"product" gorm:"type:varchar(255);not null"`
	Price     int       `json:"price" gorm:"not null;check:price >= 0"` // Non-negative price constraint
	Status     string    `json:"state,omitempty" gorm:"type:varchar(50);default:'pending'"` // Status of order
	CreatedOn time.Time `json:"created_on" gorm:"autoCreateTime"` // Timestamp on creation
	UpdatedOn time.Time `json:"updated_on" gorm:"autoUpdateTime"` // Timestamp on update
}

type Response struct {
	Message string      `json:"message"`
	Info    string      `json:"info"`
	Result  any `json:"result,omitempty"`
}

type PaymentStatus struct {
	OrderRef     string `json:"order_ref"`
	Status string `json:"status"`
	PaymentID    string `json:"payment_id,omitempty"`
	ErrorInfo    string `json:"error_info,omitempty"`
}

type UserInfo struct {
	UserID int    `json:"user_id"`
	Name   string `json:"full_name"`
	Email  string `json:"email_address"`
}

type UserReply struct {
	Message string   `json:"message"`
	Info    string   `json:"info"`
	Details UserInfo `json:"details"`
}

type MailPayload struct {
	Email string  `json:"recipient"`
	OrderID  string  `json:"order_ref"`
	Amount     float64 `json:"total"`
	Status     string  `json:"state"`
}

