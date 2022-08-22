package adapter

import (
	"github.com/darianJmy/XingXing-adapter/cmd/app/options"
	"github.com/darianJmy/XingXing-adapter/cmd/controller"
)

var AdapterV1 controller.AdapterV1Interface

// Setup 完成核心应用接口的设置
func Setup(o *options.Options) {
	AdapterV1 = &controller.Adapter{
		Ch1:        make(chan []byte),
		Ch2:        make(chan []byte),
		Collection: o.Collection,
		Conn:       o.Conn,
	}
}
