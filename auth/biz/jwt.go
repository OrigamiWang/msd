package biz

import (
	"github.com/OrigamiWang/msd/micro/util/jwt"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"time"
)

// 授权
func Authorize(uid int, uname string) string {
	jwtPayload := &jwt.JwtPayload{
		Uid:   uid,
		Uname: uname,
		Exp:   time.Now().Add(time.Hour * 24 * 365), // one year to expire
	}
	j := jwt.EncodeJwt(jwtPayload)
	return j
}

// 鉴权
func Authenticate(j string) (uid int, uname string) {
	jwtPayload, err := jwt.DecodeJwt(j)
	if err != nil {
		logutil.Error("decode jwt failed, error: %v", err)
		return -1, ""
	}
	// time has expired
	if jwtPayload.Exp.Before(time.Now()) {
		logutil.Warn("jwt token has expired")
		return -1, ""
	}
	return jwtPayload.Uid, jwtPayload.Uname
}
