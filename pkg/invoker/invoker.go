package invoker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

type LambdaInvokerConfig struct {
	LambdaName string `json:"lambdaName"`
	Request    string `json:"request"`
	Region     string `json:"region"`
	WorkingDir string `json:"workingDir"`
}

type LambdaClient interface {
	Invoke(ctx context.Context, params *lambda.InvokeInput, optFns ...func(*lambda.Options)) (*lambda.InvokeOutput, error)
}

type LambdaInvokerService struct {
	LambdaClient LambdaClient
	LambdaName   string
}

func InitializeLambdaClient(ctx context.Context, region string) (*lambda.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}

	return lambda.NewFromConfig(cfg), nil
}

func (lambdaInvokerService *LambdaInvokerService) InvokeLambda(ctx context.Context, payload []byte) (*lambda.InvokeOutput, error) {

	return lambdaInvokerService.LambdaClient.Invoke(ctx, &lambda.InvokeInput{
		FunctionName:   aws.String(lambdaInvokerService.LambdaName),
		InvocationType: types.InvocationTypeRequestResponse,
		Payload:        payload,
	})
}

func InvokeLambdaCmd(ctx context.Context, config *LambdaInvokerConfig) {
	client, err := InitializeLambdaClient(ctx, config.Region)
	if err != nil {
		fmt.Printf("Unable to create Lambda client error=[%s]\n", err.Error())
		return
	}

	lambdaInvokerService := &LambdaInvokerService{
		LambdaClient: client,
		LambdaName:   config.LambdaName,
	}

	requestPayload, err := GetPayloadFromConfig(config)

	if err != nil {
		fmt.Printf("Unable to get request data error=[%s]\n", err.Error())
		return
	}

	apiGatewayRequest, err := CreateAPIGatewayProxyRequest(requestPayload)

	if err != nil {
		fmt.Printf("Unable to convert your request into an API Gateway Request error=[%s]\n", err.Error())
		return
	}

	payload, err := json.Marshal(apiGatewayRequest)

	if err != nil {
		fmt.Printf("Unable to convert api gateway request to byte array error=[%s]\n", err.Error())
		return
	}

	fmt.Printf("Invoking lambda=[%s] region=[%s] request=[%s]\n", config.LambdaName, config.Region, string(payload))
	response, err := lambdaInvokerService.InvokeLambda(ctx, payload)

	if err != nil {
		fmt.Printf("Unable to invoke lambda=[%s] error=[%s]\n", lambdaInvokerService.LambdaName, err.Error())
		return
	}

	fmt.Printf("Response=[%s]", response.Payload)
}
