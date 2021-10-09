package controller

import (
	. "stencil-go/app/controller/base"
	"stencil-go/app/model"
	"stencil-go/app/service"

	"github.com/kataras/iris/v12/mvc"
)

type TodoVC struct {
	BaseVC

	Service service.Todo
}

func (vc *TodoVC) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/", "Post")
	b.Handle("GET", "/", "Get")
}

func (vc *TodoVC) Get() {
	vc.Ctx.JSON(vc.Service.Get("0"))
}

func (vc *TodoVC) Post() {
	var items []model.Item
	vc.Ctx.ReadJSON(&items)
	if err := vc.Service.Save("0", items); err != nil {
		vc.Ctx.JSON(PostItemResponse{Success: false})
		return
	}
	vc.Ctx.JSON(PostItemResponse{Success: true})
}

type PostItemResponse struct {
	Success bool `json:"success"`
}
