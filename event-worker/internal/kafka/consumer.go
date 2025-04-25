package kafka

import (
	"event-worker/internal/kafka"
	"fmt"
	"log"
)

func StartConsumer(broker, topic, groupID string) error{
	kafkaConsumer, err := kafka.NewConsumer(broker, topic, groupID)
	if err != nil {
		return fmt.Errorf("failed to create consumer: %w", err)
	}
	defer kafkaConsumer.Close()
	for {		
		message, err := kafkaConsumer.Consume()
		if err != nil {
			log.Printf("Error consuming message: %v", err)
			continue
		}
		log.Printf("Received message: %s", message)
	}
	return nil
}
