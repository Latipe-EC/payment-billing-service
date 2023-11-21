package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

const (
	//payment status
	UNPAID = 0
	PAID   = 1
)

type OrderPayment struct {
	OrderPaymentType string
	ID               uuid.UUID         `gorm:"not null;type:varchar(36);primary_key"`
	OrderID          string            `gorm:"not null;type:varchar(250)" json:"order_id"`
	Method           int               `gorm:"not null;type:int" json:"method"`
	CreatedBy        string            `gorm:"not null;type:varchar(250)" json:"user_id"`
	Amount           int               `gorm:"not null;type:bigint" json:"amount"`
	Status           int               `gorm:"not null;type:int" json:"status"`
	CreatedAt        time.Time         `gorm:"autoUpdateTime;type:datetime(6)" json:"created_at"`
	UpdatedAt        time.Time         `gorm:"autoUpdateTime;type:datetime(6)" json:"updated_at"`
	PaymentHistory   []*PaymentHistory `gorm:"constraint:OnUpdate:CASCADE;polymorphic:OrderPayment;"`
}

func (o *OrderPayment) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New()
	return nil
}

func (OrderPayment) TableName() string {
	return "order_payments"
}

type PaymentHistory struct {
	OrderPaymentType string
	ID               int           `gorm:"not null;autoIncrement;primaryKey;type:bigint" json:"id"`
	TransactionID    string        `gorm:"type:varchar(250)" json:"transaction_id"`
	OrderPaymentID   string        `gorm:"not null;type:varchar(250)" json:"order_payment_id"`
	OrderPayment     *OrderPayment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID           string        `gorm:"not null;type:varchar(250)" json:"user_id"`
	Message          string        `gorm:"not null;type:varchar(250)" json:"message"`
	CreatedAt        time.Time     `gorm:"autoUpdateTime;type:datetime(6)" json:"created_at"`
}

func (PaymentHistory) TableName() string {
	return "payment_logs"
}
