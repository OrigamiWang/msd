package biz

import (
	"time"

	"github.com/IBM/sarama"
	"github.com/OrigamiWang/msd/micro/const/db"
	"github.com/OrigamiWang/msd/micro/const/mq"
	"github.com/OrigamiWang/msd/micro/mq/kafka"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/OrigamiWang/msd/register-center/model/dao"
)

// heartbeat listener
func ListenHeartBeat() {
	go kafka.ConsumeMsg(mq.KAFKA_HEART_BEAT, heartBeatHandler)
}

func heartBeatHandler(svcName string) {
	key := db.HEART_BEAT_REDIS_PREFIX + svcName
	dao.RC.Set(key, "", time.Second*5)
}

// heartbeat beater
func BeatHeartBeat(val string) {
	go func() {
		for {
			msg := &sarama.ProducerMessage{
				Topic: mq.KAFKA_CONF_CENTER,
				Value: sarama.StringEncoder(val),
			}
			e := kafka.ProduceMsg(msg)
			if e != nil {
				logutil.Error("kafka produce msg failed, err: %v", e)
			}
			time.Sleep(time.Second * 2)
		}
	}()
}
