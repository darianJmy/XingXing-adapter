package controller

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type AdapterV1Interface interface {
	WriteMessagesToMongo(ctx context.Context)
	WriteMessagesToKafka(ctx context.Context)
	HandleMessages(ctx context.Context, body []byte)
}

type Adapter struct {
	Ch1        chan []byte
	Ch2        chan []byte
	Collection *mongo.Collection
	Conn       *kafka.Conn
}

func (a *Adapter) HandleMessages(ctx context.Context, body []byte) {
	a.Ch1 <- body
}

func (a *Adapter) WriteMessagesToMongo(ctx context.Context) {
	for {
		msg := <-a.Ch1

		time.Sleep(5 * time.Second)
		fmt.Println(string(msg))
		//err := json.Unmarshal(msg, &alertManager)
		//if err != nil {
		//	log.Println(err)
		//} else {
		//	fmt.Println(msg)
		//	o.Ch2 <- msg
		//}

	}
}

func (a *Adapter) WriteMessagesToKafka(ctx context.Context) {
	for {
		msg := <-a.Ch2
		//var kafkarequest KafkaRequest
		//err := json.Unmarshal(msg, &kafkarequest)
		//if err != nil {
		//	log.Fatal(err)
		//} else {
		//	body, err := json.Marshal(kafkarequest)
		//	if err != nil {
		//		log.Fatal(err)
		//	} else {
		//		_, err = o.Conn.WriteMessages(
		//			kafka.Message{Value: body})
		//		if err != nil {
		//			log.Fatal(err)
		//		}
		//	}
		//}
		fmt.Println(string(msg))
	}
}
