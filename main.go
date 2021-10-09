package main

import (
	"fmt"

	"github.com/kataras/iris/v12"

	config_data "stencil-go/app/bindata/config"
	"stencil-go/app/core"
	"stencil-go/app/core/config"
	"stencil-go/app/core/db"
	"stencil-go/app/core/logger"
	"stencil-go/app/core/redis"
	"stencil-go/app/middleware"
	"stencil-go/app/router"
)

func main() {
	core.Config = config.New(config_data.Asset, false)
	core.Log = logger.New()
	core.DB = db.New()
	core.Redis = redis.New()

	app := iris.New()
	if !core.Config.Production {
		app.Logger().SetLevel("debug")
	}

	// app.HandleDir("/", static_data.AssetFile()) // AssetFile 未设置成功
	// app.RegisterView(iris.HTML(static_data.AssetFile(), ".html").Reload(true))
	app.HandleDir("/", "./public")
	// app.RegisterView(iris.HTML("./public", ".html").Reload(true))

	middleware.New(app)
	router.New(app)

	if err := app.Run(
		iris.Addr(fmt.Sprintf(":%d", config.GConfig.Own.Port)),
		iris.WithConfiguration(config.GConfig.Configuration)); err != nil {
		logger.Log.Fatalf("Start admin server failed. And err:%s", err.Error())
	}

	// ssl 启动
	// app.Run(iris.TLS(":3000", "local.cn+4.pem", "local.cn+4-key.pem"))
}
