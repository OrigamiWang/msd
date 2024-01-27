package handler

import (
	"github.com/OrigamiWang/msd/gorm-demo/dal"
	"github.com/OrigamiWang/msd/gorm-demo/model/dto"
	"github.com/OrigamiWang/msd/micro/model/errx"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/gin-gonic/gin"
)

func UserBinder() interface{} {
	return &dto.UserReq{}
}

func GetFirstUser(c *gin.Context, req interface{}) (resp interface{}, err errx.ErrX) {
	r := req.(*dto.UserReq)
	logutil.Info(r)
	return dal.GetFirstUser(), nil
}
