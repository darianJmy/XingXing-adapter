package controller

import "time"

type KafkaMessage struct {
	Source Source    `json:"source"`
	Dims   Dims      `json:"dims"`
	Vals   Vals      `json:"vals"`
	Time   time.Time `json:"time"`
}

type Source struct {
	Name string `json:"name"`
	Key  string `json:"key"`
	Org  string `json:"org"`
}

type Dims struct {
	IP string `json:"ip"`
}

type Vals struct {
	AlertName string `json:"alertname"`
}

type AlertManager struct {
	Receiver string `json:"receiver"`
	Status   string `json:"status"`
	Alerts   []Alerts
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
