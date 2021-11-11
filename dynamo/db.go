package dynamo

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const TABLE_NAME = "animals"

// NewDynamoDB ...
func NewDynamoDB(s *session.Session) (*dynamodb.DynamoDB, error) {
	if s == nil {
		return nil, errors.New("[NewDynamoDB] Session is nil")
	}

	return dynamodb.New(s), nil
}

// CreateTable ...
func CreateTable(db *dynamodb.DynamoDB) error {
	result, err := db.CreateTable(&dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("CommanName"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("CommanName"),
				KeyType: aws.String("HASH"),
			},
		},
		BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
		TableName: aws.String(TABLE_NAME),
	})
	
	if err != nil {
		return err
	}

	fmt.Println("[Decribe Table]: ", result)
	return nil
}