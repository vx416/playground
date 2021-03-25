package main

import (
	"log"

	"github.com/Shopify/sarama"
)

func NewPublisher(address ...string) (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	config.ClientID = "testing"
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	return sarama.NewAsyncProducer(address, config)
}

func Publish(topics []string, produer sarama.AsyncProducer) error {
	messages := []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"}
	log.Print("publish start")
	// wg := sync.WaitGroup{}
	// wg.Add(1)
	go func() {
		for msg := range produer.Successes() {
			log.Printf("producer: publish message on topic(%s)-(%d)", msg.Topic, msg.Partition)
		}
		// wg.Done()
	}()

	go func() {
		for err := range produer.Errors() {
			log.Printf("producer: publish message on topic(%s)", err.Err)
			produer.Input() <- err.Msg
		}
		// wg.Done()
	}()

	// wg.Wait()

	for _, topic := range topics {
		for _, msg := range messages {
			pm := &sarama.ProducerMessage{}
			pm.Topic = topic
			pm.Value = sarama.StringEncoder(msg)
			produer.Input() <- pm
			log.Print("published!!!")

		}
	}

	return nil
}
