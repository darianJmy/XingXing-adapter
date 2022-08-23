package controller

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdapterV1Interface interface {
	WriteMessagesToMongo(ctx context.Context)
	WriteMessagesToKafka(ctx context.Context)
	HandleMessages(ctx context.Context, body []byte)
}

type Adapter struct {
	Ch1             chan []byte
	Ch2             chan AlertManager
	Database        *mongo.Database
	Conn            *kafka.Conn
	KafkaCollection *mongo.Collection
	MongoCollection *mongo.Collection
}

func (a *Adapter) HandleMessages(ctx context.Context, body []byte) {
	a.Ch1 <- body
}

func (a *Adapter) WriteMessagesToMongo(ctx context.Context) {
	var alertManager AlertManager
	for {
		msg := <-a.Ch1

		err := json.Unmarshal(msg, &alertManager)
		if err != nil {
			continue
		}

		iResult, err := a.MongoCollection.InsertOne(context.Background(), alertManager)
		if err != nil {
			continue
		}
		id := iResult.InsertedID.(primitive.ObjectID)
		alertManager.ID = id

		a.Ch2 <- alertManager
		alertManager.ID = primitive.ObjectID{}
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
			kafkaMessage.Vals = make(map[string]string)
			kafkaMessage.Dims = j.Labels
			kafkaMessage.Vals[j.Labels.AlertName] = j.Annotations.State
			kafkaMessage.MongoID = msg.ID

			_, err := a.KafkaCollection.InsertOne(context.Background(), kafkaMessage)
			if err != nil {
				continue
			}

			body, err := json.Marshal(kafkaMessage)
			if err != nil {
				continue
			}
			_, err = a.Conn.WriteMessages(
				kafka.Message{Value: body})
			if err != nil {
				continue
			}
		}
	}
}
