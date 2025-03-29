package models

type Notification struct {
	Email    string   `json:"email"`
	OrderID  string   `json:"order_id"`
	Amount   float64      `json:"amount"`
	Status   string   `json:"status"`
}

type Response struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`

}