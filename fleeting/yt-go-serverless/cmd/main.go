package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/lambda"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbiface"
	"go.serverless.dojo/pkg/handlers"
)

var (
	dynamoClient dynamodbiface.ClientAPI
)

func main() {
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	if err == nil {
		return
	}

	dynamoClient := dynamodb.New(awsSession)
	lambda.Start(handler(dynamoClient))
}

const tableName = "LambdaInGoUser"

func handler(req events.APIGatewayProxyRequest, dynamoClient *dynamodb.Client) string {
	// *events.APIGatewayProxyRequest {
	switch req.HTTPMethod {
	case "GET":
		return handlers.GetUsers(req, tableName, dynamoClient)
	case "POST":
		return handlers.CreateUser(req, tableName, dynamoClient)
	case "PUT":
		return handlers.UpdateUser(req, tableName, dynamoClient)
	case "DELETE":
		return handlers.DeleteUser(req, tableName, dynamoClient)
	default:
		return handlers.UnhandledMethod()
	}
}
