package options

import (
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type Options struct {
	Engine     *gin.Engine
	Address    string
	Ch1        chan []byte
	Ch2        chan []byte
	Collection *mongo.Collection
	Conn       *kafka.Conn
}

type KafkaRequest struct {
	Receiver    string `json:"receiver"`
	Status      string `json:"status"`
	ExternalURL string `json:"externalURL"`
}

type AlertManager struct {
	Receiver        string `json:"receiver"`
	Status          string `json:"status"`
	ExternalURL     string `json:"externalURL"`
	Version         string `json:"version"`
	GroupKey        string `json:"groupKey"`
	TruncatedAlerts int    `json:"truncatedAlerts"`
}
