package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Sample handler from AWS SQS
func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	// TODO
	return nil
}

func main() {
	lambda.Start(handler)
}
