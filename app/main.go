package main

import (
	"context"
	"fmt"
	"github.com/bgruey/jnk/pkg/sqs"
	"os"
	"strings"
)

func GetCSV(rows int) string {
	var buff strings.Builder
	buff.WriteString(fmt.Sprintf("id,name-for-%d", rows))
	for i := 1; i <= rows; i++ {
		buff.WriteString(fmt.Sprintf("\n%d,asd_%d", i, i*i))
	}
	return buff.String()
}

func main() {
	fmt.Println("Hello World!")
	source_queue, err := sqs.NewSQS(
		os.Getenv("AWS_ACCESS_ONE"),
		os.Getenv("AWS_SECRET_ONE"),
		"eu-central-1",
		os.Getenv("AWS_SQS_URL_ONE"),
	)
	if err != nil {
		panic(err)
	}

	dest_queue, err := sqs.NewSQS(
		os.Getenv("AWS_ACCESS_TWO"),
		os.Getenv("AWS_SECRET_TWO"),
		"eu-central-1",
		os.Getenv("AWS_SQS_URL_TWO"),
	)
	if err != nil {
		panic(err)
	}

	for i := 3; i < 8; i++ {
		source_queue.SendMessage(GetCSV(i + 1))
	}

	delCtx := context.TODO()

	fmt.Println("Sent messages, receiving now")
	msgs, err := source_queue.ReceiveMessages(10, 992)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Messages: %v", msgs)
	for _, msg := range msgs {
		dest_queue.SendMessage(*msg.Body)
	}
	fmt.Println("Deleting messages from source")
	source_queue.DeleteMessages(delCtx, msgs)

	msgs, err = dest_queue.ReceiveMessages(10, 200)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleting messages from dest")
	dest_queue.DeleteMessages(delCtx, msgs)
}
