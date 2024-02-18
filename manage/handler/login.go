package handler

import (
	"github.com/OrigamiWang/msd/manage/biz"
	"github.com/OrigamiWang/msd/manage/model/dto"
	"github.com/OrigamiWang/msd/micro/model/errx"
	"github.com/gin-gonic/gin"
)

func LoginBinder() interface{} {
	return &dto.LoginReq{}
}

func LoginHandler(c *gin.Context, req interface{}) (resp interface{}, err errx.ErrX) {
	loginReq := req.(*dto.LoginReq)
	return biz.DoLogin(loginReq)
}
