package sqs

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

func (s *sqsClient) ReceiveMessages(timeout int32, max_messages int32) (messages []types.Message, err error) {
	iterations := max_messages / 10
	if max_messages%10 != 0 {
		iterations += 1
	}
	var msgResult *sqs.ReceiveMessageOutput

	for i := 0; i < int(iterations); i++ {
		msgResult, err = s.svc.ReceiveMessage(
			context.TODO(),
			&sqs.ReceiveMessageInput{
				MessageSystemAttributeNames: []types.MessageSystemAttributeName{
					types.MessageSystemAttributeNameAll,
				},
				MessageAttributeNames: []string{
					string(types.MessageSystemAttributeNameAll),
				},
				QueueUrl:            s.QueueURL,
				MaxNumberOfMessages: *aws.Int32(10),
				VisibilityTimeout:   timeout,
				WaitTimeSeconds:     *aws.Int32(20),
			},
		)

		if err != nil {
			return
		}

		messages = append(messages, msgResult.Messages...)
	}

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
