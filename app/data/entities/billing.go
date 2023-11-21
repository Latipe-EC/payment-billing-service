package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

const (
	// store level
	IndividualLevel = 1
	MallLevel       = 2

	//transaction type
	PayInWallet  = 1
	PayOutWallet = 2

	//message transaction
	PayInForOrderFinish      = "pay in for the order finish"
	PayOutForPurchasePayment = "pay out for the purchase payment"

	//wallet type
	IndividualWallet = 1
	StoreWallet      = 2
)

type CommissionLevel struct {
	CommissionLevelType string
	ID                  int             `gorm:"not null;autoIncrement;primaryKey;type:bigint" json:"level_id"`
	Status              bool            `gorm:"type:boolean;not null" json:"status"`
	FeePerOrder         float32         `gorm:"not null" json:"fee_per_order"`
	StoreBilling        []*StoreBilling `gorm:"constraint:OnUpdate:CASCADE;polymorphic:CommissionLevel;"`
}

func (CommissionLevel) TableName() string {
	return "commission_levels"
}

type StoreBilling struct {
	CommissionLevelType string
	ID                  uuid.UUID        `gorm:"not null;type:varchar(36);primary_key"  json:"id"`
	StoreID             string           `gorm:"type:varchar(250);not null" json:"owner_id"`
	CommissionLevel     *CommissionLevel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CommissionLevelID   int              `gorm:"type:varchar(250);not null"  json:"commission_level_id"`
	CreatedAt           time.Time        `gorm:"autoUpdateTime;type:datetime(6)" json:"created_at"`
	UpdatedAt           time.Time        `gorm:"autoUpdateTime;type:datetime(6)" json:"updated_at"`
}

func (o *StoreBilling) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New()
	return nil
}

func (StoreBilling) TableName() string {
	return "store_billings"
}

type EWallet struct {
	EWalletType         string
	ID                  uuid.UUID             `gorm:"not null;type:varchar(36);primary_key" json:"id"`
	AccountBalance      int                   `gorm:"not null;type:bigint" json:"account_balance"`
	OwnerID             string                `gorm:"not null;type:varchar(250)" json:"owner_id"`
	OwnerType           int                   `gorm:"not null;type:int" json:"owner_type"`
	IsActive            bool                  `gorm:"not null;type:boolean" json:"is_active"`
	CreatedAt           time.Time             `gorm:"autoUpdateTime;type:datetime(6)" json:"created_at"`
	UpdatedAt           time.Time             `gorm:"autoUpdateTime;type:datetime(6)" json:"updated_at"`
	EWalletTransactions []*EWalletTransaction `gorm:"constraint:OnUpdate:CASCADE;polymorphic:EWallet;"`
}

func (o *EWallet) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New()
	return nil
}

func (EWallet) TableName() string {
	return "e-wallets"
}

type EWalletTransaction struct {
	EWalletType string
	ID          int       `gorm:"not null;autoIncrement;primaryKey;type:bigint" json:"id"`
	EWallet     *EWallet  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EWalletID   string    `gorm:"not null;type:varchar(250)" json:"e_wallet_id"`
	Amount      int       `gorm:"not null;type:int" json:"amount"`
	Message     string    `gorm:"not null;type:varchar(250)" json:"message"`
	Type        int       `gorm:"not null;type:int" json:"type"`
	CreatedAt   time.Time `gorm:"autoUpdateTime;type:datetime(6)" json:"created_at"`
}

func (EWalletTransaction) TableName() string {
	return "e-wallet-transactions"
}
