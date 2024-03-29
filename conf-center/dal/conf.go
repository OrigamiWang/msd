package dal

import (
	"github.com/OrigamiWang/msd/conf-center/model/dao"
	"github.com/OrigamiWang/msd/conf-center/model/dto"
	dao2 "github.com/OrigamiWang/msd/micro/model/dao"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"gorm.io/gorm"
)

var DATABSE_KEY = "conf-center-mysql"

var conn *gorm.DB

func init() {
	initMysqlConn()
}

func initMysqlConn() {
	var err error
	conn, err = dao2.MySQL(DATABSE_KEY)
	if conn == nil || err != nil {
		logutil.Error("can not connect mysql, database_key: %v, err: %v", DATABSE_KEY, err)
		panic("can not connect mysql")
	}
}

func GetConfigByName(name string) (*dao.SvcConfig, error) {
	config := &dao.SvcConfig{}
	result := conn.Where("svc_name = ?", name).First(config)
	if result.Error != nil {
		logutil.Error("get config failed, err: &v", result.Error)
		return nil, result.Error
	}
	return config, nil
}

func UpdateConfByName(name string, confReq *dto.SvcConfReq) (*dao.SvcConfig, error) {
	tx := conn.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	conf := &dao.SvcConfig{
		SvcName: name,
		Desc:    confReq.Desc,
		Env:     confReq.Env,
		Conf:    confReq.Conf,
	}
	err := tx.Model(&dao.SvcConfig{}).Where("svc_name = ?", name).Updates(conf).Error
	if err != nil {
		logutil.Error("gorm: update user failed, err: %v", err)
		tx.Rollback()
		return nil, err
	}
	err = tx.Where("svc_name = ?", name).First(conf).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return conf, nil
}
