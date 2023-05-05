package invoker

import (
	"encoding/json"
	"fmt"
	"os"
)

type APIGatewayProxyRequest struct {
	Body    string            `json:"body"`
	Headers map[string]string `json:"headers"`
}

type RequestPayload struct {
	Body    map[string]string `json:"body"`
	Headers map[string]string `json:"headers"`
}

func CreateAPIGatewayProxyRequest(requestPayload *RequestPayload) (*APIGatewayProxyRequest, error) {
	bodyStr, err := json.Marshal(requestPayload.Body)

	if err != nil {
		return nil, err
	}

	return &APIGatewayProxyRequest{
		Body:    string(bodyStr),
		Headers: requestPayload.Headers,
	}, nil
}

func GetPayloadFromConfig(config *LambdaInvokerConfig) (*RequestPayload, error) {
	filePath := fmt.Sprintf("%s/%s", config.WorkingDir, config.Request)
	val, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	requestPayload := &RequestPayload{}

	err = json.Unmarshal(val, &requestPayload)

	if err != nil {
		return nil, err
	}

	return requestPayload, nil
}
