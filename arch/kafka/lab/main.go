package main

import (
	"context"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	err := runConsumer(ctx, "localhost", "myGroup", "myTopic2")
	if err != nil {
		panic(err)
	}

	err = runConsumer2(ctx, "localhost", "myGroup2", "myTopic2")
	if err != nil {
		panic(err)
	}

	err = runProducer(ctx, "localhost", "myTopic2")
	if err != nil {
		panic(err)
	}

	time.Sleep(60 * time.Second)
	cancel()
}
