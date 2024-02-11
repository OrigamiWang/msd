package handler

import (
	"github.com/OrigamiWang/msd/auth/biz"
	"github.com/OrigamiWang/msd/auth/model/dto"
	"github.com/OrigamiWang/msd/micro/const/errcode"
	"github.com/OrigamiWang/msd/micro/model/errx"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/gin-gonic/gin"
)

func AuthenticateBinder() interface{} {
	return &dto.AuthenticateReq{}
}

func AuthenticateHandler(c *gin.Context, req interface{}) (resp interface{}, err errx.ErrX) {
	logutil.Info("authenticate...")
	authenticateReq := req.(*dto.AuthenticateReq)
	logutil.Info("authenticateReq: %v", authenticateReq.JwtToken)
	uid, uname, e := biz.Authenticate(authenticateReq.JwtToken)
	if e != nil {
		return nil, errx.New(errcode.ServerError, e.Error())
	}
	authenticateResp := dto.AuthenticateResp{
		Uid:   uid,
		Uname: uname,
	}
	return authenticateResp, nil
}
