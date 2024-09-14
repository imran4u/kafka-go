package main

import (
	"fmt"

	"github.com/IBM/sarama"
)

func main() {
	producer, _ := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	defer producer.Close()
	// for i := 0; i < 10; i++ {
	msg := &sarama.ProducerMessage{Topic: "test", Value: sarama.StringEncoder("Hello, Kafka! ${i}")}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Failed to send message:", err)
	}

	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
	// }
}
