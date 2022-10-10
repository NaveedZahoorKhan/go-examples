package main

import (
	"context"
	"log"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

func main() {
	topic := "go-examples"

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, 0)
	if err != nil {
		log.Fatal(err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(time.Now().String())},
		kafka.Message{Value: []byte(time.Now().String())},
		kafka.Message{Value: []byte(time.Now().String())},
		kafka.Message{Value: []byte(time.Now().String())},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
