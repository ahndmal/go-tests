package net

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
	"testing"
	"time"
)

const url = "pulsar://172.17.0.3:6650"
const topicName = "main"

func TestPulsar(t *testing.T) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               url,
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Printf("Error when ini client: %v", err)
	}
	defer client.Close()

	producer, err := client.CreateProducer(pulsar.ProducerOptions{Topic: topicName})
	if err != nil {
		log.Printf("Error when creating Producer: %v", err)
	}

	message := "Hello from GO"
	sentMsg, err := producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(message),
	})
	fmt.Printf(">> Sent message %s\n", sentMsg)
	if err != nil {
		log.Printf("Error when sending message in Producer: %v", err)
	}
	defer producer.Close()
}

func TestGetMessages(t *testing.T) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               url,
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Printf("Error when init client: %v", err)
	}
	defer client.Close()

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topicName,
		SubscriptionName: "sub1",
		Type:             pulsar.Shared,
	})
	if err != nil {
		log.Fatalf("Error when creating Consumer: %v", err)
	}
	defer consumer.Close()

	// 1st msg
	msg, err := consumer.Receive(context.Background())
	fmt.Println(string(msg.Payload()))

	// 2nd timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	msg, err = consumer.Receive(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(msg.Payload()))
	cancel()

}
