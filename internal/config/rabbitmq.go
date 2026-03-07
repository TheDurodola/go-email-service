package config

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbitMQConnection() *amqp.Connection {
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		log.Fatal("FATAL: RABBITMQ_URL environment variable is missing")
	}

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("FATAL: Failed to connect to RabbitMQ: %v", err)
	}

	log.Println("RabbitMQ connection successfully established.")
	return conn
}