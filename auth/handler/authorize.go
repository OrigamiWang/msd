package handler

import (
	"fmt"
	"github.com/OrigamiWang/msd/auth/model/dto"
	"github.com/OrigamiWang/msd/micro/model/errx"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/gin-gonic/gin"
)

func AuthorizeBinder() interface{} {
	return &dto.AuthorizeReq{}
}

func AuthorizeHandler(c *gin.Context, req interface{}) (resp interface{}, err errx.ErrX) {
	logutil.Info("authorize...")
	authorizeReq := req.(*dto.AuthorizeReq)
	fmt.Println(authorizeReq)
	return nil, nil
}
