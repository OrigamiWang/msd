package dal

import (
	dao2 "github.com/OrigamiWang/msd/micro/model/dao"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/OrigamiWang/msd/register-center/model/dao"
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

func GetSvcByName(name string) (*dao.SvcRegister, error) {
	s := &dao.SvcRegister{}
	result := conn.Where(&dao.SvcRegister{Name: name}).First(s)
	return s, result.Error
}

func GetAllSvc() (*[]dao.SvcRegister, error) {
	s := &[]dao.SvcRegister{}
	result := conn.Find(&s)
	return s, result.Error
}
