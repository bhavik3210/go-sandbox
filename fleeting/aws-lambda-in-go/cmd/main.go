package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	dynaClient dynamodbiface.DynamoDBAPI
)

func main() {
	region := "us-east-2"
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		return
	}
	dynaClient = dynamodb.New(awsSession)
	lambda.Start(handler)
}
