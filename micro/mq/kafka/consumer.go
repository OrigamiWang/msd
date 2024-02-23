package kafka

import (
	"github.com/IBM/sarama"
	"github.com/OrigamiWang/msd/micro/const/mq"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
)

// consume message
var (
	kafkaConsumer        sarama.Consumer
	martitionConsumerMap map[string]sarama.PartitionConsumer
)

func init() {
	var err error

	kafkaConsumer, err = sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		logutil.Error("Failed to start consumer, err: %v", err)
	}
	martitionConsumerMap = make(map[string]sarama.PartitionConsumer)
}

func ConsumeMsg(topic string, handler func(...string)) error {
	if martitionConsumerMap[topic] == nil {
		partitionConsumer, err := kafkaConsumer.ConsumePartition(topic, mq.PARTITION_NUM, sarama.OffsetNewest)
		if err != nil {
			logutil.Error("Failed to start partition consumer, err: %v", err)
			return err
		}
		martitionConsumerMap[topic] = partitionConsumer
	}
	partitionConsumer := martitionConsumerMap[topic]
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			logutil.Error("Failed to close partition consumer, err: %v", err)
		}
	}()

	for message := range partitionConsumer.Messages() {
		logutil.Info("Received message, value: %s, key: %s", string(message.Value), string(message.Key))
		handler(string(message.Key), string(message.Value))
	}
	return nil
}
