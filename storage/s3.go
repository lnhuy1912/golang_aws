package storage

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	BUCKET_NAME = "animals-bucket"
	REGION 		= "us-east-1"
)

func NewSimpleStorageService(s *session.Session) (*s3.S3, error) {
	if s == nil {
		return nil, errors.New("[NewSimpleStorageServie] Session is nil")
	}
	
	return s3.New(s), nil
}

func ListBuckets(s3cli *s3.S3) (resp *s3.ListBucketsOutput, err error) {
	resp, err = s3cli.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func CreateBucket(s3cli *s3.S3) error {
	result, err := s3cli.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(BUCKET_NAME),
		
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
			  return errors.New("Bucket name already in use !")
			case s3.ErrCodeBucketAlreadyOwnedByYou:
			  return errors.New("Bucket exists and is owned by you !")
			default:
			  return err
			}
		  }
	}

	fmt.Printf("[CreateBucket] Created a Bucket: \n%v\n", result)
	return nil
}

