package dynamo

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// NewDynamoDB ...
func NewDynamoDB(s *session.Session) (*dynamodb.DynamoDB, error) {
	if s == nil {
		return nil, errors.New("[NewDynamoDB] Session is nil")
	}

	return dynamodb.New(s), nil
}