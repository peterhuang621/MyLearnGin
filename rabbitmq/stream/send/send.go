package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"
)

func failedonerr(err error, str string) {
	if err != nil {
		log.Fatalf("failed to %s: %v", str, err)
	}
}

func main() {
	env, err := stream.NewEnvironment(stream.NewEnvironmentOptions().
		SetHost("localhost").
		SetPort(5552).
		SetUser("guest").
		SetPassword("guest"))
	failedonerr(err, "create a new environment")
	streamName := "hello-go-stream"
	env.DeclareStream(streamName, &stream.StreamOptions{MaxLengthBytes: stream.ByteCapacity{}.GB(2)})

	producer, err := env.NewProducer(streamName, stream.NewProducerOptions())
	failedonerr(err, "create producer")

	err = producer.Send(amqp.NewMessage([]byte("Hello world")))
	failedonerr(err, "send message")
	fmt.Printf(" [x] 'hello world' message sent\n")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println(" [x] Press enter to close the producer")
	_, _ = reader.ReadString('\n')
	err = producer.Close()
	failedonerr(err, "close the producer")
}
