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
