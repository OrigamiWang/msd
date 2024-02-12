package dao

import (
	"fmt"
	"github.com/OrigamiWang/msd/micro/confparser"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(db *confparser.Database) *gorm.DB {
	fmt.Println("init mysql...")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", db.User, db.Password, db.Host, db.Port, db.Name)
	fmt.Println("dsn: " + dsn)
	gormDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		logutil.Error("Open MySQL connection failed. key: %v", db.Key)
		return nil
	}

	sqlDb, _ := gormDb.DB()

	// 通过底层的sqlDb去设置更细粒度的配置
	maxIdle := db.ExtInt("maxIdle", 0)
	sqlDb.SetMaxOpenConns(maxIdle)

	maxOpen := db.ExtInt("maxOpen", 8)
	sqlDb.SetMaxOpenConns(maxOpen)

	lifetime := db.ExtDuration("maxConnLifetime", "0s")
	sqlDb.SetConnMaxLifetime(lifetime)

	return gormDb
}
