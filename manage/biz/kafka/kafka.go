package kafka

import (
	"github.com/OrigamiWang/msd/manage/dal"
	"github.com/OrigamiWang/msd/micro/const/mq"
	"github.com/OrigamiWang/msd/micro/mq/kafka"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
)

func updateConf(str ...string) {
	logutil.Info("process msg, conf-center update")
	dal.DelConf()
	dal.InitConf(str[0])
}
func InitConfCenterConsumer() {
	go kafka.ConsumeMsg(mq.KAFKA_CONF_CENTER, updateConf)
}
