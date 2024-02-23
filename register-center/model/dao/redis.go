package dao

import (
	"github.com/OrigamiWang/msd/micro/model/dao"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
)

var (
	RC          *dao.RedisClient
	DATABSE_KEY = "heart_beat_redis1"
)

func init() {
	client, err := dao.Redis(DATABSE_KEY)
	if err != nil {
		logutil.Error("fail to get redis connection, err: %v", err)
		panic(err.Error())
	}
	RC = &dao.RedisClient{
		Client: client,
	}
}
