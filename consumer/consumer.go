package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

func connectConsumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

func main() {
	// Define the topic you want to consume messages from
	topic := "comments"

	// Connect to Kafka and get the consumer (make sure the function name is connectConsumer, not connectCosumer)
	consumer, err := connectConsumer([]string{"localhost:29092"})
	if err != nil {
		// Log and exit if connection fails
		log.Fatalf("Failed to connect to kafka: %v", err)
	}

	// Start consuming from the specified topic, partition 0, from the oldest available offset
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		// Log and exit if partition consumption fails
		log.Fatalf("Failed to start partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	// Indicate that the consumer has started successfully
	fmt.Println("Consumer started")

	// Create a channel to listen for operating system signals (like Ctrl+C)
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Create a channel for signaling when processing is done
	doneCh := make(chan struct{})

	// Start a goroutine to listen for new messages from Kafka
	go func() {
		for msg := range partitionConsumer.Messages() {
			// Print each message received on the topic
			fmt.Printf("Message on %s: %s\n", msg.Topic, string(msg.Value))
		}
		// Signal that all messages have been consumed
		doneCh <- struct{}{}
	}()

	// Wait for an interrupt (Ctrl+C) or termination signal
	<-sigchan

	// Close the consumer to clean up resources
	consumer.Close()

	// Wait for the goroutine to finish processing any last messages
	<-doneCh
}
