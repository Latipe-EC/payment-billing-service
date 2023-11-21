package dto

type BaseHeader struct {
	BearerToken string `reqHeader:"Authorization"`
}

type OrderMessage struct {
	Header        BaseHeader
	UserRequest   UserRequest       `json:"user_request"`
	OrderUUID     string            `json:"order_uuid"`
	Amount        int               `json:"amount" validate:"required"`
	ShippingCost  int               ` json:"shipping_cost"`
	Discount      int               `json:"discount" validate:"required"`
	SubTotal      int               `json:"sub_total" validate:"required"`
	PaymentMethod int               `json:"payment_method" validate:"required"`
	Vouchers      []string          `json:"vouchers"`
	Address       OrderAddress      `json:"address" validate:"required"`
	Delivery      Delivery          `json:"delivery" validate:"required"`
	OrderItems    []OrderItemsCache `json:"order_items" validate:"required"`
}

type UserRequest struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
}
type OrderItemsCache struct {
	CartItemId  string      `json:"cart_item_id"`
	ProductItem ProductItem `json:"product_item"`
}

type ProductItem struct {
	ProductID   string `gorm:"not null;type:varchar(255)" json:"product_id"`
	ProductName string `gorm:"not null;type:varchar(255)" json:"product_name"`
	StoreID     string `gorm:"not null;type:varchar(255)" json:"store_id"`
	OptionID    string `gorm:"not null;type:varchar(250)" json:"option_id" `
	Quantity    int    `gorm:"not null;type:int" json:"quantity"`
	Price       int    `gorm:"not null;type:bigint" json:"price"`
	NetPrice    int    `gorm:"not null;type:bigint" json:"net_price"`
}

type OrderAddress struct {
	AddressId       string `json:"address_id"`
	ShippingName    string `json:"shipping_name" validate:"required"`
	ShippingPhone   string `json:"shipping_phone" validate:"required"`
	ShippingAddress string `json:"shipping_address" validate:"required"`
}

type Delivery struct {
	DeliveryId    string `json:"delivery_id" validate:"required"`
	Name          string `json:"name" validate:"required"`
	Cost          int    `json:"cost" validate:"required"`
	ReceivingDate string `json:"receiving_date" validate:"required"`
}
