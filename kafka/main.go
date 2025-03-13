package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	reader *kafka.Reader
	topic  = "user_click"
)

func writeKafka(ctx context.Context) {
	fmt.Println("writeKafka is staring...")

	writer := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092", "localhost:9992"),
		Topic:                  topic,
		Balancer:               &kafka.Hash{},
		WriteTimeout:           1 * time.Second,
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: true,
	}
	defer writer.Close()
	defer fmt.Println("writeKafka is ending...")

	for i := 0; i < 3; i++ {
		if err := writer.WriteMessages(
			ctx,
			kafka.Message{Key: []byte("1"), Value: []byte("big")},
			kafka.Message{Key: []byte("2"), Value: []byte("small")},
			kafka.Message{Key: []byte("3"), Value: []byte("small")},
			kafka.Message{Key: []byte("1"), Value: []byte("bigger")},
			kafka.Message{Key: []byte("1"), Value: []byte("smaller")},
		); err != nil {
			if err == kafka.LeaderNotAvailable {
				time.Sleep(500 * time.Millisecond)
				continue
			} else {
				fmt.Printf("write to kafka failed: %v\n", err)
			}
		} else {
			break
		}
	}
}

func readKafka(ctx context.Context) {
	fmt.Println("readKafka is staring...")
	reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"localhost:9092", "localhost:9992"},
		Topic:          topic,
		CommitInterval: 1 * time.Second,
		GroupID:        "rec_team",
		StartOffset:    kafka.FirstOffset,
	})
	// defer reader.Close()
	defer fmt.Println("readKafka is ending...")

	for {
		if message, err := reader.ReadMessage(ctx); err != nil {
			fmt.Printf("read kafka failed :%v", err)
			break
		} else {
			fmt.Printf("topic=%s, parition=%d, offset=%d, key=%s, value=%s\n", message.Topic, message.Partition, message.Offset, string(message.Key), string(message.Value))
		}
	}
}

func listenSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c
	fmt.Printf("get signal=%s", sig.String())
	if reader != nil {
		reader.Close()
	}
	os.Exit(0)
}

func main() {
	ctx := context.Background()
	writeKafka(ctx)

	time.Sleep(2 * time.Second)

	go listenSignal()
	readKafka(ctx)
}
