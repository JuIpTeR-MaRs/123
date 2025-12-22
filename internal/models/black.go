package models

type BlackListUid struct {
	Id        int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Uid       int64  `gorm:"not null" json:"uid"`
	Reason    string `gorm:"default:null" json:"reason"`
	CreatedAt string `gorm:"created_at" json:"created_at"`
	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
}

func (BlackListUid) TableName() string { return "blacklist_uids" }
