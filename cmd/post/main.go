package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var ApiID string = ""

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	url := fmt.Sprintf("http://%s:4566/restapis/%s/local/_user_request_/v1/simple/hey-get-send-me-some-data",
		os.Getenv("LOCALSTACK_HOSTNAME"), ApiID)

	resp, err := http.Get(url)

	var message string

	if err != nil {
		message = err.Error()
	} else {
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			message = "Invalid GET response"
		} else {
			message = fmt.Sprintf("Enriched message: %s", string(body))
		}

		resp.Body.Close()
	}

	return events.APIGatewayProxyResponse{
			Body:       message,
			StatusCode: 200},
		nil
}

func main() {
	lambda.Start(Handler)
}
