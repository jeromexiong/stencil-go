package controller

import (
	. "stencil-go/app/controller/base"
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
	b.Handle("GET", "/getUserInfo", "GetUserInfo")
}

func (vc *UserVC) SignIn() {
	type Params struct {
		UserName  string `form:"name" validate:"required"`
		Age       int
		Cellphone string
		Tail      []string `form:"tail"`
	}
	p := Params{}

	if err := vc.ReadBody(&p); err != nil {
		return
	}
	res := make(map[string]interface{})
	res["name"] = p.UserName
	res["token"] = jwt.Sign(1, 0, "none")
	vc.Success(res)
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
