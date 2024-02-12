package kafka

import (
	"github.com/OrigamiWang/msd/micro/const/mq"
	Kafka "github.com/segmentio/kafka-go"
)

// consume message
var (
	KafkaConsumer *Kafka.Reader
)

func init() {
	KafkaConsumer = Kafka.NewReader(Kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   mq.KAFKA_CONF_CENTER,
	})
}
func ConsumeMsg(id int) error {
	return nil
}
