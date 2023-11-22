package dto

type GetPaymentByOrderIDRequest struct {
	BaseHeader
	UserID    string
	OrderUUID string `json:"order_uuid" validate:"required"`
}

type GetPaymentOrderIdResponse struct {
	PaymentResponseDetail
}
