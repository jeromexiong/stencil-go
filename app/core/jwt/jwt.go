package jwt

import (
	. "stencil-go/app/core/config"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

type Platform string

const (
	NONE   Platform = "none"
	PC     Platform = "pc"
	MOBILE Platform = "mobile"
	PAD    Platform = "pad"
)

type Info struct {
	UserId   int64    `json:"id"`
	Type     uint8    `json:"type"`     // 用户类型(同id不同类型用户)
	Platform Platform `json:"platform"` // 平台唯一(none,pc,mobile,pad
}

var sigKey = []byte(GConfig.Keys)

// 生成token
func GenToken(info Info) string {
	expire := 3600 * 2 * time.Second
	if GConfig.Redis.Required {
		expire = 3600 * 24 * 7 * time.Second
		// expireTime := time.Duration(GConfig.JWTTimeout) * time.Second
		if info.Platform == MOBILE || info.Platform == PAD {
			expire = 3600 * 24 * 365 * time.Second
		}
	}

	signer := jwt.NewSigner(jwt.HS256, sigKey, expire)
	bytes, _ := signer.Sign(info)
	token := string(bytes)

	return token
}

// 解析token
func ParseToken(ctx iris.Context) (*Info, error) {
	bytes := []byte(GetToken(ctx))
	verifiedToken, err := jwt.Verify(jwt.HS256, sigKey, bytes)
	if err != nil {
		return nil, err
	}
	var info Info
	verifiedToken.Claims(&info)

	return &info, nil
}

// 获取token
func GetToken(ctx iris.Context) string {
	token := ctx.GetHeader("token")
	if token == "" {
		token = ctx.GetHeader("Token")
		if token == "" {
			token = ctx.URLParamTrim("token")
			if token == "" {
				token = ctx.FormValue("token")
			}
		}
	}
	return token
}
