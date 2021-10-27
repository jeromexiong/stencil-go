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
