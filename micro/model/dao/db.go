package dao

import (
	"github.com/OrigamiWang/msd/micro/confparser"
	"github.com/OrigamiWang/msd/micro/const/db"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"gorm.io/gorm"
)

// 初始化databases
var dbConns = make(map[string]interface{}, 0)

func init() {
	InitDb()
}
func InitDb() {
	dbs := confparser.Conf.Dbs
	for _, _db := range dbs {
		switch _db.Type {
		case db.MYSQL:
			conn := InitMysql(&_db)
			if conn == nil {
				continue
			}
			dbConns[_db.Key] = conn
		case db.REDIS:
			continue
		case db.MONGODB:
			continue
		default:
			continue
		}
	}
}

// MySQL 返回mysql数据库
func MySQL(key string) (mysql *gorm.DB, err error) {
	if v, ok := dbConns[key]; ok {
		return v.(*gorm.DB), nil
	}
	logutil.Error("connect mysql failed, err: %v", err)
	return nil, err
}
