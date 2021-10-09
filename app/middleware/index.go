package middleware

import (
	. "stencil-go/app/core"
	"stencil-go/app/middleware/jwt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
)

// 初始化中间件
func New(app *iris.Application) {
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	app.Use(recover.New())

	app.Use(func(ctx iris.Context) {
		Log.Infof("host=[%s], method=[%s], path=[%s]", ctx.Host(), ctx.Method(), ctx.Path())
		ctx.Next()
	})

	app.Use(Cors)
	app.Use(jwt.New)
}

func Cors(ctx iris.Context) {
	origin := ctx.Request().Header.Get("Origin")
	for _, domain := range Config.Domains {
		if origin == domain {
			ctx.Header("Access-Control-Allow-Origin", origin)
			if ctx.Request().Method == "OPTIONS" {
				ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
				ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
				ctx.StatusCode(204)
				return
			}
		}
	}

	ctx.Next()
}
