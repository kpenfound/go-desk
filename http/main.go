package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Height string `json:"height"`
}

func HandleRequest(ctx context.Context, height MyEvent) (string, error) {
	// TODO: Submit height to SQS
	return fmt.Sprintf("Set height to %s", height.Height), nil
}

func main() {
	lambda.Start(HandleRequest)
}
