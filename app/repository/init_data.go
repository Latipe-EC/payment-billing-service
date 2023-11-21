package repository

import (
	gormF "gorm.io/gorm"
	"latipe-payment-billing-service/app/data/entities"
	"latipe-payment-billing-service/app/repository/gorm"
)

func commissionLevelData() []*entities.CommissionLevel {
	var data []*entities.CommissionLevel

	individual := entities.CommissionLevel{
		Status:      true,
		FeePerOrder: 0.05,
	}
	data = append(data, &individual)

	mall := entities.CommissionLevel{
		Status:      true,
		FeePerOrder: 0.03,
	}
	data = append(data, &mall)

	return data

}

func InitCommissionLevelData(db gorm.Gorm) error {
	err := db.DB().AutoMigrate(
		&entities.CommissionLevel{},
		&entities.StoreBilling{},
		&entities.EWallet{},
		&entities.EWalletTransaction{},
	)

	var totalRecord int64

	dataInit := commissionLevelData()

	err = db.DB().Model(&entities.CommissionLevel{}).Count(&totalRecord).Error
	if err != nil {
		return err
	}

	if totalRecord == 0 {

		if err := db.DB().Transaction(func(tx *gormF.DB) error {
			for _, i := range dataInit {
				if err := db.DB().Create(&i).Error; err != nil {
					return nil
				}
			}
			return nil
		}); err != nil {
			return err
		}
	}

	return nil
}
