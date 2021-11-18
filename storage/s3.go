package storage

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	BUCKET_NAME = "animals-bucket"
)

// Create a S3 service client
func NewSimpleStorageService(s *session.Session) (*s3.S3, error) {
	if s == nil {
		return nil, errors.New("[NewSimpleStorageServie] Session is nil")
	}
	
	return s3.New(s), nil
}

//ListBuckets
func ListBuckets(s3cli *s3.S3) (resp *s3.ListBucketsOutput, err error) {
	resp, err = s3cli.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//CreateBucket
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

//UploadObjects
func UploadObjects(s3cli *s3.S3, filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	fmt.Printf("Uploading: %v\n", filename)
	result, err := s3cli.PutObject(&s3.PutObjectInput{
		Body: f,
		Bucket: aws.String(BUCKET_NAME),
		Key: aws.String(strings.Split(filename, "/")[1]),
		ACL: aws.String(s3.BucketCannedACLPublicRead),
	})

	if err != nil {
		return err
	}
	fmt.Println("[UploadObjects]", result)
	return nil
}

//ListObjects
func ListObjects(s3cli *s3.S3) (resp *s3.ListObjectsOutput, err error) {
	result, err := s3cli.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(BUCKET_NAME),
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}

//GetObject
func GetObject(s3cli *s3.S3, filename string) error {
	fmt.Println("Dowloading: ", filename)

	result, err := s3cli.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(BUCKET_NAME),
		Key: aws.String(filename),
	})

	if err != nil {
		return err
	}

	folder := "s3-dowload"
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(folder + "/" +filename, body, 0644)
	if err != nil {
		return err
	}
	return nil
}