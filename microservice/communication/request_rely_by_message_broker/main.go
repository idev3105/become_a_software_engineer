package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

type Request struct {
	CorrelationID string `json:"correlation_id"`
	Data          string `json:"data"`
}

type Reply struct {
	CorrelationID string `json:"correlation_id"`
	Response      string `json:"response"`
}

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Consumer.Return.Errors = true

	// create producer
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}
	defer producer.Close()

	// create consumer
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}
	defer consumer.Close()

	// send request to "request_topic"
	correlationID := fmt.Sprintf("%d", time.Now().UnixNano())
	request := Request{
		CorrelationID: correlationID,
		Data:          "Hello, Kafka!",
	}
	requestJSON, _ := json.Marshal(request)

	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic: "request_topic",
		Value: sarama.StringEncoder(requestJSON),
	})
	if err != nil {
		log.Fatalf("Failed to send message: %s", err)
	}

	// After send request, service consumer will receive and handle it
	// Then, service consumer send it to "reply_topic"

	// consume reply
	partitionConsumer, err := consumer.ConsumePartition("reply_topic", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to create partition consumer: %s", err)
	}
	defer partitionConsumer.Close()

	// set timeout for this listening
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// wait the response
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			var reply Reply
			json.Unmarshal(msg.Value, &reply)
			// if this is matched response, handle it and return
			if reply.CorrelationID == correlationID {
				fmt.Printf("Received reply: %s\n", reply.Response)
				return
			}
		case <-ctx.Done():
			// if timeout, return
			log.Println("Timeout waiting for reply")
			return
		}
	}
}
