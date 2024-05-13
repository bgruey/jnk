package sqs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type sqsClient struct {
	QueueURL *string
	svc      *sqs.Client
}

func NewSQS(key, secret, region, queueName string) (client *sqsClient, err error) {
	client = new(sqsClient)

	_, err = fmt.Printf("Credentials: %+v",
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			key, secret, "",
		)),
	)

	if err != nil {
		panic(err)
	}

	_, err = fmt.Printf(
		"Region: %+v",
		config.WithRegion("eu-central-1"),
	)

	if err != nil {
		panic(err)
	}

	config, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			key, secret, "",
		)),
		config.WithRegion("eu-central-1"),
	)

	if err != nil {
		panic(err)
	}

	client.svc = sqs.NewFromConfig(config)

	url_result, err := client.svc.GetQueueUrl(
		context.TODO(),
		&sqs.GetQueueUrlInput{
			QueueName: aws.String(queueName),
		},
	)

	if err != nil {
		panic(err)
	}

	client.QueueURL = url_result.QueueUrl

	return
}
