package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// ListenCommand Command set for getting artifacts from artifactory
type ListenCommand struct{}

// Help text
func (l *ListenCommand) Help() string {
	helpText := `
Usage: `
	return strings.TrimSpace(helpText)
}

// Name Interface for printing Name
func (l *ListenCommand) Name() string {
	return "listen"
}

// Synopsis Interface for printing description
func (l *ListenCommand) Synopsis() string {
	return "Commands for listening"
}

func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()
}

// Run execution of ListenCommand
func (l *ListenCommand) Run(args []string) int {
	var queueURL string

	cmdFlags := defaultFlagSet(l.Name())
	cmdFlags.StringVar(&queueURL, "queue", "", "queue url")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	if queueURL == "" {
		log.Print("Missing required argument: queue")
		return 1
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)

	SetupCloseHandler()

	desk := Idasen{}

	for {
		// Do listening
		msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
			AttributeNames: []*string{
				aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
			},
			MessageAttributeNames: []*string{
				aws.String(sqs.QueueAttributeNameAll),
			},
			QueueUrl:        &queueURL,
			WaitTimeSeconds: aws.Int64(10),
		})
		if err != nil {
			log.Println(err)
			return 1
		}

		// Look at each message received (should be either 0 or 1 messages)
		for _, msg := range msgResult.Messages {
			fmt.Print("\n")
			log.Println("Received Message")
			position := *msg.MessageAttributes["Position"].StringValue
			log.Printf("Setting desk position to %s\n", position)
			// Remove message from queue now
			_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      &queueURL,
				ReceiptHandle: msg.ReceiptHandle,
			})
			if err != nil {
				log.Println(err)
				return 1
			}
			err = desk.To(position)
			if err != nil {
				log.Println(err)
				return 1
			}
			time.Sleep(20 * time.Second) // cooldown

		}
		time.Sleep(5 * time.Second)
		fmt.Print(".")
	}

}
