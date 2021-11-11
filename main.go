package main

import (
	"errors"
	"fmt"
	"log"

	"dact_training/golang_aws/dynamo"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func NewSimpleStorageService(s *session.Session) (*s3.S3, error) {
	if s == nil {
		return nil, errors.New("[NewSimpleStorageServie] Session is nil")
	}
	
	return s3.New(s), nil
}

func main() {
	// Initialize a session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("test", "test", ""),
		Endpoint: aws.String("http://localhost:4566/"),
	})
	if err != nil {
		log.Fatal("[NewSession] Can't create new session.")
	}

	// Create S3 service client
	s3Cli, err := NewSimpleStorageService(sess)

	if err !=nil {
		log.Fatal(err)
	}

	fmt.Println("Session is working at ", s3Cli.Endpoint)

	// Create DynamoDB service client
	dynamoDB, err := dynamo.NewDynamoDB(sess)
	if err != nil {
		log.Fatal(err)
	}

	// Create table 
	err = dynamo.CreateTable(dynamoDB)
	if err != nil {
		log.Fatal(err)
	}
}