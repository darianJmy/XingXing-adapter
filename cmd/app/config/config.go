package config

type Config struct {
	Default DefaultOptions `yaml:"default"`
	Mongo   MongoOptions   `yaml:"mongo"`
	Kafka   KafkaOptions   `yaml:"kafka"`
}

type DefaultOptions struct {
	Listen   int    `yaml:"listen"`
	LogDir   string `yaml:"log_dir"`
	LogLevel string `yaml:"log_level"`
}

type MongoOptions struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	DataBase string `yaml:"database"`
}

type KafkaOptions struct {
	Topic     string `yaml:"topic"`
	Partition int    `yaml:"partition"`
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
}
