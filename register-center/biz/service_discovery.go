package biz

import (
	"time"

	"github.com/OrigamiWang/msd/micro/const/db"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/OrigamiWang/msd/register-center/model/dao"
)

func ListenServiceDiscovery() {
	go func() {
		for {
			logutil.Info("listen service discovery...")
			var svc_cnt int64 = 10
			if keys, _, err := dao.RC.Scan(db.HEART_BEAT_REDIS_PREFIX, svc_cnt); err == nil {
				for _, key := range keys {
					if val, err := dao.RC.Get(key); err == nil {
						logutil.Info("key: %v, val: %v", key, val)
					}
				}
			}
			time.Sleep(time.Second * 10)
		}
	}()
}
