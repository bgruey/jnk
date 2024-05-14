package sqs

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

func (s *sqsClient) ReceiveMessages(timeout int, max_messages int) (messages []types.Message, err error) {

	var msgResult *sqs.ReceiveMessageOutput
	msgsLeft := max_messages
	getXMsgs := 10 // AWS SDK max value

	for true {
		// Should only ever be 0
		if msgsLeft < 1 {
			if msgsLeft != 0 {
				panic("negative messages left, miscount in ReceiveMessages")
			}
			break
		}

		// Less than AWS Max left to get
		if getXMsgs > msgsLeft {
			getXMsgs = msgsLeft
		}

		// Get the messages
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
				MaxNumberOfMessages: *aws.Int32(int32(getXMsgs)),
				VisibilityTimeout:   *aws.Int32(int32(timeout)),
				WaitTimeSeconds:     *aws.Int32(5),
			},
		)

		if err != nil {
			return
		}
		msgsLeft -= len(msgResult.Messages)
		messages = append(messages, msgResult.Messages...)

		// Break on timeout condition, not as many messages available as requested
		if len(msgResult.Messages) < getXMsgs {
			break
		}

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
