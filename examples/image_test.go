package examples

import (
	"fmt"
	"testing"
)

// 实例函数，模拟创建一个聊天会话集
func TestCreateImagesGenerations(t *testing.T) {
	// for client
	cli := getClient()
	// for params
	params := map[string]interface{}{
		"model":  cli.Opts.ImageModel,
		"prompt": "A cute baby sea otter",
		"n":      1,
		"size":   "1024x1024",
	}
	// for headers
	headers := map[string]interface{}{
		"Authorization": fmt.Sprintf("Bearer %s", cli.Opts.ApiKey),
	}

	res, err := cli.Post("/v1/images/generations", params, headers)

	// res
	outputResResultForConsole(res, err)
}
