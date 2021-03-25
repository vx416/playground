package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"runtime"
	"strconv"

	"github.com/Shopify/sarama"
)

func NewConsumerGroup(groupID string, address ...string) (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()
	config.ClientID = "test"
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Group.Member.UserData = []byte("test_member_1")
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Net.MaxOpenRequests = 5
	config.RackID = "testinstanceid"

	return sarama.NewConsumerGroup(address, groupID, config)
}

func ListenOn(topics []string, consumer sarama.ConsumerGroup) error {

	c := &Consumer{ready: make(chan bool)}
	ctx, _ := context.WithCancel(context.Background())
	// sarama.Logger = log.New(os.Stdout, fmt.Sprintf("[%s]", "consumer"), log.LstdFlags)

	go func() {
		for err := range consumer.Errors() {
			log.Printf("Sarama err:%+v", err)
		}
	}()
	go func() {
		for {
			err := consumer.Consume(ctx, topics, c)
			log.Printf("err:%+v", err)
		}
	}()

	// cancel()

	// c.Ready()
	log.Println("Sarama consumer up and running!...")
	return nil
}

type Consumer struct {
	ready chan bool
}

func (consumer *Consumer) Ready() {
	<-consumer.ready
}

func (consumer *Consumer) Setup(seesion sarama.ConsumerGroupSession) error {
	// close(consumer.ready)
	log.Println("Sarama consumer setup")

	return nil
}

func (consumer *Consumer) Cleanup(seesion sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(seesion sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	fmt.Printf("goroutine id:%d\n", GetGID())

	for {
		select {
		case message := <-claim.Messages():
			log.Printf("Message claimed: value = %s, offset = %d, topic = %s", string(message.Value), message.Offset, message.Topic)
			// log.Printf("Message claimed: header:%+v", message.Headers)
			// default:
			seesion.MarkMessage(message, "")

			// time.Sleep(30 * time.Microsecond)
		}
	}

}

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
