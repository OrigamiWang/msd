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
	// var err error

	// kafkaConsumer, err = sarama.NewConsumer([]string{"localhost:9092"}, nil)
	// if err != nil {
	// 	logutil.Error("Failed to start consumer, err: %v", err)
	// }
	// martitionConsumerMap = make(map[string]sarama.PartitionConsumer)
}

func ConsumeMsg(topic string, handler func(string)) error {
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
		logutil.Info("Received message: %s", string(message.Value))
		handler(string(message.Value))
	}
	return nil
}
func ConsumeMsgTest(topic string, handler func(string)) error {
	consumer, _ := sarama.NewConsumer([]string{"localhost:9092"}, nil)

	partitionConsumer, _ := consumer.ConsumePartition(mq.KAFKA_CONF_CENTER, 0, sarama.OffsetNewest)
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			logutil.Error("Failed to close partition consumer, err: %v", err)
		}
	}()

	for message := range partitionConsumer.Messages() {
		logutil.Info("Received message: %s", string(message.Value))
		handler(string(message.Value))
	}
	return nil
}
