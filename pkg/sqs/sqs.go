package sqs

import (
	"context"
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

	config, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			"AWS_ACCESS_ONE", "AWS_SECRET_ONE", "",
		)),
		config.WithRegion("eu-central-1"),
	)

	if err != nil {
		return
	}

	client.svc = sqs.NewFromConfig(config)

	url_result, err := client.svc.GetQueueUrl(
		context.TODO(),
		&sqs.GetQueueUrlInput{
			QueueName: aws.String(queueName),
		},
	)

	if err != nil {
		return
	}

	client.QueueURL = url_result.QueueUrl

	return
}
