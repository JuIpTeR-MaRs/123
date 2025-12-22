package models

type Order struct {
	Id              int64   `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID          int64   `gorm:"not null" json:"user_id"`
	OrderNo         string  `gorm:"not null;unique" json:"order_no"`
	ProductID       int64   `gorm:"not null" json:"product_id"`
	ProductName     string  `gorm:"not null" json:"product_name"`
	Quantity        int64   `gorm:"not null" json:"quantity"`
	TotalAmount     float64 `gorm:"not null" json:"total_amount"`
	PayAmount       float64 `gorm:"not null" json:"pay_amount"`
	UnitPrice       float64 `gorm:"not null" json:"unit_price"`
	ReceiverName    string  `gorm:"not null" json:"receiver_name"`
	ReceiverPhone   string  `gorm:"not null" json:"receiver_phone"`
	ReceiverAddress string  `gorm:"not null" json:"receiver_address"`
	Status          string  `gorm:"not null;default:'0'" json:"status"` //0:待付款 1:待发货 2:待收货 3:已完成
}
