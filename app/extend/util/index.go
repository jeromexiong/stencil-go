package util

import (
	"stencil-go/app/extend/util/hashUtil"
)

/** 校验密钥
- pwd 密码
- salt 盐
- encryptPwd 加盐加密后的密码
*/
func CheckPwd(pwd string, salt string, encryptPwd string) bool {
	_pwd := hashUtil.Md5String(hashUtil.Md5String(pwd) + salt)
	return _pwd == encryptPwd
}
