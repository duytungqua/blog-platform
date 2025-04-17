package kafka

func StartConsumer(broker, topic, groupID string) error{
	// Create a new Sarama configuration
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create a new consumer group
	consumerGroup, err := sarama.NewConsumerGroup([]string{broker}, groupID, config)
	if err != nil {
		return fmt.Errorf("failed to create consumer group: %w", err)
	}
	defer consumerGroup.Close()

	// Start consuming messages
	for {
		err := consumerGroup.Consume(context.Background(), []string{topic}, &Consumer{})
		if err != nil {
			return fmt.Errorf("failed to consume messages: %w", err)
		}
	}

	return nil
}