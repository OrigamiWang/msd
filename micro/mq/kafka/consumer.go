package kafka

import (
	"github.com/IBM/sarama"
	"github.com/OrigamiWang/msd/micro/const/mq"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
)

// consume message
var (
	kafkaConsumer sarama.Consumer
)

func init() {
	var err error

	kafkaConsumer, err = sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		logutil.Error("Failed to start consumer, err: %v", err)
	}
}

func ConsumeMsg(topic string, handler func(string)) error {
	partitionConsumer, err := kafkaConsumer.ConsumePartition(topic, mq.PARTITION_NUM, sarama.OffsetNewest)
	if err != nil {
		logutil.Error("Failed to start partition consumer, err: %v", err)
		return err
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			logutil.Error("Failed to close partition consumer, err: %v", err)
		}
	}()

	for message := range partitionConsumer.Messages() {
		logutil.Info("Received message: %s, key: %s", string(message.Value), string(message.Key))
		handler(string(message.Key))
	}
	return nil
}

// func ConsumeMsg(key string) (bool, error) {
// 	ctx := context.Background()
// 	res := false
// 	msg, err := KafkaConsumer.ReadMessage(ctx)
// 	if err != nil {
// 		logutil.Error("Error reading message: %v\n", err)
// 		return res, err
// 	}
// 	logutil.Info("consume msg, key: %s, value: %s", string(msg.Key), string(msg.Value))
// 	if string(msg.Key) == key {
// 		res = true
// 	}
// 	KafkaConsumer.SetOffset(Kafka.LastOffset)
// 	return res, nil
// }
// func PollConsume(key string, stopChan <-chan struct{}, resultChan chan<- bool) {
// 	for {
// 		select {
// 		case <-stopChan:
// 			logutil.Info("Received stop signal. Exiting poll function.")
// 			return
// 		default:
// 			// polling
// 			res, err := ConsumeMsg(key)
// 			if err != nil {
// 				logutil.Error("Error consuming message: %v", err)
// 				break
// 			}
// 			resultChan <- res
// 			time.Sleep(time.Second * 1)
// 		}

// 	}
// }
