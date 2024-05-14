package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

// Send a message to the SQS
func (s *sqsClient) DeleteMessages(msgs []types.Message) (err error) {

	for _, msg := range msgs {
		_, err = s.svc.DeleteMessage(
			context.TODO(),
			&sqs.DeleteMessageInput{
				QueueUrl:      s.QueueURL,
				ReceiptHandle: msg.ReceiptHandle,
			},
		)
		if err != nil {
			panic(err)
		}
	}

	// Batch Delete probably faster, but unsure of Id so far.

	// s.svc.DeleteMessageBatch(
	// 	context.Background(),
	// 	&sqs.DeleteMessageBatchInput{
	// 		QueueUrl: s.QueueURL,
	// 		Entries: []types.DeleteMessageBatchRequestEntry{
	// 			Id: "???",
	// 			ReceiptHandle: msgs[0].ReceiptHandle
	// 		},
	// 	}
	// )

	return
}
