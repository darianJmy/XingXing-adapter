package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type AdapterV1Interface interface {
	WriteMessagesToMongo(ctx context.Context)
	WriteMessagesToKafka(ctx context.Context)
	HandleMessages(ctx context.Context, body []byte)
}

type Adapter struct {
	Ch1      chan []byte
	Ch2      chan AlertManager
	Database *mongo.Database
	Conn     *kafka.Conn
}

func (a *Adapter) HandleMessages(ctx context.Context, body []byte) {
	a.Ch1 <- body
}

func (a *Adapter) WriteMessagesToMongo(ctx context.Context) {
	var alertManager AlertManager
	for {
		msg := <-a.Ch1

		time.Sleep(5 * time.Second)
		fmt.Println(string(msg))
		err := json.Unmarshal(msg, &alertManager)
		if err != nil {
			log.Println(err)
			continue
		}
		collection := a.Database.Collection("new")
		_, err = collection.InsertOne(context.Background(), alertManager)
		if err != nil {
			continue
		}
		a.Ch2 <- alertManager
	}
}

func (a *Adapter) WriteMessagesToKafka(ctx context.Context) {
	var kafkaMessage KafkaMessage
	kafkaMessage.Source.Name = "HOST"
	kafkaMessage.Source.Org = "21211129"
	kafkaMessage.Source.Key = "host999"
	for {
		msg := <-a.Ch2

		for _, j := range msg.Alerts {
			kafkaMessage.Dims.IP = j.Labels.Instance
			kafkaMessage.Vals.AlertName = j.Annotations.Description

			body, err := json.Marshal(kafkaMessage)
			if err != nil {
				continue
			}
			_, err = a.Conn.WriteMessages(
				kafka.Message{Value: body})
			if err != nil {
				continue
			}
			collection := a.Database.Collection("old")
			_, err = collection.InsertOne(context.Background(), kafkaMessage)
			if err != nil {
				continue
			}
		}
	}
}
