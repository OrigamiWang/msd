package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/OrigamiWang/msd/micro/const/mq"
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
		Topic:    mq.KAFKA_CONF_CENTER,
		MaxWait:  time.Second * 10,
		MinBytes: 10e3,
		MaxBytes: 1e6,
	})
}
func ConsumeMsg(key string) error {
	// 设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 读取一条消息
	msg, err := KafkaConsumer.ReadMessage(ctx)
	if err != nil {
		fmt.Printf("Error reading message: %v\n", err)
		return err
	}
	// 打印消息的key和value
	fmt.Printf("Received message with key: %s, value: %s\n", string(msg.Key), string(msg.Value))
	logutil.Info("Received message with key: %s, value: %s\n", string(msg.Key), string(msg.Value))
	return nil
}
