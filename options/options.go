package options

import (
	"encoding/json"
	"fmt"
	"github.com/darianJmy/XingXing-adapter/cmd"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func NewHttpServer(addr string) *Options {
	o := Options{
		Address:    addr,
		Engine:     gin.Default(),
		Ch1:        make(chan []byte, 20),
		Ch2:        make(chan []byte, 20),
		Collection: cmd.InitMongoClient(),
		Conn:       cmd.InitKafkaClient(),
	}
	o.RegisterHttpRoute()

	return &o
}

func (o *Options) Run() {
	o.Engine.Run(o.Address)
}

func (o *Options) RegisterHttpRoute() {
	o.Engine.POST("/", o.ReadMessages)
}

func (o *Options) ReadMessages(c *gin.Context) {

	body, err := c.GetRawData()
	if err != nil {
		c.String(400, "Error")
		return
	}
	o.Ch1 <- body
	c.String(200, "Success")
}

func (o *Options) WriteMessagesToMongo() {

	for {
		msg := <-o.Ch1
		var alertManager AlertManager

		time.Sleep(5 * time.Second)
		err := json.Unmarshal(msg, &alertManager)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(msg)
			o.Ch2 <- msg
		}
	}
}

func (o *Options) WriteMessagesToKafka() {
	for {
		msg := <-o.Ch2
		var kafkarequest KafkaRequest
		err := json.Unmarshal(msg, &kafkarequest)
		if err != nil {
			log.Fatal(err)
		} else {
			body, err := json.Marshal(kafkarequest)
			if err != nil {
				log.Fatal(err)
			} else {
				_, err = o.Conn.WriteMessages(
					kafka.Message{Value: body})
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
