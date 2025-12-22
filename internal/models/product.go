package models

import (
	"gorm.io/datatypes"
)

type Product struct {
	Id          int64                       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string                      `gorm:"not null" json:"name"`
	CategoryId  string                      `gorm:"not null" json:"category_id"`
	Description string                      `gorm:"default:null" json:"description"`
	Price       float64                     `gorm:"not null" json:"price"`
	Sku         string                      `gorm:"unique;not null" json:"sku"`
	Stock       int64                       `gorm:"not null" json:"stock"`
	CoverImage  string                      `gorm:"default:null" json:"cover_image"`
	Images      datatypes.JSONSlice[string] `gorm:"default:null" json:"images"`
	IsOnSale    bool                        `gorm:"default:true" json:"is_on_sale"`
	CreatedAt   string                      `gorm:"created_at" json:"created_at"`
	UpdatedAt   string                      `gorm:"updated_at" json:"updated_at"`
	DeletedAt   string                      `gorm:"deleted_at" json:"deleted_at"`
}
