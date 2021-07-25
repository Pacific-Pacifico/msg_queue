package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Rabbitmq")
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

	q, err := ch.QueueDeclare(
		"TaskQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

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

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Message published:", str)
}
