package openai_super

// Options 客户端配置参数
type Options struct {
	//  url 字符串
	BaseUrl string
	// apikey
	ApiKey string
	// 文字引擎版本
	TextModel string
	// 图片引擎版本
	ImageModel string
	// completions 引擎版本
	//gpt-3.5-turbo-instruct
}
