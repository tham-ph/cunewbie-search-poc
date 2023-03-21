package database

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectRabbitMQ() (*amqp.Connection, error) {
	return amqp.Dial("amqp://guest:guest@localhost:5672/")
}
