package util

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
	"stencil-go/app/extend/util/hashUtil"
)

// 返回指定位数随机数字字符串
func RandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

// 返回指定位数随机数字
func RandomNumber(len int) string {
	var numbers = []byte{1, 2, 3, 4, 5, 7, 8, 9}
	var container string
	length := bytes.NewReader(numbers).Len()

	for i := 1; i <= len; i++ {
		random, _ := rand.Int(rand.Reader, big.NewInt(int64(length)))
		container += fmt.Sprintf("%d", numbers[random.Int64()])
	}
	return container
}

// 二分幂法 求x^n
func Powerf(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else {
		return x * Powerf(x, n-1)
	}
}

// 加盐生成密钥
func EncryptPwd(pwd string) (salt string, encrypt string) {
	salt = RandomString(16)
	encrypt = hashUtil.Md5String(hashUtil.Md5String(pwd) + salt)
	return
}

/** 校验密钥
- pwd 密码
- salt 盐
- encryptPwd 加盐加密后的密码
*/
func CheckPwd(pwd string, salt string, encryptPwd string) bool {
	_pwd := hashUtil.Md5String(hashUtil.Md5String(pwd) + salt)
	return _pwd == encryptPwd
}
