package kafka

import (
	"context"
	"fmt"
	"time"

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
func ConsumeMsg(key string) error {
	ctx := context.Background()
	msg, err := KafkaConsumer.ReadMessage(ctx)
	if err != nil {
		fmt.Printf("Error reading message: %v\n", err)
		return err
	}
	if string(msg.Key) == key {
		fmt.Printf("Received message with key: %s, value: %s\n", string(msg.Key), string(msg.Value))
	} else {
		fmt.Printf("key not match\n")
	}
	KafkaConsumer.SetOffset(Kafka.LastOffset)
	return nil
}
func PollConsume(stopChan <-chan struct{}) {
	for {
		select {
		case <-stopChan:
			fmt.Println("Received stop signal. Exiting poll function.")
			return
		default:
			// polling
			fmt.Println("polling...")
			err := ConsumeMsg("consumer1")
			if err != nil {
				fmt.Printf("Error consuming message: %v\n", err)
				break
			}
			time.Sleep(time.Second * 1)
		}

	}
}
