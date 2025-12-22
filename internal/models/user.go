package models

// 用户结构
type User struct {
	Uid       int64  `gorm:"primaryKey;autoIncrement" json:"uid"`
	Username  string `gorm:"unique;not null" json:"username"`
	Nickname  string `gorm:"not null" json:"nickname"`
	Password  string `gorm:"not null" json:"password"`
	Email     string `gorm:"unique;not null" json:"email"`
	Headicon  string `gorm:"default:null" json:"headicon"`
	Phone     string `gorm:"default:null" json:"phone"`
	CreatedAt string `gorm:"created_at" json:"created_at"`
	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
}

func (User) TableName() string {
	return "users"
}
