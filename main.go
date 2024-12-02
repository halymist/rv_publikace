package main

import (
	"encoding/json"
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
	// Define a structure for the incoming request body
	type RequestBody struct {
		Email string `json:"email"`
	}

	var body RequestBody
	// Parse JSON from the request body
	if err := json.Unmarshal([]byte(request.Body), &body); err != nil {
		log.Println("Failed to parse request body:", err)
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 400,
			Body:       "Invalid request body",
		}, nil
	}

	// Log the received email
	log.Printf("Received email: %s\n", body.Email)

	// Return a successful response
	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       "Email received successfully!",
		Headers: map[string]string{
			"Content-Type": "text/plain; charset=utf-8",
		},
	}, nil
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
