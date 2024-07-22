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

func router(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	log.Printf("Received request: Path=%s, Method=%s\n", request.RequestContext.HTTP.Path, request.RequestContext.HTTP.Method)

	switch request.RequestContext.HTTP.Path {
	case "/login":
		if request.RequestContext.HTTP.Method == http.MethodGet {
			return handleLogin(request)
		}
	case "/data":
		if request.RequestContext.HTTP.Method == http.MethodGet {
			return handleData(request)
		}
	default:
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 404,
			Body:       "Path not found",
		}, nil
	}
	return events.APIGatewayV2HTTPResponse{
		StatusCode: 405,
		Body:       "Method not allowed",
	}, nil
}

func handleLogin(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	log.Println("login")

	response := events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       "login attempt!",
		Headers: map[string]string{
			"Content-Type": "text/plain; charset=utf-8",
		},
	}

	return response, nil
}

func handleData(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	log.Println("data")
	response := events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       "Handling data",
		Headers: map[string]string{
			"Content-Type": "text/plain; charset=utf-8",
		},
	}
	return response, nil
}
