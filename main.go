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

	// Check if the path is /login and handle POST requests
	switch request.RequestContext.HTTP.Path {
	case "/login":
		if request.RequestContext.HTTP.Method == http.MethodPost { // Handle POST method
			return handleLogin(request)
		}
	default:
		// Return 404 for paths that are not found
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 404,
			Body:       "Path not found",
		}, nil
	}
	// Return 405 for unsupported methods
	return events.APIGatewayV2HTTPResponse{
		StatusCode: 405,
		Body:       "Method not allowed",
	}, nil
}

func handleLogin(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	// Define a structure for the incoming request body (email)
	type RequestBody struct {
		Email string `json:"email"`
	}

	var body RequestBody
	// Parse the JSON body
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
