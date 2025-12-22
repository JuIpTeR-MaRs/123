package models

type CartItem struct {
	Id        int64 `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int64 `gorm:"not null" json:"user_id"`
	ProductID int64 `gorm:"not null" json:"product_id"`
	Quantity  int64 `gorm:"not null" json:"quantity"`
}
