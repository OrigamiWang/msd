package dal

import (
	"github.com/OrigamiWang/msd/conf-center/model/dao"
	dao2 "github.com/OrigamiWang/msd/micro/model/dao"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"gorm.io/gorm"
)

var DATABSE_KEY = "sample_mysql1"

var conn *gorm.DB

func init() {
	var err error
	conn, err = dao2.MySQL(DATABSE_KEY)
	if err != nil {
		logutil.Error("can not connect mysql, database_key: %v, err: %v", DATABSE_KEY, err)
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
