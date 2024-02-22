package handler

import (
	"github.com/OrigamiWang/msd/micro/model/errx"
	"github.com/OrigamiWang/msd/register-center/biz"
	"github.com/gin-gonic/gin"
)

func GetSvcRegisterListHandler(c *gin.Context, req interface{}) (resp interface{}, err errx.ErrX) {
	return biz.GetSvcList()
}
func GetSvcRegisterHandler(c *gin.Context, req interface{}) (resp interface{}, err errx.ErrX) {
	name := c.Param("name")
	return biz.GetSvcByName(name)
}
