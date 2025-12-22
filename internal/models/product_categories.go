package models

type ProductCategory struct {
	Id          int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"default:null" json:"description"`
	ParentId    int64  `gorm:"default:null" json:"parent_id"`
	CreatedAt   string `gorm:"created_at" json:"created_at"`
	UpdatedAt   string `gorm:"updated_at" json:"updated_at"`
	DeletedAt   string `gorm:"deleted_at" json:"deleted_at"`
}
