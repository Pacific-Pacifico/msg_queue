package main

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

var (
	amqpConn *amqp.Connection
	ch       *amqp.Channel
	err      error
)

func LogIfError(fatal bool, err error, msg string) {
	if err != nil {
		if fatal {
			log.Fatalf("Error: %s \nmsg: %s", err, msg)
		} else {
			log.Printf("Error: %s \nmsg: %s", err, msg)
		}
	}
}

func getAmqpConnection() {
	rmq_env := os.Getenv("RMQ_ENV")
	amqpConn, err = amqp.Dial(rmq_env)
	LogIfError(true, err, "Error connecting to RabbitMQ instance")
	// defer amqpConn.Close()
	fmt.Println("connected to rabbitmq instance")
}

func OpenChannel() {
	ch, err = amqpConn.Channel()
	LogIfError(true, err, "Failed to open a channel")
	// defer ch.Close()
}

func main() {
	getAmqpConnection()
	OpenChannel()
	addToQueue("sample.mp4")
}
