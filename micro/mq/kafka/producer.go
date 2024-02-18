package kafka

import (
	"context"
	"time"

	"github.com/OrigamiWang/msd/micro/const/mq"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	Kafka "github.com/segmentio/kafka-go"
)

// ProduceMsg produce message
// id: 分区键, msg: 消息
var (
	KafkaConn *Kafka.Conn
)

func init() {
	var err error
	KafkaConn, err = Kafka.DialLeader(context.Background(), "tcp", "localhost:9092", mq.KAFKA_CONF_CENTER, mq.PARTITION_NUM)
	if err != nil {
		logutil.Error("kafka connect failed, err: %v", err)
		panic(err.Error())
	}
}
func ProduceMsg(key string, msg string) error {
	kafkaMsg := Kafka.Message{
		Key:   []byte(key),
		Value: []byte(msg),
	}
	err := KafkaConn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	if err != nil {
		logutil.Error("kafka set write deadline failed, err: %v", err)
		return err
	}
	_, err = KafkaConn.WriteMessages(kafkaMsg)
	if err != nil {
		logutil.Error("kafka produce msg failed, err: %v", err)
		return err
	}
	return nil
}
