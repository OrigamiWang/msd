package kafka

import (
	"github.com/IBM/sarama"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
)

// ProduceMsg produce message
// id: 分区键, msg: 消息
var (
	kafkaProducer sarama.SyncProducer
)

func init() {
	var err error

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	kafkaProducer, err = sarama.NewSyncProducer([]string{"localhost:9092"}, config)

	if err != nil {
		logutil.Error("Failed to start Kafka producer, err: %v", err)
	}
}

func ProduceMsg(msg *sarama.ProducerMessage) error {
	partition, offset, err := kafkaProducer.SendMessage(msg)
	if err != nil {
		logutil.Error("Failed to send message, err: %v", err)
		return err
	}

	logutil.Info("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", msg.Topic, partition, offset)
	return nil
}
