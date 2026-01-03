package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
)

type Comment struct {
	Text string `json:"text"`
}

func main() {
	// create new fiber app
	app := fiber.New()
	// group API endpoints under /api/v1
	api := app.Group("/api/v1")
	// route POST /comments to createComment handler
	api.Post("/comments", createComment)

	// start server on port 3000
	app.Listen(":3000")
}

func ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {
	// create new sarama config for producer
	config := sarama.NewConfig()
	// ensure the producer returns on success
	config.Producer.Return.Successes = true
	// set acks to WaitForAll for durability
	config.Producer.RequiredAcks = sarama.WaitForAll
	// set max retry for producer
	config.Producer.Retry.Max = 5

	// create a new SyncProducer with the given brokers and config
	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func PushCommentToQueue(topic string, message []byte) error {
	// set kafka broker(s) url
	brokersUrl := []string{"localhost:29092"}

	// connect to kafka producer
	producer, err := ConnectProducer(brokersUrl)
	if err != nil {
		// return on connection error
		log.Printf("Failed to connect to kafka: %v", err)
		return err
	}
	defer producer.Close()

	// create a ProducerMessage for kafka
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	// send the message to kafka
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		// return on publishing error
		log.Printf("Failed to send message to kafka: %v", err)
		return err
	}

	// print successfully sent message details
	fmt.Printf("Message sent to topic %s, partition %d, offset %d\n", topic, partition, offset)
	return nil
}

func createComment(c *fiber.Ctx) error {

	// Pointer of comment struct
	cmt := new(Comment)

	// read the body parse(json) it and store it in cmt variable(pointer)
	if err := c.BodyParser(cmt); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// convert struct in bytes to send in kafka
	msgBytes, err := json.Marshal(cmt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to marshal comment",
		})
	}

	// push the message to kafka topic
	if err := PushCommentToQueue("comments", msgBytes); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to send message to Kafka",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Comment created successfully",
		"comment": cmt,
	})

}
