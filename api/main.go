package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type MyEvent struct {
	Position string `json:"position"`
}

func httpError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	position := new(MyEvent)
	err := json.Unmarshal([]byte(req.Body), position)
	if err != nil {
		fmt.Print(err.Error())
		return httpError(500)
	}

	// Validate position value
	if position.Position != "sit" && position.Position != "stand" {
		fmt.Printf("Invalid position specified: ", position.Position)
		return httpError(500)
	}

	queueURL := os.Getenv("QUEUE_URL")

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)

	_, err = svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(0),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Position": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(position.Position),
			},
		},
		MessageBody: aws.String("Position value is in the Position message attribute"),
		QueueUrl:    &queueURL,
	})
	if err != nil {
		fmt.Printf(err.Error())
		return httpError(500)
	}

	fmt.Printf("Set position to %s\n", position.Position)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("Set position to %s\n", position.Position),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
