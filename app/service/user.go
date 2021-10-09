package service

import (
	"errors"
	"stencil-go/app/core"
	"stencil-go/app/model"
)

type User struct{}

func (u *User) GetById(id int) (model.Users, error) {
	user := model.Users{}
	result := core.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return user, errors.New("不存在的用户")
	}
	return user, nil
}
