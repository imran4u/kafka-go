package main

import (
	"fmt"

	"github.com/IBM/sarama"
)

func main() {
	consumer, _ := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	defer consumer.Close()

	partitionConsumer, _ := consumer.ConsumePartition("test", 0, sarama.OffsetOldest)
	defer partitionConsumer.Close()

	for message := range partitionConsumer.Messages() {
		fmt.Printf("Consumed message: %s\n", string(message.Value))
	}
}
