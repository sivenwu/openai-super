package openai_super

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Client 结构体
type Client struct {
	Opts          *Options
	RequestClient *http.Client
}

// NewClient 初始化客户端
func NewClient(opts *Options) *Client {
	// := 运算符可以使变量在不声明的情况下直接被赋值使用
	// & 初始化对象并且获得对象的指针地址
	cli := &Client{Opts: opts}
	cli.RequestClient = &http.Client{}

	return cli
}

// 以下是类方法

func (cli *Client) Get(url string, args ...map[string]interface{}) (*http.Response, error) {

	// set url
	url = cli.Opts.BaseUrl + url

	// parse params、headers
	params, headers := cli.parseArgs(args...)

	// payload and header
	jsonData, _ := json.Marshal(params)
	req, _ := http.NewRequest("GET", url, bytes.NewBuffer(jsonData))
	for key, value := range headers {
		req.Header.Add(key, fmt.Sprint(value))
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	return cli.getRequestClient().Do(req)
}

func (cli *Client) Post(url string, args ...map[string]interface{}) (*http.Response, error) {

	// set url
	url = cli.Opts.BaseUrl + url

	// parse params、headers
	params, headers := cli.parseArgs(args...)

	// payload and header
	jsonData, _ := json.Marshal(params)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(jsonData)))
	for key, value := range headers {
		req.Header.Add(key, fmt.Sprint(value))
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	return cli.getRequestClient().Do(req)
}

// 解析 args，0 代表请求参数，1代表请求headers
func (cli *Client) parseArgs(args ...map[string]interface{}) (map[string]interface{}, map[string]interface{}) {
	params := map[string]interface{}{}
	headers := map[string]interface{}{}

	if len(args) > 0 {
		params = args[0]
	}

	if len(args) > 1 {
		headers = args[1]
	}

	return params, headers
}

// getRequestClient: get request handler

func (cli *Client) getRequestClient() *http.Client {
	return cli.RequestClient
}
