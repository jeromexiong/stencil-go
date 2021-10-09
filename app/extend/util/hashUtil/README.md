# go 常用hash函数

## 示例

```
package main

import (
	"fmt"

	"github.com/ThreeKing2018/goutil/hash"
)

func main() {
	/* 字符串 */
	// md5
	md5 := hashUtil.Md5String("111111")
	fmt.Println(md5)
	// sha1
	sha1 := hashUtil.Sha1String("111111")
	fmt.Println(sha1)

	/* 字节 */
	// md5
	md5 = hashUtil.Md5Byte([]byte("111111"))
	fmt.Println(md5)
	// sha1
	sha1 = hashUtil.Sha1Byte([]byte("111111"))
	fmt.Println(sha1)

	/* 文件 */
	md5, err := hashUtil.Md5File("./test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(md5)
}

```