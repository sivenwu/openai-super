package examples

import (
	"fmt"
	"io"
	"net/http"
	"super_openai"
)

// 常量配置
var (
	apiBaseUri string = "https://api.openai.com"
	apiKey     string = ""
	textModel  string = "gpt-3.5-turbo"
	imageModel string = "dall-e-3"
)

func getClient() *openai_super.Client {
	c := openai_super.NewClient(&openai_super.Options{
		BaseUrl:    apiBaseUri,
		ApiKey:     apiKey,
		TextModel:  textModel,
		ImageModel: imageModel,
	})
	return c
}

func outputResResultForConsole(res *http.Response, err error) {
	if err != nil {
		fmt.Printf("request api failed: %v", err)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("request api failed: %v", err)
		return
	}

	fmt.Printf("res is %s", string(body))
}
