package examples

import (
	"fmt"
	"testing"
)

func TestCreateCompletions(t *testing.T) {

	cli := getClient()

	params := map[string]interface{}{
		"model":       "gpt-3.5-turbo-instruct",
		"prompt":      "Say this is a test",
		"max_tokens":  7,
		"temperature": 0,
	}

	headers := map[string]interface{}{
		"Authorization": fmt.Sprintf("Bearer %s", cli.Opts.ApiKey),
	}

	res, err := cli.Post("/v1/completions", params, headers)
	outputResResultForConsole(res, err)
}
