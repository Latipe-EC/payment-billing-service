package repository

import (
	gormF "gorm.io/gorm"
	"latipe-payment-billing-service/app/data/entities"
	"latipe-payment-billing-service/app/repository/gorm"
)

type BillingRepository struct {
	client gorm.Gorm
}

func NewBillingRepository(client gorm.Gorm) BillingRepository {
	// auto migrate
	err := client.DB().AutoMigrate(
		&entities.CommissionLevel{},
		&entities.StoreBilling{},
		&entities.EWallet{},
		&entities.EWalletTransaction{},
	)
	if err != nil {
		panic(err)
	}
	return BillingRepository{
		client: client,
	}
}

func (pm BillingRepository) CreateStoreBilling(entity *entities.StoreBilling, wallet *entities.EWallet) error {
	err := pm.client.Transaction(func(tx *gormF.DB) error {
		if err := tx.Create(&entity).Error; err != nil {
			return err
		}

		if err := tx.Create(&wallet).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
