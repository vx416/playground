package main

import (
	"context"
	"fmt"
	"log"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func runConsumer2(ctx context.Context, server, groupID string, topics ...string) error {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        server,
		"group.id":                 groupID,
		"group.instance.id":        "c_2",
		"auto.offset.reset":        "earliest",
		"enable.auto.offset.store": "false",
		"enable.auto.commit":       "true",
		"auto.commit.interval.ms":  "1",
	})

	if err != nil {
		return err
	}

	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		return err
	}

	go func(ctx context.Context, consumer *kafka.Consumer) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				msg, err := consumer.ReadMessage(-1)

				if err == nil {
					fmt.Printf("Message 2 on %s: %s, %d\n", msg.TopicPartition, string(msg.Value), msg.TopicPartition.Offset)
				} else {
					fmt.Printf("Consumer error: %v (%v)\n", err, msg)
				}
				msg.TopicPartition.Offset = msg.TopicPartition.Offset + 1
				offset := append([]kafka.TopicPartition{}, msg.TopicPartition)
				consumer.StoreOffsets(offset)
			}
		}
	}(ctx, consumer)
	return nil
}

func runConsumer(ctx context.Context, server, groupID string, topics ...string) error {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        server,
		"group.id":                 groupID,
		"group.instance.id":        "c_1",
		"auto.offset.reset":        "earliest",
		"enable.auto.offset.store": "false",
		"enable.auto.commit":       "true",
		"auto.commit.interval.ms":  "1",
		"queued.min.messages":      "11",
	})

	if err != nil {
		return err
	}

	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		return err
	}

	go func(ctx context.Context, consumer *kafka.Consumer) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				msg, err := consumer.ReadMessage(-1)
				offset := append([]kafka.TopicPartition{}, msg.TopicPartition)
				consumer.StoreOffsets(offset)
				// fmt.Printf("res:%+v", res)
				if err == nil {
					fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
				} else {
					fmt.Printf("Consumer error: %v (%v)\n", err, msg)
				}
			}
		}
	}(ctx, consumer)
	return nil
}

func runProducer(ctx context.Context, server string, topics ...string) error {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": server,
	})
	if err != nil {
		return err
	}
	defer producer.Close()
	admin, err := kafka.NewAdminClientFromProducer(producer)
	if err != nil {
		return err
	}
	topicSpec := kafka.TopicSpecification{
		Topic:             topics[0],
		NumPartitions:     2,
		ReplicationFactor: 1,
	}
	res, err := admin.CreateTopics(ctx, []kafka.TopicSpecification{topicSpec})
	if err != nil {
		return err
	}
	log.Printf("topic:%+v\n", res)
	result, err := admin.CreatePartitions(ctx, []kafka.PartitionsSpecification{{
		Topic:      topics[0],
		IncreaseTo: 2}})
	log.Printf("create partition:%+v\n", result)

	// producer.BeginTransaction
	// Delivery report handler for produced messages
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	topic := topics[0]
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: 1},
			Value:          []byte(word),
			Headers:        []kafka.Header{{Key: "test"}},
		}, nil)
	}

	producer.Flush(15 * 1000)

	return nil
}
