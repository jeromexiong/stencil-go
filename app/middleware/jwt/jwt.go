package jwt

import (
	"context"
	"fmt"
	"regexp"
	. "stencil-go/app/core"
	. "stencil-go/app/core/jwt"
	"time"

	"github.com/kataras/iris/v12"
)

var _ctx = context.Background()

func New(ctx iris.Context) {
	// 白名单判断
	if func(path string) bool {
		for _, v := range Config.Own.IgnoreURLs {
			reg := regexp.MustCompile(v)
			result := reg.FindAllString(path, -1)
			if result != nil {
				return true
			}
		}
		return false
	}(ctx.Path()) {
		ctx.Next()
		return
	}

	info := Verify(ctx)
	if info == nil {
		ctx.StopWithStatus(401)
		return
	}
	ctx.Values().Set("user", info)
	ctx.Next()
}

// 获取用户信息
func User(ctx iris.Context) *Info {
	if val := ctx.Values().Get("user"); val != nil {
		user := val.(*Info)
		return user
	}
	return nil
}

// jwt 加密
func Sign(UserId int64, Type uint8, Platform Platform) string {
	info := Info{UserId: UserId, Type: Type, Platform: Platform}
	token := GenToken(info)

	if Config.Redis.Required {
		key := fmt.Sprintf("%d_%d_%s", info.UserId, info.Type, info.Platform)
		expireTime := time.Duration(Config.JWTTimeout) * time.Second
		if info.Platform == MOBILE || info.Platform == PAD {
			expireTime = 3600 * 24 * 365 * time.Second
		}
		if err := Redis.Set(_ctx, key, token, expireTime).Err(); err != nil {
			Log.Error(err)
		}
	}

	return token
}

// JWT 解密
func Verify(ctx iris.Context) *Info {
	info, err := ParseToken(ctx)
	if err != nil {
		return nil
	}

	if Config.Redis.Required {
		key := fmt.Sprintf("%d_%d_%s", info.UserId, info.Type, info.Platform)
		expireTime := time.Duration(Config.JWTTimeout) * time.Second
		token, _ := Redis.Get(_ctx, key).Result()
		if token == GetToken(ctx) {
			Redis.Expire(_ctx, key, expireTime)
			return info
		}
		return nil
	}

	return info
}

func Remove(ctx iris.Context) {
	info, err := ParseToken(ctx)
	if err == nil {
		key := fmt.Sprintf("%d_%d_%s", info.UserId, info.Type, info.Platform)
		Redis.Del(_ctx, key)
	}
}
