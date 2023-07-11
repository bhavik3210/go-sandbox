package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	region := "us-east-2"
	session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

}
