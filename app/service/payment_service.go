package service

import (
	order "latipe-payment-billing-service/app/data/dto"
	"latipe-payment-billing-service/app/data/entities"
	"latipe-payment-billing-service/app/repository"
)

type PaymentService struct {
	paymentRepos *repository.PaymentRepository
}

func NewPaymentService(paymentRepo *repository.PaymentRepository) PaymentService {
	return PaymentService{paymentRepos: paymentRepo}
}

func (pm PaymentService) CreatePaymentOfOrder(message *order.OrderMessage) error {
	payment := entities.OrderPayment{
		Method:    message.PaymentMethod,
		CreatedBy: message.UserRequest.UserId,
		Amount:    message.Amount,
		Status:    entities.UNPAID,
		OrderID:   message.OrderUUID,
	}

	if err := pm.paymentRepos.CreatePayment(&payment); err != nil {
		return err
	}

	return nil
}
