package sqs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

func (s *sqsClient) ReceiveMessages(timeout int, max_messages int) (messages []types.Message, err error) {

	var msgResult *sqs.ReceiveMessageOutput
	msgsLeft := max_messages
	getXMsgs := 10 // AWS SDK max value

	for {
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
			fmt.Printf("Received only %d of %d messages, assuming queue empty.", len(messages), max_messages)
			break
		}

	}

	return
}
