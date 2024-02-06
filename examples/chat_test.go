package examples

import (
	"fmt"
	"testing"
)

// 实例函数，模拟创建一个聊天会话集
func TestCreateChatCompletion(t *testing.T) {
	// for client
	cli := getClient()
	// for params
	params := map[string]interface{}{
		"model": cli.Opts.TextModel,
		"messages": []map[string]interface{}{
			{"role": "user", "content": "hello"},
		},
	}
	// for headers
	headers := map[string]interface{}{
		"Authorization": fmt.Sprintf("Bearer %s", cli.Opts.ApiKey),
	}

	res, err := cli.Post("/v1/chat/completions", params, headers)

	// res
	outputResResultForConsole(res, err)
}
