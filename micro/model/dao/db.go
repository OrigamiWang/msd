package dao

import (
	"fmt"

	"github.com/OrigamiWang/msd/micro/confparser"
	"github.com/OrigamiWang/msd/micro/const/db"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// 初始化databases
var dbConns = make(map[string]interface{}, 0)

func init() {
	InitDb()
}
func InitDb() {
	fmt.Println("init db...")
	if confparser.Conf == nil {
		return
	}
	dbs := confparser.Conf.Dbs
	if dbs == nil {
		logutil.Warn("dbs is nil")
		return
	}
	for _, _db := range dbs {
		switch _db.Type {
		case db.MYSQL:
			conn := InitMysql(&_db)
			if conn == nil {
				continue
			}
			dbConns[_db.Key] = conn
		case db.REDIS:
			conn := InitRedis(&_db)
			if conn == nil {
				continue
			}
			dbConns[_db.Key] = conn
		case db.MONGODB:
			continue
		default:
			continue
		}
	}
}

func DelDB(key string) {
	dbConns = make(map[string]interface{}, 0)
}

// MySQL 返回mysql数据库
func MySQL(key string) (*gorm.DB, error) {
	fmt.Println("mysql...")
	if v, ok := dbConns[key]; ok {
		return v.(*gorm.DB), nil
	}
	logutil.Error("get mysql failed, err")
	return nil, fmt.Errorf("get mysql failed, err")
}

func Redis(key string) (*redis.Client, error) {
	fmt.Println("redis...")
	if v, ok := dbConns[key]; ok {
		return v.(*redis.Client), nil
	}
	logutil.Error("get redis failed")
	return nil, fmt.Errorf("get mysql redis")
}
