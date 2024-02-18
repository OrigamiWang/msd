package crypto

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func Md5Encode(rawStr string) string {
	hash := md5.Sum([]byte(rawStr))
	md5str := fmt.Sprintf("%x", hash) //将[]byte转成16进制
	// 32 bit result
	return strings.ToUpper(md5str)
}
