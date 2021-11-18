package main

import (
	"fmt"
	// "io/ioutil"
	"log"

	"golang_aws/dynamo"
	"golang_aws/entity"
	"golang_aws/helper"
	"golang_aws/dynamo/animals"
	"golang_aws/queue"
	"golang_aws/storage"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const (
	ENDPOINT    = "http://localhost:4566/"
	REGION      = "us-east-1"
)

var (
	s3Cli *s3.S3
	dynamoDB *dynamodb.DynamoDB
	sqsCli *sqs.SQS
)

func init() {
	// Initialize a session
	sess, err := session.NewSession(&aws.Config{
		Region		: aws.String(REGION),
		Credentials	: credentials.NewStaticCredentials("test", "test", ""),
		S3ForcePathStyle: aws.Bool(true),
		Endpoint	: aws.String(ENDPOINT),
	})
	if err != nil {
		log.Fatal("[NewSession Error] Can't create new session.")
	}
	// Create S3 service client
	s3Cli, err = storage.NewSimpleStorageService(sess)
	if err !=nil {
		log.Fatal(err)
	}
	fmt.Println("Session is working at ", s3Cli.Endpoint)

	// Create DynamoDB service client
	dynamoDB, err = dynamo.NewDynamoDB(sess)
	if err != nil {
		log.Fatalf("[NewDynamoDB Error] %v", err)
	}

	// Create SQS Service client
	sqsCli, err = queue.NewSQSService(sess)
	if err != nil {
		log.Fatalf("[NewSQSService] %v", err)
	}
}

func main() {
	//Create table 
	err := animals.CreateTable(dynamoDB)
	if err != nil {
		log.Fatalf("[CreateTable Error] %v", err)
	}

	//PutItem 
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

	//UpdateItem
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

	//GetItem
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

	//CreateBucket
	// err := storage.CreateBucket(s3Cli)
	// if err != nil {
	// 	log.Fatalf("[CreateBucket Error] %v", err)
	// }

	// ListBuckets
	// listBuckets, err := storage.ListBuckets(s3Cli)
	// if err != nil {
	// 	log.Fatalf("[ListBuckets Error] %v", err)
	// }
	// fmt.Printf("[ListBuckets] \n%s\n", helper.FormatStruct(listBuckets.Buckets))
	
	// //UploadObjects
	// folder := "files"
	// files, err := ioutil.ReadDir(folder)
	// if err != nil {
	// 	log.Fatalf("[UploadObject Error] %v", err)
	// }
	// for _, file := range files {
	// 	if file.IsDir(){
	// 		continue
	// 	} else {
	// 		err = storage.UploadObjects(s3Cli, folder + "/" + file.Name())
	// 		if err != nil {
	// 			log.Fatalf("[UploadObject Error] %v", err)
	// 		}
	// 	}
	// }

	// //ListObjects
	// listObjects, err := storage.ListObjects(s3Cli)
	// if err != nil {
	// 	log.Fatalf("[ListObjects Error] %v", err)
	// }
	// fmt.Printf("[LitsObjects] \n%s\n", helper.FormatStruct(listObjects))
	
	// //GetObjects
	// for _, object := range listObjects.Contents {
	// 	err = storage.GetObject(s3Cli, *object.Key)
	// 	if err != nil {
	// 		log.Fatalf("[GetObject Error] %v", err)
	// 	}
	// }
	
	//CreateQueue
	// queueUrl, err := queue.CreateQueue(sqsCli)
	// if err != nil {
	// 	log.Fatalf("[CreateQueue Error] %v", err)
	// }

	// //SendMessage
	// messageBody := "There's a animal information about ambious that It had added."
	// err = queue.SendMessage(sqsCli, messageBody, queueUrl)
	// if err != nil {
	// 	log.Fatalf("[SendMessage Error] %v", err)
	// }

	// //ReceiveMessage
	// err = queue.ReceiveMessage(sqsCli, queueUrl)
	// if err != nil {
	// 	// log.Fatalf("[ReceiveMessage Error] %v", err)
	// }
}