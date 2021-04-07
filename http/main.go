package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Height string `json:"height"`
}

func httpError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	height := new(MyEvent)
	err := json.Unmarshal([]byte(req.Body), height)
	if err != nil {
		return httpError(500)
	}
	// TODO: Submit height to SQS
	fmt.Printf("Set height to %s\n", height.Height)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("Set height to %s\n", height.Height),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
