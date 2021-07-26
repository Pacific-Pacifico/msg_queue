package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/streadway/amqp"
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

func main() {
	rmq_env := os.Getenv("RMQ_ENV")
	conn, err := amqp.Dial(rmq_env)
	LogIfError(true, err, "Error connecting to RabbitMQ instance")
	defer conn.Close()

	fmt.Println("connected to rabbitmq instance")

	ch, err := conn.Channel()
	LogIfError(true, err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TaskQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	LogIfError(true, err, "Failed to declare Queue")
	fmt.Println(q)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	str := fmt.Sprintf("%s %d", "Hello world", r1.Intn(100))
	err = ch.Publish(
		"",
		"TaskQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(str),
		},
	)
	LogIfError(true, err, "Failed to publish to work queue")
	fmt.Println("Message published:", str)
}
