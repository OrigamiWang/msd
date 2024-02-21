package kafka

import (
	"github.com/OrigamiWang/msd/manage/dal"
	"github.com/OrigamiWang/msd/micro/mq/kafka"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
)

func processResult(svcName string, resultChan <-chan bool) {
	for res := range resultChan {
		if res {
			logutil.Info("process msg, conf-center update")
			// get conf
			dal.DelConf()
			dal.InitConf(svcName)
		}
	}
}
func InitKafkaConsumer(svcName string) {
	// start kafka consumer
	stopChan := make(chan struct{})
	resultChan := make(chan bool)
	go kafka.PollConsume(svcName, stopChan, resultChan)
	go processResult(svcName, resultChan)
}
