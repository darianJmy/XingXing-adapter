package cmd

import (
	"context"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
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

func InitMongoClient() *mongo.Collection {
	client, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb://root:root@localhost:27017").SetConnectTimeout(5*time.Second))
	if err != nil {
		panic(err)
	}
	db := client.Database("prometheus")
	collection := db.Collection("new")
	return collection
}
