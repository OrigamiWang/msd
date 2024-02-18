package biz

import (
	"github.com/OrigamiWang/msd/conf-center/dal"
	"github.com/OrigamiWang/msd/conf-center/model/dto"
	"github.com/OrigamiWang/msd/micro/const/errcode"
	"github.com/OrigamiWang/msd/micro/model/errx"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
)

func UpdateConf(svcName string, req *dto.SvcConfReq) (resp interface{}, err errx.ErrX) {
	confDao, e := dal.UpdateConfByName(svcName, req)
	if e != nil {
		logutil.Error("update conf failed, err: %v", e)
		return nil, errx.New(errcode.MysqlErr, e.Error())
	}
	return confDao, nil
}
