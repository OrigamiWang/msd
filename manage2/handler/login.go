package handler

import (
	"github.com/OrigamiWang/msd/manage2/dal"
	"github.com/OrigamiWang/msd/manage2/model/dto"
	"github.com/OrigamiWang/msd/micro/auth/crypto"
	"github.com/OrigamiWang/msd/micro/const/errcode"
	"github.com/OrigamiWang/msd/micro/model/errx"
	"github.com/gin-gonic/gin"
)

func LoginBinder() interface{} {
	return &dto.LoginReq{}
}

func LoginHandler(c *gin.Context, req interface{}) (resp interface{}, err errx.ErrX) {
	loginReq := req.(*dto.LoginReq)
	name := loginReq.Name
	pswd := loginReq.Pswd
	encodedPswd := crypto.Md5Encode(pswd)
	resp, e := dal.Login(name, encodedPswd)
	if e != nil {
		return nil, errx.New(errcode.Success, "name or password incorrect")
	}
	return resp, nil
}
