package main

import (
	"context"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func main() {
	p, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
		producer.WithRetry(2),
	)
	if err != nil {
		panic(err)
	}

	err1 := p.Start()
	if err1 != nil {
		panic(err1)
	}

	res, err2 := p.SendSync(context.Background(), primitive.NewMessage("test", []byte("Hello Rocketmq")))
	if err2 != nil {
		fmt.Printf("Send failed: %s\n", err2)
	} else {
		fmt.Printf("Send successfully: %s\n", res.String())
	}
	p.Shutdown()

	c, _ := rocketmq.NewPushConsumer(consumer.WithNameServer([]string{"127.0.0.1:9876"}),
		consumer.WithGroupName("testGroup"),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromFirstOffset),
	)
	c.Subscribe("test", consumer.MessageSelector{}, func(ctx context.Context, me ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range me {
			fmt.Printf("Message received: %s\n", me[i].Body)
		}
		return consumer.ConsumeSuccess, nil
	})

	err3 := c.Start()
	if err3 != nil {
		panic(err3)
	}
}
