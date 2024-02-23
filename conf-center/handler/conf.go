package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/IBM/sarama"
	"github.com/OrigamiWang/msd/conf-center/biz"
	"github.com/OrigamiWang/msd/conf-center/dal"
	"github.com/OrigamiWang/msd/conf-center/model/dto"
	"github.com/OrigamiWang/msd/micro/confparser"
	"github.com/OrigamiWang/msd/micro/const/errcode"
	"github.com/OrigamiWang/msd/micro/const/mq"
	"github.com/OrigamiWang/msd/micro/model"
	"github.com/OrigamiWang/msd/micro/model/errx"
	"github.com/OrigamiWang/msd/micro/mq/kafka"
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

	// produce msg to kafka
	msg := &sarama.ProducerMessage{
		Topic: mq.KAFKA_CONF_CENTER,
		Value: sarama.StringEncoder(svcName),
		Key:   sarama.StringEncoder(svcName),
	}

	resp, err = biz.UpdateConf(svcName, svcConfReq)
	e := kafka.ProduceMsg(msg)
	if e != nil {
		logutil.Error("kafka produce msg failed, err: %v", e)
		return nil, errx.New(errcode.ServerError, "kafka produce msg failed")
	}
	return
}
