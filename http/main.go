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
		fmt.Printf(err.Error())
		return httpError(500)
	}

	queueURL := os.Getenv("QUEUE_URL")

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)
	message_group_id := "main"

	_, err = svc.SendMessage(&sqs.SendMessageInput{
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Height": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(height.Height),
			},
		},
		MessageBody:    aws.String("Height value is in the Height message attribute"),
		QueueUrl:       &queueURL,
		MessageGroupId: &message_group_id,
	})
	if err != nil {
		fmt.Printf(err.Error())
		return httpError(500)
	}

	fmt.Printf("Set height to %s\n", height.Height)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("Set height to %s\n", height.Height),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
