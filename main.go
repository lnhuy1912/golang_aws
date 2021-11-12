package main

import (
	"errors"
	"fmt"
	"log"

	"golang_aws/dynamo"
	"golang_aws/entity"
	"golang_aws/helper"
	"golang_aws/dynamo/animals"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	ENDPOINT = "http://localhost:4566/"
	REGION   = "us-east-1"
)

var (
	s3Cli *s3.S3
	dynamoDB *dynamodb.DynamoDB
)

func init() {
	// Initialize a session
	sess, err := session.NewSession(&aws.Config{
		Region		: aws.String(REGION),
		Credentials	: credentials.NewStaticCredentials("test", "test", ""),
		Endpoint	: aws.String(ENDPOINT),
	})
	if err != nil {
		log.Fatal("[NewSession Error] Can't create new session.")
	}
	// Create S3 service client
	s3Cli, err = NewSimpleStorageService(sess)

	if err !=nil {
		log.Fatal(err)
	}

	fmt.Println("Session is working at ", s3Cli.Endpoint)

	// Create DynamoDB service 
	dynamoDB, err = dynamo.NewDynamoDB(sess)
	if err != nil {
		log.Fatalf("[NewDynamoDB Error] %v", err)
	}
}

func NewSimpleStorageService(s *session.Session) (*s3.S3, error) {
	if s == nil {
		return nil, errors.New("[NewSimpleStorageServie] Session is nil")
	}
	
	return s3.New(s), nil
}

func main() {
	// Create table 
	err := animals.CreateTable(dynamoDB)
	if err != nil {
		log.Fatalf("[CreateTable Error] %v", err)
	}

	// PutItem 
	newAnimal := entity.Animal{
		CommonName		: "Secretary Bird",
		ScientificName	: "Sagittarius serpentarius",
		Type			: "Birds",
		Diet			: "Carnivore",
		AverageLife		: "Ten to 15 years",
		Size			: "3.9 feet",
		Weight			: "5 to 9.4 pounds",
	}
	err = animals.PutItem(dynamoDB, newAnimal)
	if err != nil {
		log.Fatalf("[PutItem Error] %v", err)
	}

	// UpdateItem
	updateItem := entity.Animal{
		CommonName		: "Secretary Bird",
		ScientificName	: "Sagittarius serpentarius",
		Type			: "Birds",
		Diet			: "Carnivore",
		AverageLife		: "Ten to 15 years",
		Size			: "3.8 feet",
		Weight			: "5 to 9.4 pounds",
	}
	err = animals.UpdateItem(dynamoDB, updateItem)
	if err != nil {
		log.Fatalf("[UpdateItem Error]  %v", err)
	}

	// GetItem
	commonNameRequest := "Secretary Bird"
	animalItem, err := animals.GetItem(dynamoDB, commonNameRequest)
	if err != nil {
		log.Fatalf("[GetItem Error] \n%v\n", err)
	}
	fmt.Printf("[GetItem] Got a Item with key \"%s\" \n%s\n", commonNameRequest, helper.FormatStruct(animalItem))

	//DeleteItem
	animalDeleteKey := "Secretary Bird"
	err = animals.DeleteItem(dynamoDB, animalDeleteKey)
	if err != nil {
		log.Fatalf("[DeleteItem Error] \n%v\n", err)
	}
}