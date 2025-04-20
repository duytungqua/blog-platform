package kafka

import (
	"event-worker/internal/kafka"
	"fmt"
	"log"
)

func StartConsumer(broker, topic, groupID string) error{
	// Create a new Sarama configuration
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        "group.id":          "go-consumer-group",
        "auto.offset.reset": "earliest",
    })
	if err != nil {
        log.Fatalf("Failed to create consumer: %s", err)
    }

    defer consumer.Close()

    err = consumer.Subscribe("my-topic", nil)
    if err != nil {
        log.Fatalf("Failed to subscribe: %s", err)
    }

    fmt.Println("Consumer is up and listening...")

    for {
        msg, err := consumer.ReadMessage(-1)
        if err == nil {
            fmt.Printf("Received message: %s from topic %s\n", string(msg.Value), *msg.TopicPartition.Topic)
        } else {
            fmt.Printf("Consumer error: %v (%v)\n", err, msg)
        }
    }
	return nil
}
