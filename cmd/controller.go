package cmd

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

func InitKafkaClient() *kafka.Conn {

	topic := "kafkatest"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "10.249.4.155:9092", topic, partition)
	if err != nil {
		return nil
	}

	return conn
}

func SetWriteToMessage(conn *kafka.Conn, body []byte) {
	_, err := conn.WriteMessages(kafka.Message{Value: body})
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
}
