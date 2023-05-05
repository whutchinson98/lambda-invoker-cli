package invoker

import (
	"fmt"
	"testing"
)

func TestGetPayloadFromConfig(t *testing.T) {
	cases := []struct {
		testCase        string
		config          *LambdaInvokerConfig
		payloadExpected *RequestPayload
		err             error
	}{
		{
			testCase: "Gets payload",
			config: &LambdaInvokerConfig{
				WorkingDir: "./",
				Request:    "test_payload.json",
			},
			payloadExpected: &RequestPayload{
				Body:    map[string]string{"hello": "world"},
				Headers: map[string]string{"Authorization": "Bearer test"},
			},
			err: nil,
		},
		{
			testCase: "request is not found",
			config: &LambdaInvokerConfig{
				WorkingDir: "./",
				Request:    "bad_payload.json",
			},
			payloadExpected: nil,
			err:             fmt.Errorf("open .//bad_payload.json: no such file or directory"),
		},
	}
	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			payload, err := GetPayloadFromConfig(c.config)

			if c.err == nil {
				if err != nil {
					t.Errorf("expected no error but got error=[%s]\n", err.Error())
				}
				if !arePayloadsEqual(payload, c.payloadExpected) {
					t.Errorf("expected payload=[%v] but got payload=[%v]\n", c.payloadExpected, payload)
				}
			}

			if c.err != nil {
				if err == nil {
					t.Errorf("expected error=[%s] but got no error\n", c.err.Error())
				}
				if c.err.Error() != err.Error() {
					t.Errorf("wanted error=[%s] got error=[%s]\n", c.err.Error(), err.Error())
				}
			}
		})
	}
}

func TestCreateAPIGatewayProxyRequest(t *testing.T) {
	cases := []struct {
		testCase       string
		requestPayload *RequestPayload
		result         *APIGatewayProxyRequest
		err            error
	}{
		{
			testCase: "Creates request",
			requestPayload: &RequestPayload{
				Body:    map[string]string{"hello": "world"},
				Headers: map[string]string{"Authorization": "test"},
			},
			result: &APIGatewayProxyRequest{
				Body:    "{\"hello\":\"world\"}",
				Headers: map[string]string{"Authorization": "test"},
			},
			err: nil,
		},
	}
	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			response, err := CreateAPIGatewayProxyRequest(c.requestPayload)

			if c.err == nil {
				if err != nil {
					t.Errorf("expected no error but got error=[%s]\n", err.Error())
				}
				if !areApiGatewayProxyRequestsEqual(response, c.result) {
					t.Errorf("expected request=[%v] but got request=[%v]\n", c.result, response)
				}
			}

			if c.err != nil {
				if err == nil {
					t.Errorf("expected error=[%s] but got no error\n", c.err.Error())
				}
				if c.err.Error() != err.Error() {
					t.Errorf("wanted error=[%s] got error=[%s]\n", c.err.Error(), err.Error())
				}
			}
		})
	}
}

func arePayloadsEqual(got *RequestPayload, expect *RequestPayload) bool {
	if len(got.Body) != len(expect.Body) {
		return false
	}

	if len(got.Headers) != len(expect.Headers) {
		return false
	}

	for k, v := range got.Body {
		if expect.Body[k] != v {
			return false
		}
	}

	for k, v := range got.Headers {
		if expect.Headers[k] != v {
			return false
		}
	}

	return true
}

func areApiGatewayProxyRequestsEqual(got *APIGatewayProxyRequest, expect *APIGatewayProxyRequest) bool {
	if got.Body != expect.Body {
		return false
	}

	if len(got.Headers) != len(expect.Headers) {
		return false
	}

	for k, v := range got.Headers {
		if expect.Headers[k] != v {
			return false
		}
	}

	return true
}
