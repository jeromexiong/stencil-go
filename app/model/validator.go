package model

import (
	"time"

	"gorm.io/gorm"
)

// 有`UpdatedTime`字段的模型需要继承该模型 [gorm Hook](https://gorm.io/zh_CN/docs/hooks.html)
type Model struct{}

func (u *Model) BeforeCreate(db *gorm.DB) (err error) {
	db.Statement.SetColumn("CreatedTime", time.Now().Unix())
	db.Statement.SetColumn("UpdatedTime", time.Now().Unix())
	return
}

func (u *Model) BeforeUpdate(db *gorm.DB) (err error) {
	db.Statement.SetColumn("UpdatedTime", time.Now().Unix())
	return
}

// 分页查询参数
type ListQuery struct {
	Page     int `json:"page" form:"page" url:"page"`
	PageSize int `json:"page_size" form:"page_size" url:"page_size"`
}

type AdminParams struct {
	ListQuery
	Keywords string `json:"keywords" form:"keywords"`
}
