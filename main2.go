package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"10.249.4.155:9092"},
		Topic:     "kafkatest",
		Partition: 0,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})
	r.SetOffset(43)
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s \n", m.Offset, string(m.Key), m.Value)
	}
	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
