package base

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

// 封装的输出类，返回指定的格式信息
const (
	STATUS_SUCCESS = 1
	STATUS_ERROR   = 0
	STATUS_PARAMS  = -1
)

// 输出结构体
type BaseVC struct {
	code    int
	message string
	data    interface{}
	Ctx     iris.Context
}

func (vc *BaseVC) HandleError(ctx iris.Context, err error) {
	if iris.IsErrPath(err) {
		// to ignore any "schema: invalid path" you can check the error type
		return // continue.
	}
}

// 获取请求参数并校验；struct 使用form `UserName string `form:"name"``映射字段
func (vc *BaseVC) ReadBody(ptr interface{}) error {
	if err := vc.Ctx.ReadBody(ptr); err != nil {
		// 参数不匹配错误不算
		if !iris.IsErrPath(err) {
			vc.Error(STATUS_PARAMS, err.Error())
			return err
		}
	}

	// [参数校验](https://blog.csdn.net/guyan0319/article/details/105918559/)
	validate := validator.New()
	if err := validate.Struct(ptr); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			vc.Error(STATUS_PARAMS, err.Error())
			return err
		}
	}

	return nil
}

// 输出成功
func (vc *BaseVC) Success(data interface{}) {
	vc.code = STATUS_SUCCESS
	vc.message = "ok"
	vc.data = data

	_, _ = vc.Ctx.JSON(vc.getData())
}

// 输出标准错误
func (vc *BaseVC) Failed(message string) {
	vc.code = STATUS_ERROR
	if message != "" {
		vc.message = message
	} else {
		vc.message = "error"
	}

	_, _ = vc.Ctx.JSON(vc.getData())
}

// 输出指定状态错误
func (vc *BaseVC) Error(code int, message string) {
	vc.code = code
	if message != "" {
		vc.message = message
	} else {
		vc.message = "error"
	}

	// vc.Ctx.StatusCode(code)
	_, _ = vc.Ctx.JSON(vc.getData())
}

// 组装为指定格式的基础数据
func (vc *BaseVC) getData() map[string]interface{} {
	data := make(map[string]interface{})
	data["message"] = vc.message
	data["code"] = vc.code
	data["data"] = vc.data

	return data
}
