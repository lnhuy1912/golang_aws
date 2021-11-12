package s3cli

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	BUCKET_NAME = "animals"
	REGION 		= "us-east-1"
)

func ListBuckets(s3cli *s3.S3) (resp *s3.ListBucketsOutput, err error) {
	resp, err = s3cli.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func CreateBucket(s3cli *s3.S3) (resp *s3.CreateBucketOutput, err error) {
	resp, err = s3cli.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(BUCKET_NAME),
		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
			LocationConstraint: aws.String(REGION),
		},
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
			  return nil, errors.New("Bucket name already in use !")
			case s3.ErrCodeBucketAlreadyOwnedByYou:
			  return nil, errors.New("Bucket exists and is owned by you !")
			default:
			  return nil, err
			}
		  }
	}

	return resp, nil
}

