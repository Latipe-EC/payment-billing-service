package service

import (
	"errors"
	order "latipe-payment-billing-service/app/data/dto"
	"latipe-payment-billing-service/app/data/entities"
	"latipe-payment-billing-service/app/pkg/mapper"
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

func (pm PaymentService) PaymentCompleteStatus(dto *order.CompleteOrderPaymentStatus) error {
	payment, err := pm.paymentRepos.GetPaymentByOrderId(dto.OrderID)
	if err != nil {
		return err
	}

	if payment.Status == entities.PAID {
		return errors.New("order was paid")
	}

	payment.Status = entities.PAID

	log := entities.PaymentHistory{
		OrderPaymentID: payment.ID.String(),
		UserID:         dto.UserID,
		Message:        "the order was paid",
	}

	if err := pm.paymentRepos.UpdatePaymentStatus(*payment); err != nil {
		return err
	}
	if err := pm.paymentRepos.CreatePaymentLog(&log); err != nil {
		return err
	}

	return nil
}

func (pm PaymentService) GetPaymentOfOrder(dto *order.GetPaymentByOrderIDRequest) (*order.GetPaymentOrderIdResponse, error) {
	payment, err := pm.paymentRepos.GetPaymentByOrderId(dto.OrderUUID)
	if err != nil {
		return nil, err
	}

	resp := order.GetPaymentOrderIdResponse{}

	if err := mapper.BindingStruct(payment, &resp); err != nil {
		return nil, err
	}
	return &resp, nil

}
