package models

type ProductReview struct {
	Id        int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int64  `gorm:"not null" json:"user_id"`
	ProductID int64  `gorm:"not null" json:"product_id"`
	Rating    int    `gorm:"not null" json:"rating"`
	Comment   string `gorm:"type:text;not null" json:"comment"`
	Images    string `gorm:"type:text" json:"images"`
}
