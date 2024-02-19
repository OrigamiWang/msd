package kafka

import (
	"github.com/OrigamiWang/msd/manage/biz"
	"github.com/OrigamiWang/msd/micro/const/svc"
	"github.com/OrigamiWang/msd/micro/mq/kafka"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
)

func processResult(resultChan <-chan bool) {
	for res := range resultChan {
		if res {
			logutil.Info("process msg, conf-center update")
			// get conf
			biz.DelConf()
			biz.InitConf(svc.MANAGE)
		}
	}
}
func InitKafkaConsumer() {
	// start kafka consumer
	stopChan := make(chan struct{})
	resultChan := make(chan bool)
	go kafka.PollConsume(svc.MANAGE, stopChan, resultChan)
	go processResult(resultChan)
}
