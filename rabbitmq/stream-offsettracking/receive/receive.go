package main

import (
	"errors"
	"fmt"
	"log"
	"sync/atomic"
	"time"

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
	streamName := "stream-offset-tracking-go"
	env.DeclareStream(streamName, &stream.StreamOptions{MaxLengthBytes: stream.ByteCapacity{}.GB(2)})

	var firstOffset, messageCount int64 = -1, -1
	var lastOffset atomic.Int64

	ch := make(chan bool)
	messageHandler := func(consumerContext stream.ConsumerContext, message *amqp.Message) {
		time.Sleep(100 * time.Millisecond)
		if atomic.CompareAndSwapInt64(&firstOffset, -1, consumerContext.Consumer.GetOffset()) {
			fmt.Println("First message received.")
		}
		if atomic.AddInt64(&messageCount, 1)%10 == 0 {
			_ = consumerContext.Consumer.StoreOffset()
		}
		if string(message.GetData()) == "marker" {
			lastOffset.Store(consumerContext.Consumer.GetOffset())
			_ = consumerContext.Consumer.StoreOffset()
			_ = consumerContext.Consumer.Close()
			ch <- true
		}
	}

	var OffsetSpecification stream.OffsetSpecification
	consumerName := "offset-tracking-tutorial"
	storedOffset, err := env.QueryOffset(consumerName, streamName)
	if errors.Is(err, stream.OffsetNotFoundError) {
		OffsetSpecification = stream.OffsetSpecification{}.First()
	} else {
		OffsetSpecification = stream.OffsetSpecification{}.Offset(storedOffset + 1)
	}

	_, err = env.NewConsumer(streamName, messageHandler, stream.NewConsumerOptions().
		SetManualCommit().
		SetConsumerName(consumerName).
		SetOffset(OffsetSpecification))
	failedonerr(err, "create a consumer")

	fmt.Println("Started consuming...")
	_ = <-ch
	fmt.Printf("Done consuming, first offset %d, last offset %d.\n", firstOffset, lastOffset.Load())
}
