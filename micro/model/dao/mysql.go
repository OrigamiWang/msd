package dao

import (
	"database/sql"
	"fmt"
	"github.com/OrigamiWang/msd/micro/confparser"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(db *confparser.Database) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/msd?charset=utf8mb4&parseTime=True&loc=Local", db.User, db.Password, db.Host, db.Port)
	gormDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		logutil.Error("mysql init failed, err: %v", err)
		return nil
	}
	sqlDb, err := gormDb.DB()
	if err != nil {
		logutil.Error("transfer mysql failed, err: %v", err)
		return nil
	}
	err = sqlDb.Ping()
	if err != nil {
		logutil.Error("ping mysql failed, err: %v", err)
		return nil
	}
	return sqlDb
}
