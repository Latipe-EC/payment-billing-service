package dto

import (
	"time"
)

type PaymentResponseDetail struct {
	ID             string           `json:"id"`
	OrderID        string           ` json:"order_uuid"`
	Method         int              ` json:"method"`
	CreatedBy      string           `json:"user_id"`
	Amount         int              `json:"amount"`
	Status         int              `json:"status"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
	PaymentHistory []PaymentHistory `json:"payment_history"`
}

type PaymentHistory struct {
	OrderPaymentType string
	ID               int       `json:"id"`
	TransactionID    string    `json:"transaction_id,omitempty"`
	OrderPaymentID   string    `json:"order_payment_id"`
	UserID           string    `json:"user_id"`
	Message          string    `json:"message"`
	CreatedAt        time.Time `json:"created_at"`
}
