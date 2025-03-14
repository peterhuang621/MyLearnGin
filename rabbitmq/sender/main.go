package main

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "failed to connect rabbitmq")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "failed to create the channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("hello", true, false, false, false, nil)
	failOnError(err, "failed to create the queue")

	body := "this is a message aaa"
	ch.PublishWithContext(context.Background(), "", q.Name, false, false, amqp.Publishing{ContentType: "text/pain", Body: []byte(body)})
}
