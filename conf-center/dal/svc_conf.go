package dal

import (
	"github.com/OrigamiWang/msd/conf-center/model/dao"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"gorm.io/gorm"
)

var DATABSE_KEY = "conf-center-mysql"

var conn *gorm.DB

func GetConfigByName(name string) (*dao.SvcConfig, error) {
	config := &dao.SvcConfig{}
	result := conn.Where("svc_name = ?", name).First(config)
	if result.Error != nil {
		logutil.Error("get config failed, err: &v", result.Error)
		return nil, result.Error
	}
	return config, nil
}
