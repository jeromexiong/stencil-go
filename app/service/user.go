package service

import (
	"errors"
	"stencil-go/app/core"
	"stencil-go/app/model"
)

type User struct{}

func (u *User) LoginAccount(username string) (model.Admin, error) {
	user := model.Admin{}
	result := core.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return user, errors.New("不存在的用户")
	}
	return user, nil
}

func (u *User) GetById(id int) (model.Admin, error) {
	user := model.Admin{}
	result := core.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return user, errors.New("不存在的用户")
	}
	return user, nil
}

func (u *User) UpdatePwd(id int, salt string, pwd string) {
	update := &model.Admin{
		ID:   uint32(id),
		Salt: salt,
		Pwd:  pwd,
	}
	core.DB.Model(update).Updates(update)
}

func (u *User) GetList(params *model.AdminParams) (res []model.Admin, count int64, err error) {
	if len(params.Keywords) > 0 {
		core.DB.Where("username like %?", params.Keywords)
	}
	if params.Page > 0 && params.PageSize > 0 {
		core.DB.Limit(params.PageSize).Offset((params.Page - 1) * params.PageSize)
	}
	core.DB.Select(
		"id as user_id",
		"nickname",
		"username",
		"telephone",
		"last_ip",
		"created_time",
		"updated_time",
	)
	if err = core.DB.Find(&res).Count(&count).Error; err != nil {
		return
	}
	return
}
