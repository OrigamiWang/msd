package kafka

import (
	"context"
	"time"

	logutil "github.com/OrigamiWang/msd/micro/util/log"
	Kafka "github.com/segmentio/kafka-go"
)

// consume message
var (
	KafkaConsumer *Kafka.Reader
)

func init() {
	KafkaConsumer = Kafka.NewReader(Kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "test",
		MaxWait:  time.Second * 10,
		MinBytes: 10e3,
		MaxBytes: 1e6,
	})
}
func ConsumeMsg(key string) (bool, error) {
	ctx := context.Background()
	res := false
	msg, err := KafkaConsumer.ReadMessage(ctx)
	if err != nil {
		logutil.Error("Error reading message: %v\n", err)
		return res, err
	}
	logutil.Info("key1: %s, key2: %s", string(msg.Key), key)
	if string(msg.Key) == key {
		res = true
	}
	KafkaConsumer.SetOffset(Kafka.LastOffset)
	return res, nil
}
func PollConsume(key string, stopChan <-chan struct{}, resultChan chan<- bool) {
	for {
		logutil.Info("polling...")
		select {
		case <-stopChan:
			logutil.Info("Received stop signal. Exiting poll function.")
			return
		default:
			// polling
			res, err := ConsumeMsg(key)
			if err != nil {
				logutil.Error("Error consuming message: %v", err)
				break
			}
			logutil.Info("res: %v", res)
			resultChan <- res
			time.Sleep(time.Second * 1)
		}

	}
}
