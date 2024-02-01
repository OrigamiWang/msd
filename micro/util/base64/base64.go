package base64

import "encoding/base64"

func EncodeBase64(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func DecodeBase64(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}
