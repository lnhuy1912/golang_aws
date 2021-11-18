package queue

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const (
	QUEUE_NAME = "animals-queue"
)

//NewSQSService
func NewSQSService(s *session.Session) (*sqs.SQS, error) {
	if s == nil {
		return nil, errors.New("[NewSQSService] Session is nil")
	}
	
	return sqs.New(s), nil
}

//CreateQueue
func CreateQueue(sqsCli *sqs.SQS) (string, error) {
	result, err := sqsCli.CreateQueue(&sqs.CreateQueueInput{
		QueueName: aws.String(QUEUE_NAME),
		Attributes: map[string]*string{
			"DelaySeconds": aws.String("60"),
			"MessageRetentionPeriod": aws.String("86400"),
		},
	})
	
	if err != nil {
		return "", err
	}

	fmt.Println("[CreateQueue] ", result)

	return *result.QueueUrl, nil
}

//SendMessage
func SendMessage(sqsCli *sqs.SQS, messageBody string, queueUrl string) error {
	result, err := sqsCli.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageBody: aws.String(messageBody),
		QueueUrl: aws.String(queueUrl),
	})
	if err != nil {
		return err
	}

	fmt.Println("[SendMessage] ", result)
	return nil
}

//ReceiveMessage
func ReceiveMessage(sqsCli *sqs.SQS, queueUrl string) error {
	result, err := sqsCli.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
            aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
        },
        MessageAttributeNames: []*string{
            aws.String(sqs.QueueAttributeNameAll),
        },
		QueueUrl: aws.String(queueUrl),
		MaxNumberOfMessages: aws.Int64(10),
		VisibilityTimeout: aws.Int64(60),
		WaitTimeSeconds: aws.Int64(0),
	})
	if err != nil {
		return err
	}

	fmt.Println("[ReceiveMessage] ", result.Messages)
	return nil
}

//GetQueueUrl
func GetQueueUrl(sqsCli *sqs.SQS, queueName string) (string, error) {
	result, err := sqsCli.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	})
	if err != nil {
		return "", err
	}

	return *result.QueueUrl, nil
}