package controller

import (
	. "stencil-go/app/controller/base"
	"stencil-go/app/extend/util"
	"stencil-go/app/middleware/jwt"
	"stencil-go/app/model"
	"stencil-go/app/service"

	"github.com/kataras/iris/v12/mvc"
)

type UserVC struct {
	BaseVC

	User service.User
}

func (vc *UserVC) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/signIn", "SignIn")
	b.Handle("POST", "/signOut", "SignOut")
	b.Handle("POST", "/updatePwd", "UpdatePwd")
	b.Handle("GET POST", "/getUserInfo", "GetInfo")
	b.Handle("GET POST", "/getUserList", "GetList")
}

func (vc *UserVC) SignIn() {
	type Params struct {
		Account string `form:"account" validate:"required"`
		Pwd     string `form:"pwd" validate:"required"`
	}
	p := Params{}
	if err := vc.ReadBody(&p); err != nil {
		return
	}

	user, err := vc.User.LoginAccount(p.Account)
	if err != nil {
		vc.Failed(err.Error())
		return
	}
	if !util.CheckPwd(p.Pwd, user.Salt, user.Pwd) {
		vc.Failed("密码错误")
		return
	}

	vc.Success(Map{
		"user_id": user.ID,
		"token":   jwt.Sign(int64(user.ID), 0, "none"),
	})
}

func (vc *UserVC) SignOut() {
	jwt.Remove(vc.Ctx)
	vc.Success(nil)
}

func (vc *UserVC) GetInfo() {
	id := int(vc.TokenUser().UserId)
	// id = vc.Ctx.URLParamIntDefault("id", 1)
	info, err := vc.User.GetById(id)
	if err != nil {
		vc.Failed(err.Error())
		return
	}
	vc.Success(Map{
		"user_id":      info.ID,
		"nickname":     info.Nickname,
		"username":     info.Username,
		"telephone":    info.Telephone,
		"last_ip":      info.LastIP,
		"created_time": info.CreatedTime,
		"updated_time": info.UpdatedTime,
	})
}

func (vc *UserVC) UpdatePwd() {
	type Params struct {
		OldPwd string `form:"old_pwd" validate:"required"`
		NewPwd string `form:"new_pwd" validate:"required"`
	}
	p := Params{}
	if err := vc.ReadBody(&p); err != nil {
		return
	}

	id := int(vc.TokenUser().UserId)
	info, _ := vc.User.GetById(id)
	if !util.CheckPwd(p.OldPwd, info.Salt, info.Pwd) {
		vc.Failed("密码错误")
		return
	}
	if p.OldPwd == p.NewPwd {
		vc.Failed("新密码与旧密码一致")
		return
	}

	salt, pwd := util.EncryptPwd(p.NewPwd)
	vc.User.UpdatePwd(id, salt, pwd)
	vc.Success(nil)
}

func (vc *UserVC) GetList() {
	p := model.AdminParams{}
	if err := vc.ReadBody(&p); err != nil {
		return
	}

	list, count, err := vc.User.GetList(&p)
	if err != nil {
		vc.Failed(err.Error())
		return
	}

	vc.Success(Map{
		"total": count,
		"list":  list,
	})
}
