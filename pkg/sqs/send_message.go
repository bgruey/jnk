package sqs

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// Send a message to the SQS
func (s *sqsClient) SendMessage(body string) (msgResult *sqs.SendMessageOutput, err error) {
	msgResult, err = s.svc.SendMessage(
		context.TODO(),
		&sqs.SendMessageInput{
			DelaySeconds: *aws.Int32(0),
			MessageBody:  aws.String(body),
			QueueUrl:     s.QueueURL,
		})
	if err != nil {
		return
	}

	return
}
