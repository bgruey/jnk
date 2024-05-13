package sqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// Send a message to the SQS
func (s *sqsConfig) SendMessage(body string) (msgResult *sqs.SendMessageOutput, err error) {
	msgResult, err = s.svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(0),
		MessageBody:  aws.String(body),
		QueueUrl:     s.QueueURL,
	})
	if err != nil {
		return
	}

	return
}
