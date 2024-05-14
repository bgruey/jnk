package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type sqsClient struct {
	QueueURL *string
	svc      *sqs.Client
}

func NewSQS(key string, secret string, region string, queueName string) (client *sqsClient, err error) {

	client = new(sqsClient)

	// Create static credentials from params
	config, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			key, secret, "",
		)),
		config.WithRegion(region),
	)

	if err != nil {
		panic(err)
	}

	// New sqs client
	client.svc = sqs.NewFromConfig(config)
	client.QueueURL = &queueName

	client.PrintAttributes()

	return
}
