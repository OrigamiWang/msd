package biz

import (
	"github.com/OrigamiWang/msd/manage/cli"
	"github.com/OrigamiWang/msd/manage/dal"
	"github.com/OrigamiWang/msd/micro/confparser"
	"github.com/OrigamiWang/msd/micro/model/dao"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/mitchellh/mapstructure"
)

func Init(svcName string) {
	InitConf(svcName)
}
func InitConf(svcName string) {
	resp, err := cli.Conf.GetConf(svcName)
	if err != nil {
		logutil.Error("get conf failed, err: %v", err)
	}
	m := resp.(map[string]interface{})
	logutil.Info(m)
	var conf *confparser.Config
	err = mapstructure.Decode(m, &conf)
	if err != nil {
		logutil.Error("marshal json failed, err: %v", err)
		panic("marshal json failed")
	}
	confparser.Conf = conf
	dao.InitDb()
	dal.InitConn()
}

func DelConf() {
	confparser.Conf = nil
	dao.DelDB()
	dal.DelConn()
}
