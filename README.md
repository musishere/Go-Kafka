# Go Kafka Producer-Consumer

A simple Go application demonstrating Kafka message production and consumption. The project consists of a Fiber web API that produces comment messages to Kafka and a consumer that reads and displays those messages.

## Project Structure

```
.
├── producer/
│   └── producer.go    # Fiber API server that produces messages to Kafka
├── consumer/
│   └── consumer.go    # Kafka consumer that reads messages
├── go.mod             # Go module dependencies
└── README.md          # This file
```

## Features

- **Producer**: REST API built with Fiber that accepts comment creation requests and publishes them to Kafka
- **Consumer**: Kafka consumer that reads messages from the "comments" topic and displays them
- **Kafka Integration**: Uses IBM Sarama library for Kafka client operations

## Prerequisites

- Go 1.25.5 or higher
- Kafka broker running on `localhost:29092` (default port)
- Docker (optional, for running Kafka)

## Installation

1. Clone or navigate to the project directory:

```bash
cd "/home/furqi/Documents/Mustafa/Learn go"
```

2. Install dependencies:

```bash
go mod download
```

## Running Kafka (Docker)

If you don't have Kafka running, you can start it using Docker:

```bash
docker run -d \
  --name kafka \
  -p 29092:29092 \
  -e KAFKA_ZOOKEEPER_CONNECT=localhost:2181 \
  -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:29092 \
  -e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 \
  confluentinc/cp-kafka:latest
```

Or use docker-compose if you have a setup configured.

## Usage

### 1. Start the Producer (API Server)

In one terminal, run the producer:

```bash
go run producer/producer.go
```

The API server will start on `http://localhost:3000`

### 2. Start the Consumer

In another terminal, run the consumer:

```bash
go run consumer/consumer.go
```

The consumer will start listening for messages on the "comments" topic.

### 3. Create a Comment

Send a POST request to create a comment:

```bash
curl -X POST http://localhost:3000/api/v1/comments \
  -H "Content-Type: application/json" \
  -d '{"text": "This is a test comment"}'
```

You should see:

- The producer logs showing the message was sent to Kafka
- The consumer displaying the received message

## API Endpoints

### POST `/api/v1/comments`

Creates a new comment and publishes it to Kafka.

**Request Body:**

```json
{
  "text": "Your comment text here"
}
```

**Success Response (200):**

```json
{
  "status": "success",
  "message": "Comment created successfully",
  "comment": {
    "text": "Your comment text here"
  }
}
```

**Error Responses:**

- `400 Bad Request`: Invalid request body
- `500 Internal Server Error`: Failed to marshal comment or send to Kafka

## Configuration

### Kafka Broker

The default Kafka broker URL is set to `localhost:29092` in both producer and consumer. To change it, modify:

- **Producer**: Line 49 in `producer/producer.go`
- **Consumer**: Line 28 in `consumer/consumer.go`

### API Port

The producer API server runs on port `3000` by default. To change it, modify line 25 in `producer/producer.go`.

## Dependencies

- [Fiber v2](https://github.com/gofiber/fiber) - Web framework
- [IBM Sarama](https://github.com/IBM/sarama) - Kafka client library

## Building

To build the binaries:

```bash
# Build producer
go build -o producer ./producer

# Build consumer
go build -o consumer ./consumer
```

Run the binaries:

```bash
./producer
./consumer
```

## License

This is a learning project.
