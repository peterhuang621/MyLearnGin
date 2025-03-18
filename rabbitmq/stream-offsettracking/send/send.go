package main

import (
	"fmt"
	"log"

	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"
)

func failedonerr(err error, str string) {
	if err != nil {
		log.Fatalf("failed to %s: %v", str, err)
	}
}

func handlePublishConfirm(confirms stream.ChannelPublishConfirm, messageCount int, ch chan bool) {
	go func() {
		confirmedCount := 0
		for confirmed := range confirms {
			for _, msg := range confirmed {
				if msg.IsConfirmed() {
					confirmedCount++
					if confirmedCount == messageCount {
						ch <- true
					}
				}
			}
		}
	}()
}

func main() {
	env, err := stream.NewEnvironment(stream.NewEnvironmentOptions().
		SetHost("localhost").
		SetPort(5552).
		SetUser("guest").
		SetPassword("guest"))
	failedonerr(err, "create a new environment")
	streamName := "stream-offset-tracking-go"
	env.DeclareStream(streamName, &stream.StreamOptions{MaxLengthBytes: stream.ByteCapacity{}.GB(2)})

	producer, err := env.NewProducer(streamName, stream.NewProducerOptions())
	failedonerr(err, "create producer")

	messageCount := 100
	ch := make(chan bool)
	chPublishConfirm := producer.NotifyPublishConfirmation()
	handlePublishConfirm(chPublishConfirm, messageCount, ch)

	fmt.Printf("Publishing %d messages\n", messageCount)
	for i := 0; i < messageCount; i++ {
		var body string
		if i == messageCount-1 {
			body = "marker"
		} else {
			body = "hello"
		}
		producer.Send(amqp.NewMessage([]byte(body)))
	}
	_ = <-ch
	fmt.Println("Messages confirmed")

	err = producer.Close()
	failedonerr(err, "close the producer")
}
