package router

import (
	"context"
	"fmt"
	"github.com/darianJmy/XingXing-adapter/pkg/adapter"
	"github.com/gin-gonic/gin"
)

func HandleMessages(c *gin.Context) {

	body, err := c.GetRawData()
	if err != nil {
		c.String(400, "Error")
		return
	}
	fmt.Println(string(body))
	adapter.AdapterV1.HandleMessages(context.Background(), body)
	c.String(200, "Success")
}
