package biz

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"github.com/OrigamiWang/msd/micro/const/db"
	"github.com/OrigamiWang/msd/micro/const/mq"
	"github.com/OrigamiWang/msd/micro/mq/kafka"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/OrigamiWang/msd/register-center/model/dao"
	"github.com/OrigamiWang/msd/register-center/model/dto"
)

// heartbeat listener
func ListenHeartBeat() {
	go kafka.ConsumeMsg(mq.KAFKA_HEART_BEAT, heartBeatHandler)
}

func heartBeatHandler(strArr ...string) {
	kafkaKey := strArr[0]
	kafkaVal := strArr[1]
	instConf := dto.InstanceConf{}
	json.Unmarshal([]byte(kafkaVal), &instConf)
	key := fmt.Sprintf("%s_%s_%d", db.HEART_BEAT_REDIS_PREFIX, kafkaKey, instConf.InstanceId)
	dao.RC.Set(key, kafkaVal, time.Second*45)
	// dao.RC.Set(key, kafkaVal, time.Minute*1)
}

// heartbeat beater
func BeatHeartBeat(svcName, val string) {
	go func() {
		for {
			msg := &sarama.ProducerMessage{
				Topic: mq.KAFKA_HEART_BEAT,
				Key:   sarama.StringEncoder(svcName),
				Value: sarama.StringEncoder(val),
			}
			e := kafka.ProduceMsg(msg)
			if e != nil {
				logutil.Error("kafka produce msg failed, err: %v", e)
			}
			time.Sleep(time.Second * 30)
		}
	}()
}
