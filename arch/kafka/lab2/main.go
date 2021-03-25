package main

import (
	"log"
	"os"
	"time"
)

func main() {
	address := "localhost:9092"
	topics := []string{"test_topic_1", "test_topic_2"}

	client, err := NewConsumerGroup("test_group", address)
	if err != nil {
		log.Panicf("err:%+v", err)
		os.Exit(1)
	}

	err = ListenOn(topics, client)
	if err != nil {
		log.Panicf("err:%+v", err)
	}
	log.Print("listen on")

	p, err := NewPublisher(address)
	if err != nil {
		log.Panicf("err:%+v", err)
		os.Exit(1)
	}
	err = Publish(topics, p)
	if err != nil {
		log.Panicf("err:%+v", err)
		os.Exit(1)
	}

	time.Sleep(15 * time.Second)
}
