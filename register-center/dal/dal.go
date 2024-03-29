package dal

import (
	dao2 "github.com/OrigamiWang/msd/micro/model/dao"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"gorm.io/gorm"
)

var (
	DATABSE_KEY = "register-center-mysql"
	conn        *gorm.DB
)

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
