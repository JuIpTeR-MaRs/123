package models

type Admin struct {
	AdminId   int64  `gorm:"primaryKey;autoIncrement" json:"admin_id"`
	Username  string `gorm:"unique;not null" json:"username"`
	Password  string `gorm:"not null" json:"password"`
	Email     string `gorm:"unique;not null" json:"email"`
	CreatedAt string `gorm:"created_at" json:"created_at"`
	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
}

func (Admin) TableName() string { return "admins" }
