package dal

import (
	"fmt"

	"github.com/OrigamiWang/msd/manage/cli"
	"github.com/OrigamiWang/msd/micro/confparser"
	"github.com/OrigamiWang/msd/micro/const/svc"
	"github.com/OrigamiWang/msd/micro/model/dao"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

var (
	DATABSE_KEY = "sample_mysql1"
	conn        *gorm.DB
)

func init() {
	InitConf(svc.MANAGE)
	dao.InitDb()
	InitMysqlConn()
}
func InitConf(svcName string) {
	fmt.Println("init conf...")
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
}

func DelConf() {
	confparser.Conf = nil
	dao.DelDB()
	DelConn()
}

func InitMysqlConn() {
	var err error
	conn, err = dao.MySQL(DATABSE_KEY)
	if conn == nil || err != nil {
		logutil.Error("can not connect mysql, database_key: %v, err: %v", DATABSE_KEY, err)
		panic("can not connect mysql")
	}
}

func DelConn() {
	conn = nil
}
