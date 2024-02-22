package biz

import (
	"github.com/OrigamiWang/msd/micro/const/errcode"
	"github.com/OrigamiWang/msd/micro/model/errx"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/OrigamiWang/msd/register-center/dal"
	"github.com/OrigamiWang/msd/register-center/model/dto"
)

func GetSvcList() (resp interface{}, err errx.ErrX) {
	svcDaoArr, e := dal.GetAllSvc()
	if e != nil {
		logutil.Error("mysql get all svc failed, err: %v", e)
		return nil, errx.New(errcode.MysqlErr, e.Error())
	}
	svcRespArr := []dto.SvcRegisterResp{}
	for _, svcDao := range *svcDaoArr {
		svcRespArr = append(svcRespArr, dto.SvcRegisterResp{
			ID:     svcDao.ID,
			Name:   svcDao.Name,
			Config: svcDao.Config,
		})
	}
	resp = &svcRespArr
	return
}

func GetSvcByName(name string) (resp interface{}, err errx.ErrX) {
	svcDao, e := dal.GetSvcByName(name)
	if e != nil {
		logutil.Error("mysql get svc by name failed, err: %v", e)
		return nil, errx.New(errcode.MysqlErr, e.Error())
	}
	resp = &dto.SvcRegisterResp{
		ID:     svcDao.ID,
		Name:   svcDao.Name,
		Config: svcDao.Config,
	}
	return
}
