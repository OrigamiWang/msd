package dal

import (
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	var err error
	dsn := "root:abc123@tcp(127.0.0.1:3306)/msd?charset=utf8mb4&parseTime=True&loc=Local"
	// 关闭默认事务
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		logutil.Error("mysql init error")
	}
}
