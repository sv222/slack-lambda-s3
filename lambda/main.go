package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func handler(ctx context.Context, s3Event events.S3Event) error {
	bucketName := s3Event.Records[0].S3.Bucket.Name
	objectKey := s3Event.Records[0].S3.Object.Key

	sess := session.Must(session.NewSession())
	svc := s3.New(sess)

	// Download the file from S3
	resp, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		log.Fatalf("Failed to download file from S3: %v", err)
	}
	defer resp.Body.Close()

	// Read the contents of the file
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Print the contents of the file
	fmt.Println(string(body))

	return nil
}

func main() {
	lambda.Start(handler)
}
