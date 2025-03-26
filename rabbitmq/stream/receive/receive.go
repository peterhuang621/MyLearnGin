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

	messageHandler := func(consumerContext stream.ConsumerContext, message *amqp.Message) {
		fmt.Printf("Stream: %s - Received message: %s\n", consumerContext.Consumer.GetStreamName(),
			message.Data)
	}

	consumer, err := env.NewConsumer(streamName, messageHandler, stream.NewConsumerOptions().SetOffset(stream.OffsetSpecification{}.First()))
	failedonerr(err, "create a consumer")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println(" [x] Press enter to close the producer")
	_, _ = reader.ReadString('\n')
	err = consumer.Close()
	failedonerr(err, "close the consumer")
}
