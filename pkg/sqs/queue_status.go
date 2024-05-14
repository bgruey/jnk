package sqs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

// Send a message to the SQS
func (s *sqsClient) GetAttributes() {
	ret, err := s.svc.GetQueueAttributes(
		context.TODO(),
		&sqs.GetQueueAttributesInput{
			QueueUrl: s.QueueURL,
			AttributeNames: []types.QueueAttributeName{
				types.QueueAttributeName(types.QueueAttributeNameAll),
			},
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Attributes: %v", ret)
}
