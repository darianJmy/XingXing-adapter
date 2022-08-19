package main

import (
	"github.com/darianJmy/XingXing-adapter/options"
)

func main() {

	s := options.NewHttpServer(":8081")

	for i := 0; i < 5; i++ {
		go s.WriteMessagesToMongo()
	}

	for i := 0; i < 5; i++ {
		go s.WriteMessagesToKafka()
	}

	s.Run()
}
