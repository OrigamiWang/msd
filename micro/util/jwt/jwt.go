package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/OrigamiWang/msd/micro/util/base64"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
)

const SECRET = "ecef6883c93214e253352b8ac36ea93cf5ac4c34f8bb1f3217bc8376145661fb"

type JwtPayload struct {
	Uid   int       `json:"uid"`   // id
	Uname string    `json:"uname"` // username
	Exp   time.Time `json:"exp"`   // expire time
}

func encodeHeader() string {
	header := map[string]string{}
	header["alg"] = "HS256"
	header["typ"] = "JWT"
	headByte, err := json.Marshal(header)
	if err != nil {
		logutil.Error("jwt header marshal failed, err: %v", err)
		return ""
	}
	headerBase64 := base64.EncodeBase64(headByte)
	return headerBase64
}

func encodePayload(payload *JwtPayload) string {
	payloadByte, err := json.Marshal(payload)
	if err != nil {
		logutil.Error("jwt payload marshal failed, err: %v", err)
		return ""
	}
	payloadBase64 := base64.EncodeBase64(payloadByte)
	return strings.TrimRight(payloadBase64, "=")
}

func encodeSignature(data string) string {
	hmacHasher := hmac.New(sha256.New, []byte(SECRET))
	hmacHasher.Write([]byte(data))
	hmacHashed := hmacHasher.Sum(nil)
	signature := hex.EncodeToString(hmacHashed)
	return signature
}

func EncodeJwt(jwtPayload *JwtPayload) string {
	header := encodeHeader()
	payload := encodePayload(jwtPayload)
	if header == "" || payload == "" {
		logutil.Error("jwt header or payload is empty")
		return ""
	}
	data := header + "." + payload
	signature := encodeSignature(data)
	jwt := data + "." + signature
	return jwt
}
func DecodeJwt(jwt string) (*JwtPayload, error) {
	if jwt == "" {
		logutil.Error("jwt is empty")
		return nil, fmt.Errorf("jwt is empty")
	}
	jwt = strings.TrimSpace(jwt)
	arr := strings.Split(jwt, ".")
	if len(arr) != 3 {
		logutil.Error("jwt is not valid")
		return nil, fmt.Errorf("jwt is not valid")
	}
	// header
	headerBase64 := arr[0]
	err := decodeHeader(headerBase64)
	if err != nil {
		logutil.Error("decode jwt header base64 failed, err: %v", err)
		return nil, nil
	}
	// payload
	payloadBase64 := arr[1]
	jwtPayload := &JwtPayload{}
	err = decodePayload(payloadBase64, jwtPayload)
	if err != nil {
		logutil.Error("decode jwt payload failed, err: %v", err)
		return nil, nil
	}
	// signature
	signature := arr[2]
	data := headerBase64 + "." + payloadBase64
	if checkSignature(data, signature) {
		fmt.Println("signature is valid")
	} else {
		fmt.Println("signature is not valid")
		return nil, nil
	}
	return jwtPayload, nil
}

func decodeHeader(headerBase64 string) error {
	headerBase, error := base64.DecodeBase64(headerBase64)
	if error != nil {
		logutil.Error("decode jwt header base64 failed, err: %v", error)
		return error
	}
	header := map[string]string{}
	error = json.Unmarshal(headerBase, &header)
	if error != nil {
		logutil.Error("jwt header json unmarshal failed, err: %v", error)
		return error
	}
	return nil
}
func decodePayload(payloadBase64 string, jwtPayload *JwtPayload) error {
	payloadBase, error := base64.DecodeBase64(payloadBase64)
	if error != nil {
		logutil.Error("decode jwt payload base64 failed, err: %v", error)
		return error
	}
	error = json.Unmarshal(payloadBase, jwtPayload)
	if error != nil {
		logutil.Error("jwt payload json unmarshal failed, err: %v", error)
		return error
	}
	return nil
}
func checkSignature(data, rawSignature string) bool {
	return rawSignature == encodeSignature(data)
}
