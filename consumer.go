package main

import (
	"fmt"
	"os"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer")
	rmq_env := os.Getenv("RMQ_ENV")
	conn, err := amqp.Dial(rmq_env)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	fmt.Println("connected to rabbitmq instance")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"TaskQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved msg:%s\n", d.Body)
			time.Sleep(time.Second * 5)
		}
	}()

	fmt.Println("Successfully connected to Rabbitmq instance")
	fmt.Println("~~waiting for msg~~")
	<-forever
}
