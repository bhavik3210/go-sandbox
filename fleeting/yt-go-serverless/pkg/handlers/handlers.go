package handlers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func GetUsers(req events.APIGatewayProxyRequest, tableName string, dynamoClient *dynamodb.Client) string {
	return "Potato"
}

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynamoClient *dynamodb.Client) string {
	return "Potato"
}

func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynamoClient *dynamodb.Client) string {
	return "Potato"
}

func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynamoClient *dynamodb.Client) string {
	return "Potato"
}

func UnhandledMethod() string {
	return "Potato"
}
