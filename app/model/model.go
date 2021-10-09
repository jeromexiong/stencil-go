package model

// Users [...]
type Users struct {
	ID          uint32 `gorm:"primaryKey;column:id;type:int(10) unsigned;not null" json:"id"`
	Name        string `gorm:"column:name;type:varchar(64);not null;default:''" json:"name"` // 用户名
	Sex         bool   `gorm:"column:sex;type:tinyint(1);not null;default:0" json:"sex"`     // 0  未知| 1 男 | 2 女
	CreatedTime int64  `gorm:"column:created_time;type:bigint(20);not null;default:0" json:"created_time"`
	UpdatedTime int64  `gorm:"column:updated_time;type:bigint(20);not null;default:0" json:"updated_time"`
}
