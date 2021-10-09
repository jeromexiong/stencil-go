package base

import (
	"github.com/kataras/iris/v12"
)

// 通用错误处理
func ErrorHandle(app *iris.Application) {
	httpError := HttpError{}

	app.OnErrorCode(iris.StatusNotFound, httpError.notFound)
	app.OnErrorCode(iris.StatusInternalServerError, httpError.internalServerError)
	app.OnErrorCode(401, httpError.authError)
}

type HttpError struct{ BaseVC }

func (h *HttpError) notFound(ctx iris.Context) {
	h.Ctx = ctx
	h.Error(404, "404 not found")
}

func (h *HttpError) internalServerError(ctx iris.Context) {
	h.Ctx = ctx
	h.Error(500, "Oops something went wrong, try again")
}

func (h *HttpError) authError(ctx iris.Context) {
	h.Ctx = ctx
	h.Error(401, "Oops auth fail")
}
