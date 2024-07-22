package main

import (
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(router)
}

func router(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Received request: Path=%s, Method=%s\n", request.Path, request.HTTPMethod)
	switch request.Path {
	case "/login":
		if request.HTTPMethod == http.MethodPost {
			return handleLogin()
		}
	case "/data":
		if request.HTTPMethod == http.MethodGet {
			return handleData()
		}
	default:
		return events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body:       "Path not found",
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 405,
		Body:       "Method not allowed",
	}, nil
}

func handleLogin() (events.APIGatewayProxyResponse, error) {
	log.Println("login")

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "login attemt!",
		Headers: map[string]string{
			"Content-Type": "text/plain; charset=utf-8",
		},
	}

	return response, nil
}

func handleData() (events.APIGatewayProxyResponse, error) {
	log.Println("data")
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Handling data",
		Headers: map[string]string{
			"Content-Type": "text/plain; charset=utf-8",
		},
	}
	return response, nil
}
