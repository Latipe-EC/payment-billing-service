package service

import (
	"latipe-payment-billing-service/app/data/dto"
	"latipe-payment-billing-service/app/data/entities"
	"latipe-payment-billing-service/app/repository"
)

type BillingService struct {
	paymentRepos *repository.PaymentRepository
	billingRepos *repository.BillingRepository
}

func NewBillingService(paymentRepo *repository.PaymentRepository,
	billingRepo *repository.BillingRepository) BillingService {

	return BillingService{
		paymentRepos: paymentRepo,
		billingRepos: billingRepo,
	}
}

func (pm BillingService) CreateBillingFromMessage(message *dto.StoreMessage) error {

	billing := entities.StoreBilling{
		StoreID:           message.StoreId,
		CommissionLevelID: entities.IndividualLevel,
	}

	wallet := entities.EWallet{
		AccountBalance: 0,
		OwnerID:        message.StoreId,
		OwnerType:      entities.StoreWallet,
		IsActive:       true,
	}
	if err := pm.billingRepos.CreateStoreBilling(&billing, &wallet); err != nil {
		return err
	}

	return nil
}
