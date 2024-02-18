package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/OrigamiWang/msd/conf-center/biz"
	"github.com/OrigamiWang/msd/conf-center/dal"
	"github.com/OrigamiWang/msd/conf-center/model/dto"
	"github.com/OrigamiWang/msd/micro/confparser"
	"github.com/OrigamiWang/msd/micro/model"
	"github.com/OrigamiWang/msd/micro/model/errx"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

func ConfBinder() interface{} {
	return &dto.SvcConfReq{}
}

func GetConfigHandler(c *gin.Context) {
	svcName := c.Param("name")
	if svcName == "" {
		logutil.Error("incorrect param")
		c.JSON(http.StatusBadRequest, model.Response{Ts: fmt.Sprintf("%v", time.Now().Unix()), Msg: "incorrect param"})
		return
	}
	config, err := dal.GetConfigByName(svcName)
	if err != nil {
		logutil.Error("get config by name failed, err: %v", err)
		c.JSON(http.StatusInternalServerError, model.Response{Ts: fmt.Sprintf("%v", time.Now().Unix()), Msg: "mysql error"})
		return
	}
	conf := config.Conf
	res := &confparser.Config{}
	err = yaml.Unmarshal([]byte(conf), res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Ts: fmt.Sprintf("%v", time.Now().Unix()), Msg: "json unmarshal error"})
		return
	}
	c.JSON(http.StatusOK, res)
}

func UpdateConfigHandler(c *gin.Context, req interface{}) (resp interface{}, err errx.ErrX) {
	svcName := c.Param("name")
	if svcName == "" {
		logutil.Error("incorrect param")
		c.JSON(http.StatusBadRequest, model.Response{Ts: fmt.Sprintf("%v", time.Now().Unix()), Msg: "incorrect param"})
		return
	}
	svcConfReq := req.(*dto.SvcConfReq)
	return biz.UpdateConf(svcName, svcConfReq)
}
