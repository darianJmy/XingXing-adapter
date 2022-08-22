package options

import (
	"context"
	"fmt"
	pixiuConfig "github.com/caoyingjunz/pixiulib/config"
	"github.com/darianJmy/XingXing-adapter/cmd/app/config"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

const (
	defaultConfigFile = "/etc/xingxing-adapter/config.yaml"
)

type Options struct {
	ConfigFile      string
	ComponentConfig config.Config

	GinEngine *gin.Engine

	Database *mongo.Database
	Conn     *kafka.Conn
}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) Complete() error {
	configFile := o.ConfigFile
	if len(configFile) == 0 {
		configFile = os.Getenv("ConfigFile")
	}
	if len(configFile) == 0 {
		configFile = defaultConfigFile
	}

	c := pixiuConfig.New()
	c.SetConfigFile(configFile)
	c.SetConfigType("yaml")
	if err := c.Binding(&o.ComponentConfig); err != nil {
		return err
	}

	// 注册依赖组件
	if err := o.register(); err != nil {
		return err
	}
	return nil

}

func (o *Options) register() error {
	if err := o.InitKafkaClient(); err != nil {
		return err
	}

	if err := o.InitMongoClient(); err != nil {
		return err
	}

	o.GinEngine = gin.Default()
	return nil
}

func (o *Options) InitKafkaClient() error {

	topic := o.ComponentConfig.Kafka.Topic
	partition := o.ComponentConfig.Kafka.Partition
	address := fmt.Sprintf("%s:%d", o.ComponentConfig.Kafka.Host, o.ComponentConfig.Kafka.Port)

	conn, err := kafka.DialLeader(context.Background(), "tcp", address, topic, partition)
	if err != nil {
		return err
	}
	o.Conn = conn

	return nil
}

func (o *Options) InitMongoClient() error {
	address := fmt.Sprintf("mongodb://%s:%s@%s:%d",
		o.ComponentConfig.Mongo.User,
		o.ComponentConfig.Mongo.Password,
		o.ComponentConfig.Mongo.Host,
		o.ComponentConfig.Mongo.Port)

	client, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI(address).SetConnectTimeout(5*time.Second))
	if err != nil {
		return err
	}

	db := client.Database(o.ComponentConfig.Mongo.DataBase)

	o.Database = db

	return nil
}

func (o *Options) Run() {
	_ = o.GinEngine.Run(fmt.Sprintf(":%d", o.ComponentConfig.Default.Listen))
}
