package repository

import (
	"latipe-payment-billing-service/app/data/entities"
	"latipe-payment-billing-service/app/repository/gorm"
)

type PaymentRepository struct {
	client gorm.Gorm
}

func NewPaymentRepository(client gorm.Gorm) PaymentRepository {
	// auto migrate
	err := client.DB().AutoMigrate(
		&entities.OrderPayment{},
		&entities.PaymentHistory{},
	)
	if err != nil {
		panic(err)
	}
	return PaymentRepository{
		client: client,
	}
}

func (pm PaymentRepository) CreatePayment(entity *entities.OrderPayment) error {
	err := pm.client.DB().Create(&entity).Error
	if err != nil {
		return err
	}

	return nil
}

func (pm PaymentRepository) UpdatePaymentStatus(payment entities.OrderPayment) error {
	err := pm.client.DB().Updates(&payment).Error
	if err != nil {
		return err
	}

	return nil
}

func (pm PaymentRepository) GetPaymentById(paymentId string) (*entities.OrderPayment, error) {
	var entity entities.OrderPayment

	err := pm.client.DB().Preload("PaymentHistory").
		Find(&entity, "id=?", paymentId).Error
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (pm PaymentRepository) GetPaymentByOrderId(orderId string) (*entities.OrderPayment, error) {
	var entity entities.OrderPayment

	err := pm.client.DB().Preload("PaymentHistory").
		Find(&entity, "order_payments.order_id=?", orderId).Error
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (pm PaymentRepository) CreatePaymentLog(log *entities.PaymentHistory) error {

	err := pm.client.DB().Create(&log).Error
	if err != nil {
		return err
	}

	return nil
}
