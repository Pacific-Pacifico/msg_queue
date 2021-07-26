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
	LogIfError(true, err, "Error connecting to RabbitMQ instance")
	defer ch.Close()

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	LogIfError(true, err, "Failed to set QoS")
	msgs, err := ch.Consume(
		"TaskQueue",
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved msg:%s\n", d.Body)
			time.Sleep(time.Second * time.Duration(r1.Intn(10)))
			fmt.Println("Processing complete!")
			d.Ack(false)
		}
	}()

	fmt.Println("Successfully connected to Rabbitmq instance")
	fmt.Println("~~waiting for msg~~")
	<-forever
}
