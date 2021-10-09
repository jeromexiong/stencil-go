package router

import (
	"stencil-go/app/controller"
	"stencil-go/app/controller/base"
	"stencil-go/app/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/websocket"
)

// 初始化路由
func New(app *iris.Application) {
	base.ErrorHandle(app)
	mvc.Configure(app.Party("/api"), interfaceMvc)
	mvc.Configure(app.Party("/todos"), interfaceMvcTodo)
}

func interfaceMvc(app *mvc.Application) {
	app.Party("/user").Handle(new(controller.UserVC))
}

func interfaceMvcTodo(app *mvc.Application) {
	app.Register(
		service.NewMemoryTodo(),
	)
	app.Handle(new(controller.TodoVC))

	todosWebsocketApp := app.Party("/sync")
	todosWebsocketApp.HandleWebsocket(new(controller.WebsocketVC)).
		SetNamespace("todos")

	upgrader := websocket.DefaultGorillaUpgrader
	websocketServer := websocket.New(upgrader, todosWebsocketApp)
	idGenerator := func(ctx iris.Context) string {
		// id := sess.Start(ctx).ID()
		return "0"
	}
	todosWebsocketApp.Router.Get("/", websocket.Handler(websocketServer, idGenerator))
}
