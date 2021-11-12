package animals

import (
	"fmt"

	"golang_aws/entity"
	"golang_aws/helper"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const TABLE_NAME = "animals"

// CreateTable ...
func CreateTable(db *dynamodb.DynamoDB) error {
	result, err := db.CreateTable(&dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("CommonName"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("CommonName"),
				KeyType: aws.String("HASH"),
			},
		},
		BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
		TableName: aws.String(TABLE_NAME),
	})
	
	if err != nil {
		return err
	}

	fmt.Println("[CreateTable] Created a table: \n", result)
	return nil
}

// PutItem ...
func PutItem(db *dynamodb.DynamoDB, animal entity.Animal) error {
	_, err := db.PutItem(&dynamodb.PutItemInput{
		Item: 	map[string]*dynamodb.AttributeValue{
				"CommonName"	: 	{S:	aws.String(animal.CommonName),},
				"ScientificName":	{S: aws.String(animal.ScientificName)},
				"Type"			:	{S: aws.String(animal.Type)},
				"Diet"			:	{S: aws.String(animal.Diet)},
				"AverageLife"	:	{S: aws.String(animal.AverageLife)},
				"Size"			:	{S: aws.String(animal.Size)},
				"Weight"		:	{S: aws.String(animal.Weight)},
		},
		TableName: aws.String(TABLE_NAME),
	})

	if err != nil {
		return err
	}

	fmt.Printf("[PutItem] Put new item \n%s\n", helper.FormatStruct(animal))
	return nil
}

// UpdateItem ...
func UpdateItem(db *dynamodb.DynamoDB, animal entity.Animal) error {
	_, err := db.UpdateItem(&dynamodb.UpdateItemInput{
		ExpressionAttributeNames: map[string]*string {
			"#ScientificName"	:  	aws.String("ScientificName"),
			"#Type"				:	aws.String("Type"),
			"#Diet"				:	aws.String("Diet"),
			"#AverageLife"		: 	aws.String("AverageLife"),
			"#Size"				:	aws.String("Size"),
			"#Weight"			:	aws.String("Weight"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":ScientificName"	:	{S: aws.String(animal.ScientificName)},
			":Type"			 	:	{S: aws.String(animal.Type)},
			":Diet"			 	:	{S: aws.String(animal.Diet)},
			":AverageLife"	 	:	{S: aws.String(animal.AverageLife)},
			":Size"			 	:	{S: aws.String(animal.Size)},
			":Weight"		 	:	{S: aws.String(animal.Weight)},
		},
		Key:  map[string]*dynamodb.AttributeValue{
			"CommonName"		:	{S: aws.String(animal.CommonName)},
		},
		TableName: aws.String(TABLE_NAME),
		UpdateExpression: aws.String("SET #ScientificName =	:ScientificName, #Type = :Type,#Diet = :Diet,#AverageLife =	:AverageLife,#Size = :Size,#Weight = :Weight"),
	})

	if err != nil {
		return err
	}

	fmt.Printf("[UpdateItem] Updated Item: \n%s\n", helper.FormatStruct(animal))
	return nil
}

// GetItem ...
func GetItem(db *dynamodb.DynamoDB,commonName string) (animal entity.Animal, err error) {
	result, err := db.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"CommonName": {S: aws.String(commonName)},
		},
		TableName: aws.String(TABLE_NAME),
	})

	if err != nil {
		return animal, err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &animal)

	return animal, err
}

// DeleteItem ...
func DeleteItem(db *dynamodb.DynamoDB ,animalKey string) error {
	_, err := db.DeleteItem(&dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"CommonName": {S: aws.String(animalKey)},
		},
		TableName: aws.String(TABLE_NAME),
	})

	if err != nil {
		return err
	}
	fmt.Printf("[DeleteItem] Deleted Item with key \"%s\"", animalKey)
	return nil
}