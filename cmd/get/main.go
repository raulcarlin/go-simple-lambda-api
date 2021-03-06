package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	param := request.PathParameters["param"]

	message := fmt.Sprintf(" { \"post-sent-me-this\" : \"%s\" } ", param)

	return events.APIGatewayProxyResponse{
			Body:       message,
			StatusCode: 200},
		nil
}

func main() {
	lambda.Start(Handler)
}
