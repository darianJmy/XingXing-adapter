package options

import (
	"fmt"
	"github.com/darianJmy/XingXing-adapter/cmd"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

type Options struct {
	Engine  *gin.Engine
	Address string
	Conn    *kafka.Conn
}

func (o *Options) Run() {
	o.Engine.Run(o.Address)
}

func (o *Options) SetBodyToKafka(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		c.String(400, "Error")
		return
	}
	fmt.Println(string(body))
	if o.Conn == nil {
		c.String(400, "Error")
		return
	}
	go cmd.SetWriteToMessage(o.Conn, body)
	c.String(200, "Success")
}
