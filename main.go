package main

import (
	"github.com/darianJmy/XingXing-adapter/cmd"
	"github.com/darianJmy/XingXing-adapter/options"
	"github.com/gin-gonic/gin"
)

func main() {
	s := NewHttpServer(":8081")
	s.Run()
}

func NewHttpServer(addr string) *options.Options {
	o := options.Options{
		Address: addr,
		Engine:  gin.Default(),
		Conn:    cmd.InitKafkaClient(),
	}

	o.Engine.POST("/", o.SetBodyToKafka)

	return &o
}
