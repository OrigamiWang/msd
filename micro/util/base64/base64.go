package base64

import (
	"encoding/base64"
	"strings"
)

func EncodeBase64(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

func DecodeBase64(str string) ([]byte, error) {
	// fill '=' until the length of the str is the multiple of 4
	if l := len(str) % 4; l > 0 {
		str += strings.Repeat("=", 4-l)
	}
	decoded, err := base64.URLEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}

	return decoded, nil
}
