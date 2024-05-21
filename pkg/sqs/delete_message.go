package sqs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

// Send a message to the SQS
func (s *sqsClient) DeleteMessages(ctx context.Context, msgs []types.Message) (err error) {

	// https://docs.aws.amazon.com/code-library/latest/ug/sqs_example_sqs_DeleteMessageBatch_section.html
	entries := make([]types.DeleteMessageBatchRequestEntry, len(msgs))
	for msgIndex := range msgs {
		entries[msgIndex].Id = aws.String(fmt.Sprintf("%v", msgIndex))
		entries[msgIndex].ReceiptHandle = msgs[msgIndex].ReceiptHandle
	}
	fmt.Printf("Deleting: %+v\n", entries)

	delOutput, err := s.svc.DeleteMessageBatch(
		context.Background(),
		&sqs.DeleteMessageBatchInput{
			QueueUrl: s.QueueURL,
			Entries:  entries,
		},
	)
	if err != nil {
		return
	}

	if len(delOutput.Failed) > 0 {
		err = fmt.Errorf("message deletion failed for %+v", delOutput.Failed)
		return
	}

	return
}
