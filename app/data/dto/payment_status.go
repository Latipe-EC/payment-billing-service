package dto

type CompleteOrderPaymentStatus struct {
	BaseHeader
	UserID  string
	OrderID string `json:"order_id"`
}
