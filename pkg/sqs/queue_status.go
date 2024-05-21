package sqs

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

func (s *sqsClient) GetAttributes() string {
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

	ret_b, err := json.MarshalIndent(ret.Attributes, "", "  ")

	if err != nil {
		panic(err)
	}
	return string(ret_b)
}

func (s *sqsClient) PrintAttributes() {
	fmt.Printf("Attributes for %s\n%s\n", *s.QueueURL, s.GetAttributes())
}
