# Project Overview

## What We Are Trying to Do

This project is a **learning exercise** that demonstrates how to build a **Kafka-based message queue system** using Go. The project implements a simple **producer-consumer pattern** where:

1. **Producer**: A REST API that receives HTTP requests and publishes messages to Kafka
2. **Consumer**: A service that reads messages from Kafka and processes them

## Learning Objectives

### 1. Understanding Kafka Basics

- Learn how to connect to a Kafka broker
- Understand topics, partitions, and offsets
- Practice producing and consuming messages

### 2. Building REST APIs with Fiber

- Create HTTP endpoints using the Fiber web framework
- Handle JSON request/response parsing
- Implement proper error handling

### 3. Producer-Consumer Pattern

- Understand asynchronous message processing
- Learn how to decouple services using message queues
- See how producers and consumers can run independently

### 4. Go Concurrency

- Use goroutines for concurrent message processing
- Handle graceful shutdown with signal handling
- Manage channels for communication

## How It Works

### The Flow

```
Client Request → Producer API → Kafka Topic → Consumer → Display Message
```

1. **Client sends a POST request** to `/api/v1/comments` with a comment text
2. **Producer receives the request**, validates it, and publishes it to the "comments" Kafka topic
3. **Kafka stores the message** in the topic
4. **Consumer reads the message** from Kafka and displays it

### Key Concepts Demonstrated

- **Message Queue**: Kafka acts as a buffer between the producer and consumer
- **Decoupling**: The producer doesn't need to know about the consumer, and vice versa
- **Scalability**: Multiple consumers can read from the same topic
- **Reliability**: Messages are persisted in Kafka, so they won't be lost if a consumer crashes

## Real-World Applications

This pattern is commonly used in:

- **Microservices Architecture**: Services communicate through message queues
- **Event-Driven Systems**: Events are published and consumed asynchronously
- **Log Aggregation**: Collecting logs from multiple sources
- **Real-time Analytics**: Processing streams of data
- **Notification Systems**: Sending notifications without blocking the main application

## What You'll Learn

By working through this project, you'll gain experience with:

- ✅ Go programming fundamentals
- ✅ REST API development
- ✅ Kafka integration
- ✅ Error handling and logging
- ✅ Concurrent programming in Go
- ✅ Building distributed systems

This is a foundational project that introduces concepts you'll use in larger, more complex systems.
