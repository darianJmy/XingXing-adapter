package main

import (
	"context"
	"github.com/darianJmy/XingXing-adapter/cmd/app/options"
	"github.com/darianJmy/XingXing-adapter/pkg/adapter"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	s := options.NewOptions()

	if err := s.Complete(); err != nil {
		panic(err)
	}
	for i := 0; i < 5; i++ {
		go adapter.AdapterV1.WriteMessagesToMongo(context.Background())
	}

	for i := 0; i < 5; i++ {
		go adapter.AdapterV1.WriteMessagesToKafka(context.Background())
	}
	s.Run()
}
