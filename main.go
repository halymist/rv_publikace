package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("ahoj")

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "wuttest test test!",
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
	}

	return response, nil
}
