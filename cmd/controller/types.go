package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type KafkaMessage struct {
	Source  Source             `json:"source"`
	Dims    Labels             `json:"dims"`
	Vals    map[string]string  `json:"vals"`
	Time    time.Time          `json:"time"`
	MongoID primitive.ObjectID `json:"mongoid"`
}

type Source struct {
	Name string `json:"name"`
	Key  string `json:"key"`
	Org  string `json:"org"`
}

type AlertManager struct {
	Receiver string `json:"receiver"`
	Status   string `json:"status"`
	Alerts   []Alerts
	ID       primitive.ObjectID `json:"id"`
}

type Alerts struct {
	Status      string      `json:"status"`
	Labels      Labels      `json:"labels"`
	Annotations Annotations `json:"annotations"`
}

type Labels struct {
	AlertName string `json:"alertname"`
	Instance  string `json:"instance"`
	Job       string `json:"job"`
}

type Annotations struct {
	Description string `json:"description"`
	State       string `json:"state"`
}
