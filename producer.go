package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func addToQueue(fileName string) {
	q, err := ch.QueueDeclare(
		"TranscodeQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	LogIfError(true, err, "Failed to declare Queue")
	fmt.Println(q)

	err = ch.Publish(
		"",
		"TranscodeQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(fileName),
		},
	)
	LogIfError(true, err, "Failed to publish to work queue")
	fmt.Println("Message published:", fileName)
}
