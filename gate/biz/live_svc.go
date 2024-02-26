package biz

import (
	"github.com/OrigamiWang/msd/micro/const/db"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/OrigamiWang/msd/register-center/model/dao"
)

func GetLiveSvc() (res map[string]interface{}, err error) {
	var svc_cnt int64 = 10
	match := db.HEART_BEAT_REDIS_PREFIX + "*"
	res = make(map[string]interface{})
	if keys, _, err := dao.RC.Scan(match, svc_cnt); err == nil {
		for _, key := range keys {
			if val, err := dao.RC.Get(key); err == nil {
				logutil.Info("key: %v, val: %v", key, val)
				res[key[len(db.HEART_BEAT_REDIS_PREFIX)+1:]] = val
			} else {
				logutil.Error("get redis error: %v, err")
			}
		}
	} else {
		logutil.Error("scan redis error: %v, err")
	}
	return
}
