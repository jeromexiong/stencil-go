package controller

import (
	. "stencil-go/app/controller/base"
	"stencil-go/app/extend/util"
	"stencil-go/app/middleware/jwt"
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
	b.Handle("GET POST", "/getUserInfo", "GetUserInfo")
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

	res := make(map[string]interface{})
	res["id"] = user.ID
	res["token"] = jwt.Sign(int64(user.ID), 0, "none")
	vc.Success(res)
}

func (vc *UserVC) SignOut() {
	jwt.Remove(vc.Ctx)
	vc.Success(nil)
}

func (vc *UserVC) GetUserInfo() {
	id := vc.Ctx.URLParamIntDefault("id", 1)
	user, err := vc.User.GetById(id)
	if err != nil {
		vc.Failed(err.Error())
		return
	}
	vc.Success(user)
}
