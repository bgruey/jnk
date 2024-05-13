package sqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type sqsConfig struct {
	QueueURL *string
	svc      *sqs.SQS
}

func NewSQS(key, secret, region, queueName string) (config *sqsConfig, err error) {
	config = new(sqsConfig)

	sess, err := session.NewSession(&aws.Config{

		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
	})

	if err != nil {
		return
	}

	config.svc = sqs.New(sess)

	urlResult, err := config.svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	})

	if err != nil {
		return
	}

	config.QueueURL = urlResult.QueueUrl

	return
}
