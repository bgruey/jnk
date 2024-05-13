package sqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func (s *sqsConfig) ReceiveMessage(timeout int64, max_messages int64) (messages []*sqs.Message, err error) {
	msgResult, err := s.svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            s.QueueURL,
		MaxNumberOfMessages: aws.Int64(max_messages),
		VisibilityTimeout:   &timeout,
		WaitTimeSeconds:     aws.Int64(20),
	})

	if err != nil {
		return
	}

	messages = msgResult.Messages

	//for _, msg := range msgResult.Messages {
	//	fmt.Printf("Message ID: %s\n", *msg.MessageId)
	//	fmt.Printf("Message Body: %s\n", *msg.Body)
	//
	//	//Save Message
	//
	//	//// Delete Message
	//	//_, err1 = svc.DeleteMessage(&sqs.DeleteMessageInput{
	//	//	QueueUrl:      queueURL,
	//	//	ReceiptHandle: msg.ReceiptHandle,
	//	//})
	//	//if err1 != nil {
	//	//	panic(err1)
	//	//}
	//}
	return
}
